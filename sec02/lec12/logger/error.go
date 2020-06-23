package logger

import "fmt"

// Error functions
func Error(s string) {
	fmt.Println("[ERROR]", s)
}

func init() {
	fmt.Println("logger.init() - from error.go")
}
