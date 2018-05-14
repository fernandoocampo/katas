package main

import (
	"fmt"

	"github.com/fernandoocampo/katas/brackets"
)

func main() {
	fmt.Println(string("[ABC[23]][89]"[0:]))
	fmt.Println(brackets.WhichIsTwo("[ABC[23]][89]", 0))
	fmt.Println(brackets.WhichIsTwo("[ABC[23]][89]", 4)) // 7
}
