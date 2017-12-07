package mystudy

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//var PersonMap map[string]PersonInfo
	//PersonMap = make(map[string]PersonInfo)

	PersonMap := make(map[string]PersonInfo)

	PersonMap["1"] = PersonInfo{"1", "luzq", "beijing"}
	PersonMap["2"] = PersonInfo{"2", "cj", "beijing"}
	PersonMap["3"] = PersonInfo{"3", "gui", "beijing"}
	fmt.Println()
	delete(PersonMap, "4")
	fmt.Println(PersonMap)

	var p1 = PersonInfo{"2", "luzq", "bejing"}
	var p2 = &PersonInfo{"3", "cj", "bejing"}

	p1.Notify()
	p1.ChangeName("luzqds")
	p1.Notify()

	p2.Notify()
	p2.ChangeName("cjds")
	p2.Notify()
	//fmt.Println(cap(PersonMap))

	// person, ok := PersonMap["2"]
	// if ok {
	// 	fmt.Println(person.Name)
	// } else {
	// 	fmt.Println("no person")
	// }

	//person := getPerson(PersonMap, "2")
	//fmt.Println(person.Name)
	//
	//person = getPerson(PersonMap, "3")
	//fmt.Println(person.Name)
	//
	//delete(PersonMap, "3")
	//
	//person = getPerson(PersonMap, "3")
	//
	//fmt.Println(person.Name)
}

func (p *PersonInfo) ChangeName(name string) {
	p.Name = name
}
func (p PersonInfo) Notify() {
	fmt.Println(p.Name)
}

//注释这样声明方法也是可以的：）
// func getPerson(personMap map[string]PersonInfo, key string) (ret PersonInfo) {
func getPerson(personMap map[string]PersonInfo, key string) PersonInfo {
	person, ok := personMap[key]
	if ok {
		return person
	} else {
		return PersonInfo{"", "", ""}
	}
}
