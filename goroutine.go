//1. 数据生产者只负责创造数据
//2。 消费者拿取创造的数据
//3. main方法开通两个【(goroutine)虚拟线程】，将生产数据送给消费者方法
//4. 使用main方法进行输出
package main

//三个包进行编译
import (
	"fmt"       //导入格式化
	"math/rand" //随机数
	"time"      //时间
)

//数据生产者
func producer(header string, channel chan<- string) {
	//无限循环，不停生产数据
	for {
		//将随机数和字符串格式化为字符串发送给通道
		//字符串+随机数 格式化为字符串
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		//等待1秒
		time.Sleep(time.Second)
	}
}

//数据消费者
func customer(channel <-chan string) {
	//不停地获取数据
	for {
		//从通道口中取出数据，此处会阻塞直到信道中返回数据
		message := <-channel
		//打印数据
		fmt.Println(message)
	}

}
func main() {
	//创建一个字符串类型的通道
	channel := make(chan string)
	//创建producer() 函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	//数据消费函数
	customer(channel)
}
