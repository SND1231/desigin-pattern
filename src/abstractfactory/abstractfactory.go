package abstractfactory

import(
	"fmt"
)

type Animal interface{
	GetName()(string)
	Bark()
}

func Factory(no int) Animal {
	switch no {
		case 1:
			return NewPanda("シャンシャン")
		default:
			return NewDog(("ちくわ"))
	}
}

type Panda struct{
	name string
}

func NewPanda(name string) *Panda{
	return &Panda{
		name: name,
	}
}

func (p *Panda) GetName() string{
	return p.name
}

func (p *Panda) Bark() {
	fmt.Println("パンダだぞー")
}


type Dog struct{
	name string
}

func NewDog (name string) *Dog{
	return &Dog{
		name: name,
	}
}

func (d *Dog) GetName() string{
	return d.name
}

func (d *Dog) Bark() {
	fmt.Println("ワンワン")
}