package specs

// FnStrean is a function returning a stream.
type FnStream func() Stream

// Stream the main interface
type Stream interface {
	Head() interface{}
	Tail() Stream
	Empty() bool
	Map(func(interface{}) interface{}) Stream
	Reduce(func(interface{}, interface{}) interface{}, ...interface{}) interface{}
	Filter(func(interface{}) bool) Stream
}

type stream struct {
	first interface{}
	rest  FnStream
}

func (s *stream) Head() interface{} { // #generic #custom
	return s.first
}

func (s *stream) Tail() Stream { // #generic #custom
	return s.rest()
}

func (s *stream) Empty() bool { // #generic #custom
	return s.first == nil
}

func (s *stream) Map(f func(interface{}) interface{}) Stream { // #generic #custom
	if s.Empty() {
		return s
	}
	return &stream{f(s.Head()).(interface{}), func() Stream {
		return s.Tail().Map(f)
	}}
}

//func (s *stream) Reduce(agg func(interface{},interface{}) interface{}) interface{} { // #generic #custom
//
//}

func (s *stream) Reduce(agg func(interface{}, interface{}) interface{}, i ...interface{}) interface{} { //

	var initial interface{}
	var self Stream

	if len(i) < 1 {
		if s.Empty() {
			panic("aggregation is nil and stream is empty.")
		}
		initial = s.Head()
		self = s.Tail()
	} else {
		initial = i[0]
		self = s
	}

	if self.Empty() {
		return initial
	}

	return self.Tail().Reduce(agg, agg(initial, self.Head()))

}

func (s *stream) Filter(f func(interface{}) bool) Stream {
	if s.Empty() {
		return s
	}
	h := s.Head()
	t := s.Tail()
	if f(h) {
		return &stream{h, func() Stream {
			return t.Filter(f)
		}}
	}
	return t.Filter(f)
}

////////////////////////

// Make a stream from the input elements
func Make(elem ...interface{}) Stream { // #generic #custom
	if len(elem) == 0 {
		return &stream{nil, nil}
	}

	rest := elem[1:len(elem)]

	return &stream{elem[0].(interface{}), func() Stream {
		return Make(rest...)
	}}
}

// FromSlice build a stream from a slice
func FromIntSlice(slice []interface{}) Stream { // #int #custom
	if len(slice) == 0 {
		return &stream{nil, nil}
	}

	return &stream{slice[0], func() Stream {
		return FromIntSlice(slice[1:len(slice)])
	}}
}

// IntRange will return a stream containing the range [low - high]
func IntRange(low, high int) Stream { // #int

	if low == high {
		return Make(low)
	}

	return &stream{low, func() Stream {
		return IntRange(low+1, high)
	}}
}
