package main

import "fmt"

//exercise to create two dimentional array
func main() {
	var matrix =[][]int{{1,2,3},{4,5,6}}
	fmt.Printf("", matrix)
	println("the length of rows is ", len(matrix))
	println("the length of column is ", len(matrix[0]))
	rows :=len(matrix)
	columns:=len(matrix[0])
	newArraysRows := columns
	newArrayColm :=rows
	// creating the new matrix is the way like here
	a := make([][]int,newArraysRows)
	for i :=0 ;i < newArraysRows ; i++{
		a[i]=make([]int, newArrayColm) // for each row how many columns
	}

	println(a)

	for i:=0 ;i < columns ; i++{
		for j:=0 ; j < rows ; j++{
			a[i][j]=matrix[j][i]
		}
	}

	fmt.Printf("the transposed matrix is " , a)
}
