package specs

import (
	"testing"
)

func TestMake(t *testing.T) {
	s := Make(1, 2, 3)

	assertEquals(t, 1, s.Head())

	s = s.Tail()

	assertEquals(t, 2, s.Head())

	s = s.Tail()

	assertEquals(t, 3, s.Head())

	s = s.Tail()

	assertTrue(t, s.Empty())

}

//func TestFromIntSlice(t *testing.T) {
//	slice := make([]int, 3)
//	slice[0] = 1
//	slice[1] = 2
//	slice[2] = 3
//
//	s := FromIntSlice(slice)
//
//	assertEquals(t, 1, s.Head())
//
//	s = s.Tail()
//
//	assertEquals(t, 2, s.Head())
//
//	s = s.Tail()
//
//	assertEquals(t, 3, s.Head())
//
//	s = s.Tail()
//
//	assertTrue(t, s.Empty())
//
//}

func TestIntRange(t *testing.T) {
	s := IntRange(1, 10)

	for i := 1; i <= 10; i++ {
		assertEquals(t, i, s.Head())

		s = s.Tail()
	}

	assertTrue(t, s.Empty())

}

func TestMap(t *testing.T) {
	s := IntRange(1, 10)
	s = s.Map(func(i interface{}) interface{} { return i.(int) + 10 })

	for i := 1; i <= 10; i++ {
		assertEquals(t, i+10, s.Head())

		s = s.Tail()
	}

	assertTrue(t, s.Empty())
}

func TestReduce(t *testing.T) {
	s := IntRange(1, 10)

	adder := func(a interface{}, b interface{}) interface{} {
		return (a.(int) + b.(int))
	}

	r := s.Reduce(adder)

	assertEquals(t, 55, r)

}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expecting %s: got %s ", expected, actual)
	}
}

func assertTrue(t *testing.T, expected bool) {
	if !expected {
		t.Errorf("True Expected : got False ")
	}
}
