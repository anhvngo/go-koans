package go_koans

func aboutBasics() {
	__bool__ := true
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // in it there is nothing false

	var i int = 1
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	assert(5%2 == 1)
	assert(5*2 == 10)
	assert(5^2 == 7)

	var x int
	assert(x == 0) // zero values are valued in Go

	var f float32
	assert(f == 0) // for types of all types

	var s string
	assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == __int__)     // and types within composite types
	assert(c.f == __float32__) // which match the other types
	assert(c.s == __string__)  // in a typical way
}
