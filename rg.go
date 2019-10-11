package brdocs

import (
	"errors"
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"log"
	"strconv"
	"strings"
)

const RgType internal.DocumentType = "RG"

const (
	rgNumberOfDigits int    = 1
	rgValidatorRegex string = `^([\d]{7})([\d]{1})$`
	rgFormatterRegex string = "$1-$2"
)

type rg struct{}

func NewRg(number string) (internal.Document, error) {
	if len(number) < 8 {
		return internal.Document{}, errors.New("invalid number")
	}

	d, err := internal.NewDocument(
		RgType,
		number,
		len(number),
		rgNumberOfDigits,
		rgFormatterRegex,
		rgValidatorRegex,
		rg{},
	)
	return d, err
}

func (d rg) CalculateDigit(n string) string {
	sn := strings.Split(n, "")
	dv := sn[len(n)]
	s := make([]int, len(n) - 1)
	for i := len(n) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(sn[i])
		if err != nil {
			log.Fatalf("could not convert %s into int", sn[i]);
		}

		s[i] = num
	}
	return fmt.Sprintf("%d", dv)
}