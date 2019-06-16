package main

import (
	"github.com/pickme-go/errors"
	"log"
)

func main() {

	err1 := errors.New(`func1`, `something went wrong on func 1`)
	err2 := errors.WithPrevious(err1, `func2`, `something went wrong on func 2`)
	err3 := errors.WithPrevious(err2, `func3`, `something went wrong on func 3`)
	err4 := errors.WithPrevious(err3, `func4`, `something went wrong on func 4`)

	log.Printf(`error happens due to %s`, err4.Error())
}
