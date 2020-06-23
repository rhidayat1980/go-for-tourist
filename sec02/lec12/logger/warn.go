package logger

import "fmt"

// Warn functions
func Warn(s string) {
	fmt.Println("[WARN]", s)
}

func init() {
	fmt.Println("logger.init() - 1 from warn.go")
}
