// Learning about type assertions
// type assertion allows to access the concrete value under an interface
// The below statement asserts that interface "i" holds a value of type "T"
// v := i.(T)
//
// If the assertion is true, "v" will be assigned the concrete value of "i"
// If the assertion is false, this will cause a panic.
//
// This can be handled more gracefully when testing for specific types
// v, ok := i(.T)
// if "i" does not hold a value of type "T", it would return (nil value of "v", false)

package main

import "fmt"

// some sample interface and types to play with

type Quaker interface {
	Quack()
}

type Duck struct{}

func (d Duck) Quack() {
	fmt.Println("Quack Quack!!!")
}

type Chicken struct{}

func (c Chicken) String() {
	fmt.Println("Buck Buck!!!")
}

func main() {
	var i interface{}

	val := "gopher"
	i = val

	// v will refer to concrete value of the interface "i"
	// since val is of type string. Does the address of the
	// value pointed by v and val are same?
	v := i.(string)
	fmt.Printf("v: %v\n", v)

	// values are assigned based on the standard assignment rules
	// for simple types, value seem to be copied.
	// v & val point to different address (I think it has to. Hasn't it?)
	fmt.Printf("&v: %p, &val:%p\n", &v, &val)

	b, ok := i.(bool)
	fmt.Printf("b: %v, ok: %v\n", b, ok)

	f, ok := i.(float64)
	fmt.Printf("f: %v, ok: %v\n", f, ok)

	primes := []int{1, 2, 3, 5, 7}
	i = primes

	s, ok := i.([]int)
	fmt.Printf("s: %v, ok: %v\n", s, ok)

	// s & primes point to the same address
	fmt.Printf("&s: %p, &primes: %p\n", s, primes)

	// further operations can diverge though.
	s = append(s, 9)
	fmt.Printf("&s: %p, &primes: %p\n", s, primes)
	fmt.Printf("s: %v, primes: %v\n", s, primes)

	i = primes[4]

	// type assertion has a special format when used in switch case
	// i.(type) can be used in the switch statement to act based
	// on the type of the interface.
	format := func(a interface{}) string {
		switch v := a.(type) {
		case int:
			return fmt.Sprintf("int: v=%d", v)
		case float64:
			return fmt.Sprintf("float64: v=%f", v)
		case error:
			return fmt.Sprintf("error: v=%v", v)
		case Duck:
			return fmt.Sprintf("Duck: v=%v", v)
		case fmt.Stringer:
			return fmt.Sprintf("stringer: v=%s", v)
		default:
			return fmt.Sprintf("unknown v=%v, type=%[1]T", v)
		}
	}

	fmt.Println(format(45))

	mixedTypes := make(map[string]interface{})

	mixedTypes["string"] = "foo"
	mixedTypes["float64"] = 2484.22
	// interesting that error strings are not allowed to be capitalized
	// go-staticcheck complains otherwise.
	mixedTypes["error"] = fmt.Errorf("test error")
	mixedTypes["Duck"] = Duck{}
	mixedTypes["Stringer"] = Chicken{}

	for k := range mixedTypes {
		fmt.Println(format(mixedTypes[k]))
	}

	// type assertions are used extensively in Go
	// fmt/print.go has lot of examples of type assertion
	// fmt.Print(...) takes variable number of arguments of type "interface{}"
	// the type of the argument is then checked internally using type assertion
	// and handled accordingly.
	//
	// some example of type assertion in switch case
	// .. from go-safeweb package
	// https://github.com/google/go-safeweb/blob/d4d7197bce5996f2440b9fa9559990b1c8ae8bdf/safehttp/form.go#L107
	// https://github.com/google/go-safeweb/blob/d4d7197bce5996f2440b9fa9559990b1c8ae8bdf/safehttp/form.go#L107
}
