[![Go Reference](https://pkg.go.dev/badge/github.com/phelmkamp/gocomps.svg)](https://pkg.go.dev/github.com/phelmkamp/gocomps)
[![Go Report Card](https://goreportcard.com/badge/github.com/phelmkamp/gocomps)](https://goreportcard.com/report/github.com/phelmkamp/gocomps)


# gocomps
A component-based framework for Go

## Examples

### Simple

See [examples/simple/main.go](examples/simple/main.go).

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

See [examples/http/main.go](examples/http/main.go).

```go
func handle(ctx context.Context, props handler.Props) component.Component {
	name := props.R.URL.Query().Get("name")
	onGreet := func(msg string, err error) {
		if err != nil {
			props.W.WriteHeader(http.StatusBadRequest)
			props.W.Write([]byte(err.Error()))
			return
		}
		props.W.Write([]byte(msg))
	}
	return component.New(greetSvc, greetProps{name: name, greet: onGreet})
}

func main() {
	http.ListenAndServe(":8080", handler.New(handle))
}
```

### Parallel

See [examples/parallel/main.go](examples/parallel/main.go).

```go
func sum(ctx context.Context, p props) component.Component {
	if len(p.in) <= p.batchSize {
		return apply(ctx, p)
	}
	return combine(ctx, p)
}

func split(ctx context.Context, p props) component.Component {
	return component.NewGroup(
		component.New(sum, props{in: p.in[:len(p.in)/2], out: p.out, batchSize: p.batchSize}),
		component.New(sum, props{in: p.in[len(p.in)/2:], out: p.out, batchSize: p.batchSize}),
	)
}

func apply(ctx context.Context, p props) component.Component {
	var n int
	for _, v := range p.in {
		n += v
	}
	p.out <- n
	return component.Component{}
}

func combine(ctx context.Context, p props) component.Component {
	res := make(chan int)
	go func() {
		var n int
		for i := 0; i < 2; i++ {
			n += <-res
		}
		p.out <- n
	}()
	return component.New(split, props{in: p.in, out: res, batchSize: p.batchSize})
}
```

### State

See [examples/state/main.go](examples/state/main.go).

```go
func read(ctx context.Context, done chan struct{}) component.Component {
	type state struct {
		email string
		valid bool
	}
	emailState, setEmailState := component.UseState(ctx, state{})

	if !emailState.valid {
		if len(emailState.email) > 0 {
			fmt.Println("invalid input:", emailState.email)
		}

		onChangeEmail := func(s string) {
			setEmailState(state{email: s, valid: strings.Contains(s, "@")})
		}
		return component.New(input, onChangeEmail)
	}

	done <- struct{}{}
	return component.Component{}
}
```

### Reducer

See [examples/reducer/main.go](examples/reducer/main.go).

```go
func reduce(i int, action string) int {
	switch action {
	case "add":
		return i + 1
	case "subtract":
		return i - 1
	default:
		return i
	}
}

func read(ctx context.Context, done chan struct{}) component.Component {
	i, dispatch := component.UseReducer(ctx, reduce, 0)

	fmt.Println(i)

	onChange := func(s string) {
		if s == "+" {
			dispatch("add")
		} else if s == "-" {
			dispatch("subtract")
		} else {
			done <- struct{}{}
		}
	}
	return component.New(input, onChange)
}
```