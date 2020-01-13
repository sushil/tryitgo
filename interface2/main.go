package main

import (
	"fmt"
	"reflect"
)

func main() {
	o := newEsIndexOperator()
	h := newHoster()

	e1 := h.host(o)
	fmt.Println(e1)

	e2 := h.host(&o)
	fmt.Println(e2)

	e3 := h.host(Hoster{})
	fmt.Println(e3)

	e4 := h.host(&Hoster{})
	fmt.Println(e4)

	var h2 hoster
	h2 = newHoster()
	e5 := h.host(h2)
	fmt.Println(e5)

}

var noPointerErr = fmt.Errorf("not pointer")

type hoster interface {
	host(v interface{}) error
}

func newHoster() hoster {
	return &Hoster{}
}

type Hoster struct {
}

func (p *Hoster) host(svcPtr interface{}) error {

	var (
		svcValue = reflect.ValueOf(svcPtr)
		// svcType  = svcValue.Type()
	)

	// Check if the given service is really a pointer.
	if svcValue.Kind() != reflect.Ptr {
		return noPointerErr
	}
	return nil
}

type indexOpeator interface {
	operate() string
}

type esIndexOperator struct {
}

func newEsIndexOperator() *esIndexOperator {
	return &esIndexOperator{}
}
