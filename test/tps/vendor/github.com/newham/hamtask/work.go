package hamtask

import (
	"reflect"
)

type Data struct {
	value interface{}
}

func Int(i int) Data {
	return Data{i}
}

func String(str string) Data {
	return Data{str}
}

func Int64(i int64) Data {
	return Data{i}
}

func Float32(f float32) Data {
	return Data{f}
}

func Float64(f float64) Data {
	return Data{f}
}

func (d Data) Value() interface{} {
	return d.value
}

func (d Data) String() string {
	if d.Type() == "string" {
		return d.value.(string)
	}
	return ""
}

func (d Data) Int64() int64 {
	if d.Type() == "int64" {
		return d.value.(int64)
	}
	return 0
}

func (d Data) Int() int {
	if d.Type() == "int" {
		return d.value.(int)
	}
	return 0
}

func (d Data) Float32() float32 {
	if d.Type() == "float32" {
		return d.value.(float32)
	}
	return 0.0
}

func (d Data) Float64() float64 {
	if d.Type() == "float64" {
		return d.value.(float64)
	}
	return 0.0
}

func (d Data) Type() string {
	return reflect.TypeOf(d.value).String()
}

type Worker interface {
	Put(Data)
	Puts([]Data)
	Close()
	Start()
	Wait()
	WaitClose()
}

type worker struct {
	Worker
	wType        int
	threadNum    int
	loop         int
	array        chan Data
	done         chan bool
	consumer     func(int, Data)
	producer     func() Data
	fullProducer func(chan Data)
}

const (
	normal    = 0
	simple    = 1
	recursion = 2
)

/**
一般的多线程模型：
1个生产者，自行控制生产队列
n个消费者
*/
func NewWorker(threadNum int, consumer func(int, Data)) Worker {
	return &worker{threadNum: threadNum, array: make(chan Data, threadNum), done: make(chan bool, threadNum), consumer: consumer, wType: normal}
}

func NewFullWorker(threadNum int, consumer func(int, Data), producer func(chan Data)) Worker {
	return &worker{threadNum: threadNum, array: make(chan Data, threadNum), done: make(chan bool, threadNum), fullProducer: producer, consumer: consumer, wType: recursion}
}

/**
简单的多线程模型：
1个生产者，生产者执行loop次相同的任务
n个消费者
*/
func NewSimpleWorker(threadNum int, consumer func(int, Data), producer func() Data, loop int) Worker {
	return &worker{threadNum: threadNum, array: make(chan Data, threadNum), done: make(chan bool, threadNum), consumer: consumer, producer: producer, loop: loop, wType: simple}
}

func (w *worker) get(i int) {
	for {
		url, ok := <-w.array
		if ok {
			w.consumer(i, url)
		} else { //收到close信号
			break
		}
	}
	w.done <- true //通知已经结束
}

func (w *worker) Put(data Data) {
	w.array <- data
}

func (w *worker) Puts(dataArray []Data) {
	for _, data := range dataArray {
		w.array <- data
	}
}

func (w *worker) Start() {
	for i := 1; i <= w.threadNum; i++ { //多线程，启动消费者
		go w.get(i)
	}
	switch w.wType {
	case normal:
		break
	case simple: //简单模式，启动一个生产者
		go func() {
			for i := 0; i < w.loop; i++ {
				w.Put(w.producer())
			}
			w.Close()
		}()
		w.Wait() //等待所有消费者结束任务
		break
	case recursion:
		go func() {
			w.fullProducer(w.array)
			w.Close()
		}()
		w.Wait() //等待所有消费者结束任务
		break
	}

}

func (w *worker) Close() {
	close(w.array)
}

func (w *worker) Wait() {
	for i := 1; i <= w.threadNum; i++ { //等待消费者完成
		<-w.done
	}
}

func (w *worker) WaitClose() {
	w.Close()
	w.Wait()
}
