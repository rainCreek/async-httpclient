package main

import (
	//"async-httpclient/entities"
  "./entities"
	"fmt"
)


func main()  {
	fmt.Println("Sync cost : ")
	entities.Synchronous()
	fmt.Println("Async cost : ")
	entities.Asynchronous()
}

