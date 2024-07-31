package main

import (
	"context"
	"fmt"
)

type MiddleWares []MiddleWare
type MiddleWare interface {
	Fire(ctx *context.Context, msg string) error
}

type Prefix struct {
}

func NewPrefix() *Prefix {
	return &Prefix{}
}
func (p *Prefix) Fire(ctx *context.Context, msg string) error {
	fmt.Printf("Before Fire (ctx : %v %v, msg : %s)\n", ctx, *ctx, msg)
	*ctx = context.WithValue(*ctx, "key", "value")
	fmt.Printf("After Fire (ctx : %v %v, msg : %s)\n", ctx, *ctx, msg)
	return nil
}

func main() {
	hooks := MiddleWares{}
	hooks = append(hooks, NewPrefix())

	ctx := context.Background()
	for idx, hook := range hooks {
		fmt.Println("loop ", idx)
		hook.Fire(&ctx, "input")
	}

	fmt.Printf("End (ctx : %v, %v)", &ctx, ctx)
}
