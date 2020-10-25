# hamtask

A simple multi-threaded framework for golang

## simplest example
```golang
NewFullWorker(3, func(i int, data Data) {
    fmt.Printf("no%d:%d\n", i, data.Int())
}, func(array chan Data) {
    for _, i := range []Data{Data{1}, Data{2}, Data{3}, Data{4}} {
        array <- i
    }
}).Start() //启动消费者，开始等待生产者生产
```
## common example
```golang
w := NewWorker(3, func(i int, data Data) {
    fmt.Printf("no%d:%d\n", i, data.Int())
})
w.Start() //启动消费者，开始等待生产者生产

w.Puts([]Data{Data{1}, Data{2}, Data{3}, Data{4}})

w.Close() //结束生产队列，通知所有消费者

w.Wait() //等待消费者执行完任务
```

## more example 
**see** [work_test.go](/work_test.go)  