### 测试环境
```shell
docker version: 4.6.1
redis version: 6.0
macOS vesrion: 12.3.1
machine: 6c 16G
```

### 问题
1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
[详见](./benchmark.md)

2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

|keys|value size|memory|avg|
|----| ----  |  ----  | ----  |
|20w| 10 |  21.64M | 108.97 |
|20w| 20 | 23.16M | 116.97 |
|20w| 50 | 29.26M | 149.97 |
|20w| 100 | 39.95M | 204.97 |
|20w| 200 | 61.31M | 316.97 |
|20w| 500 | 116.24M | 604.97 |
|20w| 1024 | 262.72M | 1372.97 |
|20w| 5120 | 1.16G | 6236.97 |
