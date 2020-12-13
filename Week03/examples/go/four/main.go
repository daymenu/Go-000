package main

// 对于无缓存的chan 对c的接收要先行发生于对c的发送
var c = make(chan int)
var a string = "hiiiii"

func f() {
	a = "hello world"
	<-c
}

func main() {
	go f()
	c <- 8989898
	print(a)
}
