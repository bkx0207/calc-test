package main

import (
	"fmt"
	"git/calc"
)

func main()  {
	res:=calc.Add(10,20)
	fmt.Println("Add(10,20):",res)
	res=calc.Sub(30,20)
	fmt.Println("Sub(30,20)",res)
	res=calc.Multi(10,20)

	fmt.Println("Multi(10,20)",res)
	res=calc.Div(30,10)

	fmt.Println("Div(30,10)",res)

}