Postel’s Law
A good maxim for writing good software is Postel’s Law. It’s often said as

> “Be conservative with what you do, be liberal with you accept”

Translated to Go, the saying is as goes:

> “Accept interfaces, return structs”
  


* cycle import error
package a depends on package b, package b depends on package a

https://github.com/golang/go/wiki/CodeReviewComments#interfaces

struct field name and method name can't be the same(case-sensitive)

`init` method will be executed automatically before `main`

Data is transferred over TCP as a stream of bytes, that means:
* the application need to serialize/deserialize data before send data to client or after receive data from a client.
* the total size of data in stream is unknown to the receiver
* the sender may split data into multiple packets and sent one by one over tcp
* data packet arrival order is unknown due to the network uncertainty
* the sender may send multiple data

so we need a way to: 
* tell the receiver the total length of data
* the order of chunks
* the data identity which the packets belong to

That is a protocol:
header: length, identity
data: bytes

zinx use TLV serialization 

https://www.cnblogs.com/zhangxian/articles/4990997.html

https://blog.csdn.net/solstice/article/details/6300108 and comments

https://my.oschina.net/maxid/blog/206546?tdsourcetag=s_pctim_aiomsg

[How To Use Go Interfaces](https://blog.chewxy.com/2018/03/18/golang-interfaces/)

[Golang Interface](https://www.golangprograms.com/go-language/interface.html)

[Go 语言设计与实现](https://draveness.me/golang/)
[](https://draveness.me/)
[上下文 Context](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/)

[Golang Context 是好的设计吗？](https://segmentfault.com/a/1190000017394302)

## v0.7 decouple read/write

use channel to communicate between reader and writer

## v0.8 worker pool

create work pool to process request
* one worker has one task queue
* message handler dispatch tasks over the queues for workers based on the connection id

## v0.9 connection manager
maintain the current active connections, can add/remove/clear connections
add hooks after connection created or destroyed

## v1.0 add connection properties
