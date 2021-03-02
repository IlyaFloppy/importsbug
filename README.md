# importsbug

Repository contains minimal example of how `golang.org/x/tools/imports` chooses wrong import paths when v2 module is used.

`imports.Process` always substitutes `import "github.com/IlyaFloppy/importsbuglib"` even though `libuserv2/user.go` imports `"github.com/IlyaFloppy/importsbuglib/v2"`.

Furthermore aliased imports get removed.

Processed `libuserv1/user.go`:
```golang
package libuserv1

import "github.com/IlyaFloppy/importsbuglib"

func SomeFunc() importsbuglib.SomeType {
        typev1 := importsbuglib.SomeType{
                Value: "v1",
        }

        return typev1
}
```

Processed `libuserv2/user.go`:
```golang
package libuserv2

import "github.com/IlyaFloppy/importsbuglib"

func SomeFunc() importsbuglib.SomeType {
        typev2 := importsbuglib.SomeType{
                Value:        "v2",
                AnotherValue: "v2",
        }

        return typev2
}
```

Processed `libuseralias/user.go`:
```golang
package libuseralias

func SomeFunc() importsbuglibalias.SomeType {
        typealias := importsbuglibalias.SomeType{
                Value: "alias",
        }

        return typealias
}

```
Run `go run main.go` to reproduce.
