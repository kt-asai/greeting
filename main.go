package greeting

import "fmt"

func Greet(name string) string {
	msg := fmt.Sprintf("Hello %s\n", name)

	return msg
}
