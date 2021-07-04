package factorymethod

import(
	"fmt"
)

type person interface{
	IntoroduceMyself()
}

type personImpl struct{
	name string
	age int
}

func NewPerson(name string, age int) *personImpl {
	return &personImpl{
		name: name,
		age: age,
	}
}

func(p personImpl) IntoroduceMyself(){
	fmt.Printf("私の名前は%s, %d歳です\n", p.name, p.age)
}
