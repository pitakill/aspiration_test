package main

import (
	"fmt"

	"github.com/pitakill/test/mapper"
)

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
