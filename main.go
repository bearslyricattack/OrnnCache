package main

import (
	"os"
	"text/template"
)

// 装饰者模式的功能演示

////只使用基本功能
//var base baseclient.BaseClient
//base.Set(ctx, "test", "test", 0)
//
////使用读写锁
//var mutex mutexclient.MutexClient
//mutex.BaseClient = &base
//mutex.Set(ctx, "test", "test", 0)
//
////使用缓存屏蔽功能
//var mock mockclient.MockClient
//mock.BaseClient = &base
//mock.Set(ctx, "test", "test", 0)
//
////同时使用这两个功能
//mock.BaseClient = &mutex
//mock.Set(ctx, "test", "test", 0)

// 演示如何使用装饰者模式，根据自己的需求组装对象并使用

// Item 定义gorm model
type Item struct {
	Name    string `gorm:"column:name; type:varchar(32);"`
	Content string `gorm:"column:content; type:varchar(512);"`
	Author  string `gorm:"column:author; type:varchar(512);"`
	Type    int    `gorm:"column:type; type:varchar(32);"`
}

// TableName 定义表名
func (i Item) TableName() string {
	return "Item"
}

//client 演示
//func main() {
//	//新建数据库
//	db := mysql2.New()
//	//redis
//	//redisClient := redis.New()
//	////内存数据库
//	base := baseclient.New()
//	//新建数据库表,根据结构体创建
//	//m := db.Migrator()
//	//err := m.CreateTable(&Item{})
//	//if err != nil {
//	//	panic("新建表失败!")
//	//}
//
//	//创建路由
//	r := gin.Default()
//
//	//随机获取其中一个元素
//	r.GET("/poem", func(c *gin.Context) {
//		//获取随机key
//		randomKey := base.RandomKey()
//		//先查缓存
//		res, err := base.Get(context.Background(), randomKey)
//
//		if err {
//			item := res.(Item)
//			c.JSON(http.StatusOK, gin.H{
//				"content": item.Content,
//				"name":    item.Name,
//				"author":  item.Author,
//			})
//			return
//		}
//		//缓存没查到，就去查mysql
//		var mysql Item
//		db.First(&mysql, "name = ?", randomKey)
//		//查到之后 写回内存缓存
//		base.Set(context.Background(), randomKey, mysql, 0)
//		//返回
//		c.JSON(http.StatusOK, gin.H{
//			"content": mysql.Content,
//			"name":    mysql.Content,
//			"author":  mysql.Author,
//		})
//		return
//	})
//	//写入
//	r.GET("/set", func(c *gin.Context) {
//		name := c.Query("name")
//		content := c.Query("content")
//		author := c.Query("author")
//		//组装对象
//		item := &Item{
//			Name:    name,
//			Content: content,
//			Author:  author,
//		}
//		//存入内存
//		base.Set(context.Background(), name, item, 0)
//		//存入mysql
//		db.Create(&item)
//	})
//	//把数据库中的值存入缓存
//	r.GET("/read", func(c *gin.Context) {
//		var s []Item
//		var swap Item
//		db.Find(&s)
//		for _, v := range s {
//			swap = v
//			base.Set(context.Background(), v.Name, swap, 0)
//		}
//	})
//	r.Run(":8081")
//}

// 代码生成工具演示
func main() {
	// 模板字符串

	const base = `var base baseclient.BaseClient 
var {{.Name}} mutexclient.MutexClient 
mutex.BaseClient = &base`
	// 准备模板数据
	data := struct {
		Name string
	}{
		Name: "mutex",
	}

	// 创建模板对象并解析模板字符串
	tmpl, err := template.New("test").Parse(base)
	if err != nil {
		panic(err)
	}

	// 使用数据渲染模板，并将结果输出到标准输出
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
