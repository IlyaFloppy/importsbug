package libuserv2

import "github.com/IlyaFloppy/importsbuglib/v2"

func SomeFunc() importsbuglib.SomeType {
	typev2 := importsbuglib.SomeType{
		Value:        "v2",
		AnotherValue: "v2",
	}

	return typev2
}
