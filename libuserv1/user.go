package libuserv1

import "github.com/IlyaFloppy/importsbuglib"

func SomeFunc() importsbuglib.SomeType {
	typev1 := importsbuglib.SomeType{
		Value: "v1",
	}

	return typev1
}
