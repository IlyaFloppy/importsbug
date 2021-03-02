package libuserv2plus

import "github.com/IlyaFloppy/importsbuglib/v2"

func SomeFunc() importsbuglib.SomeType {
	typev2 := importsbuglib.SomeType{
		Value:        "v2plus",
		AnotherValue: "v2plus",
	}

	return typev2
}
