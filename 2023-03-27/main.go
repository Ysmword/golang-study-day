package main

type S struct {
}

func f(x interface{}) {

}

func g(x *interface{}) {

}

func main() {
	s := S{}
	p := &S{}
	f(s)
	// g(s)
	f(p)
	// g(p)
}
