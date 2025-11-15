package main

import "fmt"

func main(){
	var students [5]string

	students[0]="Jaba"
	students[1]="Mark"
	students[2]="John"
	students[3]="Joy"
	students[4]="Shiro"

	var i int
	for i=0; i<5; i++{
		fmt.Println(students[i])
	}
	
}