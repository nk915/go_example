package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {

	g, _ := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		i := i
		fmt.Printf("[%d] ready..\n", i)
		g.Go(func() error {
			fmt.Printf("[%d] start step...\n", i)
			if err := printIndex(i); err != nil {
				return err
			}

			fmt.Printf("[%d] end step...\n", i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Err : %v", err)
	}
}

func printIndex(n int) error {
	if n == 3 {
		fmt.Printf("[%d] goroutine err\n", n)
		return errors.New("invalid index")
	}

	fmt.Printf("[%d] ing.. \n", n)
	return nil
}
