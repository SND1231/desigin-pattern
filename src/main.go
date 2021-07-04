package main

import(
	"github.com/SND1231/design-pattern/adapter"
	"github.com/SND1231/design-pattern/factorymethod"
	"github.com/SND1231/design-pattern/singleton"
	"github.com/SND1231/design-pattern/abstractfactory"
)

func main() {
	// adapterを実行
	adapter.Execute()
	// facatorymethodを実行
	factorymethod.Execute()
	// singletonを実行
	singleton.Execute()
    // abstractfactoryを実行
	abstractfactory.Execute()
}
