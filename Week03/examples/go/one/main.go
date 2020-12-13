package main

var a string = "hi"

func f() {
	print(a)
}
func hello() {
	a = "hello, world"
	go f()
}
func main() {
	hello()
}
