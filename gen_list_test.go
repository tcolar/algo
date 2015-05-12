// History: Jan 04 14 tcolar Creation

package algo

import (
	"log"
	"reflect"
	"testing"
)

type GSlice struct {
	slice  []interface{}
	sliceV reflect.Value // value of pointer to slice
}

func New() *GSlice {
	s := &GSlice{}
	s.slice = []interface{}{}
	s.sliceV = reflect.ValueOf(&s.slice)
	return s
}

func (s *GSlice) Second(receiver interface{}) {
	obj := reflect.ValueOf(receiver).Elem()
	obj.Set(reflect.Indirect(s.sliceV).Index(1).Elem())
}

func TestGList(t *testing.T) {
	//stuff := []int{5, 6, 7}
	ml := New()
	ml.slice = append(ml.slice, 7)
	ml.slice = append(ml.slice, 9)
	ml.slice = append(ml.slice, "b")
	//log.Print(ml.slice[0].(int) + 1)

	var result int
	ml.Second(&result)
	log.Print(result)
	//ml.FirstTwo(&result)
	//log.Print(result)

	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	log.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	a := 3.4
	b := 5.5
	log.Println(floatSwap(a, b))
	log.Print(5.0 + a)
}
