package main

import (
	"fmt"
	"sort"
)
func main()  {
	sliceDemo();
	copySliceDemo();
	removeSliceElement();
	iterateSlice();
	makeSlice();
}

func makeSlice() {
	/**make will create the array with size 5 */
	println("make demo")
	a := make([]int, 5)
	print(len(a))
	print(cap(a))
	b := make([]int, 5 , 10)
	print(len(b))
	print(cap(b))
}

func iterateSlice() {
	var a=[]int{10,21,30,41,50,60}
	i :=0
	for i < len(a){
		if (a[i] % 2 == 0){
			println("Fizz")
		}
		i++
	}
}

func removeSliceElement() {
	var a=[]int{10,20,30,40,50,60};
	/**we dont have delete - we have to do append and remove */
	fmt.Println("a[:1] " ,a[:1])
	fmt.Println("a[1:3] " ,a[1:3])
	fmt.Println("a[:] " ,a[:])
	/**to remove*/
	/** ...  will make it as flat*/
	b := append(a[:2],a[3:]...);
	fmt.Println(b)
}

func copySliceDemo() {
	var a=[]int{20,10}
	/**by default copy the reference*/
	b := a;
	b[0]=14;
	fmt.Println("the a is " , a[0])
	fmt.Println("the b is " ,b[0])

	sort.Ints(a);
	fmt.Println("the b is " ,b[0])



}

func sliceDemo() {
	/**slice is not of fixed size*/
	var a=[]int{10,20};
	fmt.Println(a)
	fmt.Println("the size is " , len(a))
	fmt.Println("the capacity is " , cap(a))
	a=append(a, 20, 30);
	fmt.Println("the size is " , len(a))
	fmt.Println("the capacity is " , cap(a))
	/**capacity will be doubled for the size and on some condition , so always declare capacity if u know*/


}
