package gimme

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/text"
	"golang.org/x/term"
)

func Header(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err) // TODO
	}

	divider := text.Pad("", width-2, '=')
	fmt.Println(divider)
	fmt.Println(text.AlignCenter.Apply(msg, width))
	fmt.Println(divider)
}

func Section(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err) // TODO
	}

	var pad string
	sideSize := ((width - len(msg)) / 2) - 4
	if sideSize >= 0 {
		pad = strings.Repeat("-", sideSize)
	}
	msg = fmt.Sprintf("%s %s %s", pad, msg, pad)

	fmt.Println(text.AlignCenter.Apply(msg, width))
}

func Error(msg string, args ...any) {
	if text.ANSICodesSupported {

	} else {
		fmt.Printf("ERROR: "+msg+"\n", args...)
	}
}
