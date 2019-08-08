package main

/**iota always starts with 0 and pattern follows */
/** cat will be 1 and cow will be 2*/
const (
	dog = iota
	cat
	cow
)
/**iota always starts with 0 and pattern follows  and now dogHouse= 1 and _ means compiler consider this as not useful */
const (
	_= iota
	dogHouse
	catHouse
	cowHouse

)

const (
	_= iota *10
	dogHouseValue
	catHouseValue
	cowHouseValue

)
const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
	PETABYTE
	EXABYTE
)
func  main()  {
	enumerationDemo();
}
func enumerationDemo() {
	value := 2
	print(value==catHouse)
	print(catHouseValue)
}