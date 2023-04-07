package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/phelmkamp/gocomps/component"
)

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

func input(ctx context.Context, onChange func(string)) component.Component {
	go func() {
		sc := bufio.NewScanner(os.Stdin)
		sc.Split(bufio.ScanLines)
		for sc.Scan() {
			onChange(sc.Text())
			return
		}
	}()
	return component.Component{}
}

func main() {
	fmt.Println("Enter email:")
	done := make(chan struct{})
	component.Run(context.Background(), component.New(read, done))
	<-done
}
