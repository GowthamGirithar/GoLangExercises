package main

import (
	"regexp"
)


func main(){
	regex:=`have a nice day (.*)`
	regex1:=`(.*) have a nice day`
	str:=`have a nice day 'Gowtham Girithar'`
	str1:=`'Gowtham Girithar' have a nice day`
	var rgx =regexp.MustCompile(regex)
	var rgx1=regexp.MustCompile(regex1)
	println(rgx.FindString(str)) //have a nice day 'Gowtham Girithar'
	println(rgx.MatchString(str))//true
	data := rgx.FindStringSubmatch(str)[1]
	data1:=rgx1.FindStringSubmatch(str1)[1]
	println(data)//'Gowtham Girithar'
	println(data1)//'Gowtham Girithar'

	//case 2:
	//here regedx have ` but string have " , so not matched.
	regex2:=`have a nice day (.*)`
	str2:="have a nice day 'Gowtham Girithar'"
	var rgx2 =regexp.MustCompile(regex2)
	println(rgx2.MatchString(str2))//false
}

