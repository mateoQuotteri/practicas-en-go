package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//genero un contexto-
	ctx := context.Background()

	//le agrego valores
	ctx = context.WithValue(ctx, "key", "value")
	ctx = context.WithValue(ctx, "key-int", 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	select {
	case <-ctx2.Done():
		fmt.Println("Contexto cancelado")
		fmt.Println("Error:", ctx2.Err())
	}
}

func viewContext(ctx context.Context) {
	fmt.Println("My value is %s\n", ctx.Value("key"))
	fmt.Println("My int value is %d\n", ctx.Value("key-int"))
}
