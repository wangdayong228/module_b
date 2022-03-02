package moduleb

import "github.com/wangdayong228/module_a/m1"

func B() {
	println("moduleb.B")
	// modulea.A()
	m1.Foo()
}
