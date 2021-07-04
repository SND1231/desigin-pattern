package adapter

type Hanako struct {
	taro Taro
}

func NewHanako() Hanako {
	var h Hanako
	h.taro = Taro{}
	return h
}

func (h Hanako) OrganizeClass() {
	h.taro.EnjoyClass()
}
