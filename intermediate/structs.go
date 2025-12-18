package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
	adress    Adress
}

type Adress struct {
	city    string
	country string
}

func main() {

	p := Person{
		firstname: "rohit",
		lastname:  "Gajbhiye",
		age:       18,
		adress: Adress{
			city:    "Nagpur",
			country: "India",
		},
	}

	p1 := Person{
		firstname: "Tushar",
		age:       25,
	}

	fmt.Println(p.firstname)
	fmt.Println(p.lastname)

	fmt.Println(p.adress.city)

	// anonymous structs

	user := struct {
		username string
		email    string
	}{
		username: "rohitgajbhye",
		email:    "rohit@gmail.com",
	}

	fmt.Println(user)
	fmt.Println(p.fullname())

	fmt.Println(p == p1)

	fmt.Println("Age value before increment ", p.age)
	p.increment()
	fmt.Println("Age value after increment ", p.age)
}

func (p Person) fullname() string {
	return p.firstname + " " + p.lastname
}

// pointers allows to modify the original struct but pass by reference doesn't

func (p *Person) increment() {
	p.age++
}
