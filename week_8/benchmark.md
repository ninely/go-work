### redis benchmark
* redis-benchmark -d 10 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.68 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

21358.39 requests per second

====== GET ======
  100000 requests completed in 4.44 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

22537.75 requests per second
```

* redis-benchmark -d 20 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.45 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

22497.19 requests per second

====== GET ======
  100000 requests completed in 4.38 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

22810.22 requests per second
```

* redis-benchmark -d 50 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.86 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20563.44 requests per second

====== GET ======
  100000 requests completed in 4.86 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20593.08 requests per second
```

* redis-benchmark -d 100 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.77 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20959.97 requests per second

====== GET ======
  100000 requests completed in 4.86 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20567.67 requests per second
```

* redis-benchmark -d 200 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.98 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20096.46 requests per second

====== GET ======
  100000 requests completed in 5.01 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

19944.16 requests per second
```

* redis-benchmark -d 1024 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.91 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20370.75 requests per second

====== GET ======
  100000 requests completed in 5.76 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

17346.05 requests per second
```

* redis-benchmark -d 5120 -t get,set
```shell
====== SET ======
  100000 requests completed in 4.92 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20321.07 requests per second

====== GET ======
  100000 requests completed in 4.80 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

20828.99 requests per second
```