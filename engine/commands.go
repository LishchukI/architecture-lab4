package engine

import (
	"fmt"
	"strings"
)

type PrintcCommand struct {
	Str string
	Sep string
}

type Command interface {
	Execute(handler Handler)
}

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop Handler) {
	fmt.Println(p.Arg)
}

func (printcArgs *PrintcCommand) Execute(loop Handler) {
		res := strings.Split(printcArgs.Str, printcArgs.Sep)
		for i := 0; i < len(res); i++{
		loop.Post(&PrintCommand{Arg: res[i]})
	}

}

type finishCommand func (handler Handler)

func (c finishCommand) Execute(handler Handler) {
	c(handler)
}
