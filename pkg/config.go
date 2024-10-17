package pkg

import (
	"flag"
)

func GetPortSizes() (map[string]int, error) {
	in := flag.Int("in", 10, "input port size")
	out := flag.Int("out", 10, "output port size")

	flag.Parse()
	if *in < 0 || *out < 0 {
		panic("port size must be positive")
	}
	return map[string]int{
		"IN":  *in,
		"OUT": *out,
	}, nil
}
