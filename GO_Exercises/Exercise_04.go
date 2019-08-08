package main



func main()  {
	createMap();
}

func createMap() {
	/**map of key int and value is string*/
	var c =map[int]string{1:"Apple" , 2 : "Bat"}
	for k , v := range c{
		println(" the key and value is " , k , v)
	}
	/**add the data to the map*/
	c[4] = "hello";
	/**remove the data to the map*/
	delete(c , 2)
	/**to get the values  normally it will return 0 if not there - but it is not correct*/
	/**to get to know whether it is there or not ok is used and _ means tells compiler is is not used and dont allocate memory*/
	_ , ok := c[1]
	println(ok)

	value, ok := c[1]
	println(value)

}


