### 实现
参考 https://github.com/otoolep/hraftd （照抄的）

### 查看生成的raft.db文件
1. go get github.com/br0xen/boltbrowser
2. boltbrowser raft.db


### 查询键值对
    curl -XPOST localhost:11000/key -d '{"k1":{"value":"v1","expired":1000000000}}'
    curl -XGET localhost:11000/key/k1
