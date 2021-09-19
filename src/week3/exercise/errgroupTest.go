package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

//  go get golang.org/x/sync/errgroup
func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		// test1函数还可以在启动很多goroutine
		// 子节点都传入ctx，当test1报错，会把test1的子节点一一cancel
		return test1(ctx)
	})

	eg.Go(func() error {
		return test1(ctx)
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func test1(ctx context.Context) error {
	return errors.New("test2")
}
