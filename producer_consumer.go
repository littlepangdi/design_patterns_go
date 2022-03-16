package main

import "fmt"

func producerWithBuffer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("produce item no.%d\n", i)
		out <- i
	}
	close(out)
}

func consumerWithBuffer(in <-chan int) {
	for data := range in {
		fmt.Printf("consume item no.%d\n", data)
	}
}

//RunProducerConsumerWithBuffer
//之所以终端生产者连续打印了大于缓冲区容量的数据，
//是因为终端打印属于系统调用也是有延迟的，
//IO操作的时候，生产者同时向管道写入，请求打印，管道的写入读取与终端输出打印速度不匹配。
func RunProducerConsumerWithBuffer() {
	ch := make(chan int, 5)
	go producerWithBuffer(ch)
	consumerWithBuffer(ch)
}
