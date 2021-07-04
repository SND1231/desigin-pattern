package abstractfactory

import(
	"fmt"
)

func Execute(){
	panda := Factory(4)
	fmt.Printf("ぼくは%sだぞ〜、", panda.GetName())
	panda.Bark()
}
