package second

import (
	"fmt"
	"strconv"
)

type Sayer interface {
	Say() string
}

type Nicknamed interface {
	getNickName() string
}

type Silenced interface {
	silence()
}

type FamilyMember interface {
	Nicknamed
	Sayer
	//PrintInChar()
}

//func (p *Person) PrintInChar() {
//	fmt.Printf("%s: '%s'\n", p.getNickName(), p.Say())
//}
//
//func (p *Child) PrintInChar() {
//	fmt.Printf("%s: '%s'\n", p.getNickName(), p.Say())
//}
//
//func (p *Dog) PrintInChar() {
//	fmt.Printf("%s: '%s'\n", p.getNickName(), p.Say())
//}

func PrintInChar(m FamilyMember) {
	fmt.Printf("%s: '%s'\n", m.getNickName(), m.Say())
}

type Human struct {
	name string
	age  int
}

func (p *Human) getNickName() string {
	return p.name + strconv.Itoa(p.age)
}

func (p *Human) silence() {
	p.name = ""
}

type Person struct {
	Human
	job string
}

func (p *Person) Say() string {
	return fmt.Sprintf("My name is %s I am %d My job is %s", p.name, p.age, p.job)
}

type Child struct {
	Human
	favToy string
}

func (p *Child) Say() string {
	return fmt.Sprintf("My name is %s I am %d My fav toy is %s", p.name, p.age, p.favToy)
}

type Dog struct {
	name   string
	phrase string
}

func (p *Dog) Say() string {
	return fmt.Sprintf("%s", p.phrase)
}

func (p *Dog) getNickName() string {
	return p.name
}

//func takeFM(fm *FamilyMember) {
//	fm.
//}

func getFM(name string, age int, otherInf string) FamilyMember {
	switch {
	case age == -1:
		return &Dog{name, otherInf}
	case age > 18:
		return &Person{Human{name, age}, otherInf}
	default:
		return &Child{Human{name, age}, otherInf}
	}
}

func UseEntities() {
	{
		var fm = getFM("132", 1, "")

		switch fm.(type) {
		case *Person:
			fmt.Println("Person", fm.(*Person).name)
		case *Child:
			fmt.Println("Child", fm.(*Child).name)
		case *Dog:
			fmt.Println("Dog", fm.(*Dog).name)
		}
	}

	{
		fm := getFM("Sasha", 14, "PC")

		if s, ok := fm.(Silenced); ok {
			s.silence()
		}

		fmt.Println(fm)
	}

	{
		members := []FamilyMember{
			&Person{Human{"Sasha", 18}, "Teacher"},
			&Child{Human{"Ivan", 8}, "toy car"},
			&Dog{"Wolf", "Af af"},
		}

		for _, m := range members {
			PrintInChar(m)
		}
	}
}
