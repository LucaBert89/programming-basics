package main

import (
	"fmt"
)


func main() {
	width := 0  
	length := 0
	fmt.Print("insert the width of the room")  
	fmt.Scanln(&width)  
	fmt.Print("insert the length of the room")  
	fmt.Scanln(&length)    
	var area = width * length; 
	fmt.Println("Area of the room is ", area)
 }  
 
	




