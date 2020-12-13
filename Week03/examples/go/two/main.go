package main

func main() {
	var a string
	go func() { a = "hello" }()
	print(a)
}
