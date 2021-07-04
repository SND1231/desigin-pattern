package singleton

import(
	"fmt"
)

type singleton struct{
	id int
	name string
}

var sharedSingleton = newSingleton()

func newSingleton() *singleton{
	return &singleton{
		id: 1,
		name:"Terry",
	}
}

func GetSingleton() *singleton{
	return sharedSingleton
}

func (s *singleton) GetDetailSingleton(){
	fmt.Printf("ID:%d, Name:%s\n", s.id, s.name)
}


