package proja

import (
	"fmt"
	c "github.com/w3xse7en/third-proj-c"
)

func RpcA(i int) string {
	return fmt.Sprintf("im rpc a input int is %v", c.ToolCItoa(i))
}
