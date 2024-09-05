package main

import(
	"fmt"
	"time"
)
func ping(c1 chan string)  {
	for i := 0; ; i++{
		c1 <- "ping"
	}
}
func pong(c2 chan string){
	for i := 0; ; i++{
		c2 <- "pong"
	}
}
func imprimir(c1 chan string, c2 chan string){
	for {
		msgPing := <- c1
		fmt.Println(msgPing)
		time.Sleep(time.Second * 1)

		msgPong := <- c2
		fmt.Println(msgPong)
		time.Sleep(time.Second * 1)

	}
}
func  main(){
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go ping(c1)
	go pong(c2)
	go imprimir(c1, c2)

	var ping string 
	fmt.Scanln(&ping)	

	var pong string 
	fmt.Scanln(&pong)	
}
