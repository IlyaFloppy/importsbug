package libuseralias

import importsbuglibalias "github.com/IlyaFloppy/importsbuglib"

func SomeFunc() importsbuglibalias.SomeType {
	typealias := importsbuglibalias.SomeType{
		Value: "alias",
	}

	return typealias
}
