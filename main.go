/*
Copyright © 2026 John-Evans Wagenaar <jewagenaar@gmail.com>
*/
package main

import (
	"context"
	"ue5cli/cmd"
)

func main() {
	ctx := context.Background()
	cmd.ExecuteContext(ctx)
}
