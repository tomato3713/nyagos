package commands

import (
	"fmt"
)

type cmdHelp struct {
	usage string
	short string
	long  string
}

func (h cmdHelp) shortHelp() string {
	return h.short + "\n"
}
func (h cmdHelp) longHelp() string {
	help := fmt.Sprintf(`Usage: %s

    %s
`, h.usage, h.long)
	return help
}
