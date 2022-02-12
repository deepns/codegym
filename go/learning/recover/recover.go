// Exploring recover keyword and its usage
package main

import "fmt"

func div(a, b float64) float64 {
	fmt.Printf("a: %v, b:%v\n", a, b)

	// if the divisor is 0, go division operation returns Inf
	// no runtime error.
	// just panicking to see if this is captured in recover
	if b == 0 {
		panic(fmt.Sprintf("Invalid divisor %v\n", b))
	}

	return a / b
}

func main() {
	// recover allows to recover from panicking goroutine.
	// panic...recover is vaguely similar to try...catch.
	// recover returns the err that caused the panic, that
	// is typically the value passed to panic()
	// note: arg to panic is an interface. (panic(v interface{}))
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic. r: %v", r)
		}
	}()

	// when panic is not hit or if recover() is used outside
	// a deferred function, recover() returns nil.
	r := recover()
	if r == nil {
		fmt.Println("Just called recover() outside a deferred function")
	}

	div(5, 3)
	div(1, 0)

	// This is unreached. Since div(1, 0) would hit a panic
	// and the recover() in deferred function handles it.
	div(5, 43)

	// Looking at the usage of recover() in the go src, it is used pretty
	// commonly. one example, from the print functions in src/fmt/print.go
	//
	// switch v := p.arg.(type) {
	// case error:
	// 	handled = true
	// 	defer p.catchPanic(p.arg, verb, "Error")
	// 	p.fmtString(v.Error(), verb)
	// 	return

	// case Stringer:
	// 	handled = true
	// 	defer p.catchPanic(p.arg, verb, "String")
	// 	p.fmtString(v.String(), verb)
	// 	return
	// }
	//
	// catchPanic handles the error
	// func (p *pp) catchPanic(arg interface{}, verb rune, method string) {
	// 	if err := recover(); err != nil {
	// }

	// another common usecase for recover() is in test
	// when a testcase panics, the testcase can handle it in a deferred
	// function and check whether the panic is an expected one, and has
	// the expected error.

	// another example for recover(), from src/bufio/bufio_test.go
	//
	// type negativeReader int

	// func (r *negativeReader) Read([]byte) (int, error) { return -1, nil }

	// func TestNegativeRead(t *testing.T) {
	// 	// should panic with a description pointing at the reader, not at itself.
	// 	// (should NOT panic with slice index error, for example.)
	// 	b := NewReader(new(negativeReader))
	// 	defer func() {
	// 		switch err := recover().(type) {
	// 		case nil:
	// 			t.Fatal("read did not panic")
	// 		case error:
	// 			if !strings.Contains(err.Error(), "reader returned negative count from Read") {
	// 				t.Fatalf("wrong panic: %v", err)
	// 			}
	// 		default:
	// 			t.Fatalf("unexpected panic value: %T(%v)", err, err)
	// 		}
	// 	}()
	// 	b.Read(make([]byte, 100))
	// }

	// some examples from kubernetes src
	// https://github.com/kubernetes/kubernetes/blob/3726309bf9d59bccf28b9e22e1573764a5dd3fb5/pkg/volume/util/types/types.go#L77
	// https://github.com/kubernetes/kubernetes/blob/f90267f0629e546d1c5217c19138e2e8f68bfa92/pkg/util/flag/flags_test.go#L59
}
