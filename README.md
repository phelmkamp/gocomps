[![Go Reference](https://pkg.go.dev/badge/github.com/phelmkamp/gocomps.svg)](https://pkg.go.dev/github.com/phelmkamp/gocomps)
[![Go Report Card](https://goreportcard.com/badge/github.com/phelmkamp/gocomps)](https://goreportcard.com/report/github.com/phelmkamp/gocomps)


# gocomps
A component-based framework for Go

## Examples

### Simple

See [testdata/main.go](testdata/main.go).

```go
func root(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("root", props)
	return component.NewGroup(
		component.New(a, map[string]any{"name": "a"}),
		component.New(b, map[string]any{"name": "b"}),
	)
}

func a(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("a", props)
	return component.NewGroup(
		component.New(aa, map[string]any{"name": "aa"}),
		component.New(ab, map[string]any{"name": "ab"}),
	)
}

func b(ctx context.Context, props map[string]any) component.Component {
	fmt.Println("b", props)
	return component.New(ba, map[string]any{"name": "ba"})
}

func main() {
	root := component.New(root, map[string]any{"name": "root"})
	component.Run(context.Background(), root)
}
```

### HTTP Handler

See [testdata/http/main.go](testdata/http/main.go).

```go
func handle(ctx context.Context, props handler.Props) component.Component {
	name := props.R.URL.Query().Get("name")
	onGreet := func(s string) {
		props.W.Write([]byte(s))
	}
	return component.New(greetSvc, crud.NewWriteProps(name, onGreet, nil))
}

func main() {
	http.ListenAndServe(":8080", handler.New(handle))
}
```
