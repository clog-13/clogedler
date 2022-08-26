package middleware

import (
	"clogedler/framework"
	"fmt"
)

func Test1() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}
