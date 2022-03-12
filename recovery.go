package axisapi

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

/*
 Author：AxisZql
 Date: 2022-3-12
*/

// 获取触发panic的堆栈信息
func trace(msg string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) //skip first 3 caller
	var str strings.Builder
	str.WriteString(msg + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc) //获取引发panic的每个原文件具体位置
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandleFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(msg))
			}
		}()
		ctx.Next()
	}
}
