![image](/img/a246fc4f1afd9d4.jpg)
# OrnnCache
自由插装，动态拓展，快速构建的分布式内存数据库.

OrnnCache是使用Go语言编写的分布式内存数据库，它内嵌在应用程序代码中，可以避免使用集中式缓存所带来的网络访问、网络传输、数据复制、序列化和反序列化等操作所导致的延迟问题，进一步提升接口响应速度，节省CPU性能。

通过以嵌套的Go接口与范型为核心的构建方式和包含众多模版的代码生成器提供了快速搭建的功能，无论你是想要为已经写好的后端服务搭建一套并发量更高的前台服务，还是需要像etcd这样同时满足高可用与高可靠的数据库来存储关键数据，都可以使用简单的定义文件快速生成一套完整的代码使用，无需做额外的配置。

基于接口构建解决了在单元测试时的mock问题，提供的ctx切换策略可以为集成测试人员提供缓存屏蔽，实现全链路mock，简化测试流程。

除了在缓存的基本功能之外，OrnnCache还支持隔离级别（锁）控制，多级数据源配置（比如Redis/Mysql/http-server/rpc-server），数据更新策略（mq/raft）,缓存过期策略（手动过期/时间被动过期/LRU），缓存异常处理（击穿/穿透/雪崩）等功能，全部以插装的方式实现。

# 1.使用：

目前项目不支持使用依赖的方式进行注入。只能使用源代码的方式进行导入。

导入项目之后，根据main函数之中的组装的思想,按照自己所需的功能组装相应的对象，然后直接使用其中的方法即可。

# 2.架构：

## 1.1 整体架构：

项目架构由三部分组成：基本代码，代码生成器，示例客户端。

基本代码为每个独立功能的代码块，包含接口，对象，以及他们的实现。

代码生成器主要生成对象的组装部分，可以根据用户命令行的输入，生成对应的组装对象的方法。

示例客户端为一个使用当前内存数据库开发的项目的demo。

## 1.2 数据结构：

当前使用map[string]interface{}为基本的数据结构，后续考虑采用[]byte作为另外可选的数据结构。

# 3.设计：

## 3.1 装饰者模式：

装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。

**装饰者模式要求对象之间关系的拓展依赖于组合而不是继承**，天生适合go语言。

以基本功能为例子。首现定义数据库的基本的接口：

```
// Client 内存缓存的基本接口
type Client interface {
	Set(ctx context.Context, k string, x interface{}, d time.Duration)
	Get(ctx context.Context, k string) (interface{}, bool)
	Replace(ctx context.Context, k string, x interface{}, d time.Duration) error
	Delete(ctx context.Context, k string) (interface{}, bool)
	ItemCount(ctx context.Context) int
	Flush(ctx context.Context)
	DeleteExpired(ctx context.Context)
}
```

然后写基本的实现，没有锁/日志/缓存屏蔽以及其他的数据结构。

然后每个功能的实现都有一个自己的对象，这些对象接收一个原接口类型变量作为参数，导致他们不管是以何种顺序都可以自由的进行拓展，并且可以支持mock。在这些对象中调用上级对象的同名方法，并在这些方法周围以类似环绕通知的方式实现其功能。

以简单的使用互斥锁的client为例子：

```
func (m *MutexClient) Set(ctx context.Context, key string, value interface{}, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.BaseClient.Set(ctx, key, value, d)
}
```

## 3.2 代码结构：

代码结构截图如下：

![image-20240404200514263](https://bearsblog.oss-cn-beijing.aliyuncs.com/img/image-20240404200514263.png)

### 3.2.1 basefunction：

基本功能的实现，实现了最外层的interface接口。

baseclient为基本实现。mockclient为缓存屏蔽的实现，mutexclient为使用互斥锁屏蔽的实现。

### 3.2.2 disusestrategy：

缓存淘汰的策略。

fifo是先进先出队列的实现。lfu为最小使用次数的实现，lru为最少访问的实现。

### 3.2.3 expectionhandle：

缓存异常处理的策略。

avalanche：缓存雪崩的解决方案：给key的过期时间设置随机数，在达到一定的阈值时缓存空key。

penetration：缓存穿透问题的解决方案：使用布隆过滤器和给出缓存空key的两种解决方案。

puncture：缓存击穿的解决方案：使用singlefilght的解决方案（singlefight的实现似乎有点问题，后续考虑使用范型进行拓展。）

### 3.2.4 mock：

针对外部interface的mock代码，使用go mock mockgen生成。目前go mock不支持范型，后续考虑使用范型进行优化。

### 3.2.5 mysql:

mysql的client 只提供了new方法。使用gorm。

### 3.2.6 redis：

redis的client，基于官方的client开发，包含缓存屏蔽和mock。

### 3.2.7 remoteservice：

远程服务，获取动态数据源的方法，dadahandle为使用范型进行数据库处理（可拓展）

docall真正发送数据的方法。

sericvediscovery为服务发现的代码，当前使用nacos进行服务发现。

### 3.2.8 main:

主要三部分组成：

第一部分为装饰者模式的代码组装的演示。

第二部分为编写的客户端的demo。

第三部分为代码生成方式的演示，生成了基本的组装代码。

# 4.拓展：

当前项目十分不完善，有非常非常多可以优化的地方。

## 4.1 基本功能开发：

### 4.1.1 数据源：

进一步优化三个数据源的相关方法开发。

### 4.1.2 数据被动更新：

增添mq发送消息，主动过期数据的相关代码与代码生成的模版。

### 4.1.3 raft:

为项目提供raft的实现，配合数据源开发。

### 4.1.4 代码生成器：

现在项目的代码生成只有基本的演示demo，后续可以扩展成一个通过命令行的输入，自动生成代码的工具（参考mockgen）

### 4.1.5 更多的隔离级别：

当前只有互斥锁的级别，可以参考数据库中隔离级别的概念，使用更多的不同的锁实现，自由控制隔离级别。

### 4.1.6 其他你认为有用的任何功能

# 5.client：

使用OrnnCache开发的client是一个句子缓存软件，可以用来缓存自己喜欢的诗句和歌词。

主要包含三个方法：set，get，预热（把数据库中的数据加载进缓存）。

访问主要使用命令行的curl命令（后端程序员独有的浪漫0.0）

部分数据截图如下：

![image-20240404202545133](https://bearsblog.oss-cn-beijing.aliyuncs.com/img/image-20240404202545133.png)

# 6.最后：

Ornn是英雄联盟中，弗雷尔卓德的一名半神工匠。主掌着锻造和工艺。他在名为炉乡的火山下的溶洞中凿出了一座雄伟的工坊，独自一人在里头干活。他摆弄着熔岩沸腾的坩埚，提炼矿石，打造出无与伦比的精良物件。

其实后端程序员，很多时候在做的也是像Ornn一样的工作：打造项目/产品，华丽的也页面和功能之后，隐藏着无比丰富的优化和细节。这也是client没有ui界面，而使用命令行访问的一种设计理念：独属于幕后的创作者的退避和谦卑。

希望我们每个人都能像Ornn一样，设计并实现出令自己满意的神器作品。

如果你对这个项目和client有任何的想法，请联系我：2669184984@qq.com.
