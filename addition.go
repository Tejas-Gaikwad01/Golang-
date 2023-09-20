package main

import "fmt"
func Add(no1 int, no2 int)int{
	var Ans int = 0
	Ans = no1 + no2
	return Ans

}

func Sub(no1 int, no2 int) int{
	var Ans int = 0
	Ans = no1 - no2
	return Ans

}

func Calculate(no1 int, no2 int)(int,int){
	var result1 = 0
	var result2  = 0

	result1 = no1 + no2
	result2 = no1 - no2

	return result1 , result2
}



func main() {
	var Value1 = 10
	var Value2 = 11
	var Ret1 = 0
	var Ret2 = 0

	Ret1,Ret2 = Calculate(Value1 , Value2)
	fmt.Println("Addition is:",Ret1)
	fmt.Println("Substraction is :",Ret2)
	
}
