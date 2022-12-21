package POCDEV

import "go.k6.io/k6/js/modules"


func init() {
	modules.Register("k6/x/POCDEV", new(Demo))
}

type Demo struct{}

func (*Demo) Hello(name string) string {
	return "Hello " + name
}