package structs

import "fmt"

// Learning struct

// structs are yet another type in golang

// Version of the application
type Version struct {
	// just like the type, the struct fields also supports
	// exporting. fields starting in uppercase are exported.
	// Only exported fields can be accessed outside the struct.
	Major int
	Minor int
	name  string
}

// UpgradeMajorVersionByValue updates the major version.
// updated value not reflected in the original struct as it
// passed by value
func UpgradeMajorVersionByValue(v Version) {
	v.Major++
}

// UpgradeMajorVersion updates the major version
func UpgradeMajorVersion(v *Version) {
	v.Major++
}

func Learn() {
	fmt.Println("========== Learning Structs ==========")
	v1 := Version{1, 0, "apple"}
	fmt.Printf("v1: %v, type(v1):%T\n", v1, v1)

	// updating the fields of struct
	v1.name = "banana"
	v1.Major, v1.Minor = 2, 1
	fmt.Printf("v1: %v\n", v1)

	// structs can also be accessed using pointers
	versionPtr := &v1
	fmt.Printf("versionPtr: %p, *versionPtr:%v\n", versionPtr, *versionPtr)

	// update a struct using pointer
	// unlike C, pointers also access the fields with . notation.
	// bye bye -> operator
	versionPtr.Major = 3
	fmt.Printf("versionPtr: %p, *versionPtr:%v\n", versionPtr, *versionPtr)

	// when passing structs to functions, copies are created
	// unless they are passed as pointers
	fmt.Printf("v1: %v\n", v1)
	UpgradeMajorVersionByValue(v1)
	fmt.Printf("v1: %v\n", v1)

	UpgradeMajorVersion(&v1)
	fmt.Printf("v1: %v\n", v1)
}
