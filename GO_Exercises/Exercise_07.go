package main

import (
	"fmt"
	"runtime"
	)
func main()  {
	print(runesDemo('a'))
	fallthroughDemo()
	WhiteSpace(' ')
	println("break demo")
	breakDemo()
}


func breakDemo() {
	println(" Break is used in switch to stop executing the following statements")
	switch 1 {
	case 1:
		fmt.Println("1")
		break
		fmt.Println("11")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

/**it will proceed to the next statement*/
func fallthroughDemo() {
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}
// r is character
// argument and return type

func runesDemo(r rune) rune{

	// expression and the variable to be used for case
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	/**you can use expression too in switch */
	switch {
	case 97 <= r && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}