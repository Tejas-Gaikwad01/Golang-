
package main
import "fmt"

var no = 21

func Demo(){
	no := 11
	fmt.Println("From Demo",no)
 }

 func main(){
	Demo()
	fmt.Println("From main", no)
 }