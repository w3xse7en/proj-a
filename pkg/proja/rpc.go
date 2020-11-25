package proja

import (
	"fmt"
	"github.com/w3xse7en/proj-b/pkg/projb"
	c "github.com/w3xse7en/third-proj-c"
)

func RpcA(i int) string {
	return fmt.Sprintf("im rpc a input int is %v", c.ToolCItoa(i))
}

func AccessProjB(i int) string {
	return projb.RpcB(i)
}
