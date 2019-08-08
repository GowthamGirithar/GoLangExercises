package main

type Person struct {
	name string
	age int
	Address struct{
		streetLine1 string
		streetLine2 string
		city string
		state string
		country string
	}
}

func  main()  {
	p :=getPersonData()
	println(p.name)
	println(p.Address.city)

	p1, _ := createPersonUsingNew();
	println(p1.name)
	println(p1.Address.state)
}

func createPersonUsingNew()  (*Person  , string){
	p1:= new(Person)
	p1.name="Gowtham"
	p1.Address.state="Karanataka"
	return p1 , p1.name;
}

func getPersonData() Person {
	p := Person{
		name: "Gowtham",
		age:  27,
		Address: struct {
			streetLine1 string
			streetLine2 string
			city        string
			state       string
			country     string
		}{streetLine1:"", streetLine2:"", city:"Bangalore", state:"", country:"India"},
	}
	p1 := p;
	/** struct copies by value and not reference  (slice - reference , array - values , map - reference)*/
	p1.name ="Gowgi"
	println(p.name)
	println(p1.name)
    return p;
}