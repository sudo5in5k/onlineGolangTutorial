package main

import "fmt"

func main() {
	n := 100
	fmt.Println(n)

	p := &n
	fmt.Println(p)
	fmt.Println(*p)

	newSandBox()

	person := Person{}
	fmt.Println(person)

	person2 := &Person{} // = new(Person)
	changeName(person2)
	fmt.Println(*person2)
}

/*
 If u will return pointer type values, use new(), else make()
*/
func newSandBox() {
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)
	/*
	 Error
	*/
	// fmt.Println(*p2)
}

type Person struct {
	name, address string
	tel, age      int
}

func changeName(p *Person) {
	p.name = "Gopher" // = (*p).name
}
