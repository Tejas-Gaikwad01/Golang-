package main
import "fmt"


 func main(){
	no := 11
	p := &no 
	
	fmt.Println("value of no ",no)
	fmt.Println("value of p ",p)
	fmt.Println("Address of no ",&no)
	fmt.Println("value of no through p",*p)
 }