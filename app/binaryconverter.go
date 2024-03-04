package app

import (
	"math"

	"github.com/adityatresnobudi/bin2dec/entity"
	"github.com/adityatresnobudi/bin2dec/shared"
)

type Bin2DecConverter struct {
	binary  *entity.Input
	decimal *int
}

func NewBin2DecConverter(binary *entity.Input) *Bin2DecConverter {
	return &Bin2DecConverter{
		binary:  binary,
		decimal: nil,
	}
}

func (bd *Bin2DecConverter) Convert() (*int, error) {
	res := 0
	length := len(bd.binary.Binary)
	for idx, bin := range bd.binary.Binary {
		if bin == shared.One {
			res += int(math.Pow(2, float64(length-1-idx)))
		}
	}

	bd.decimal = &res

	return bd.decimal, nil
}

func (bd *Bin2DecConverter) IsValid() (bool, error) {
	if len(bd.binary.Binary) > 8 {
		return false, shared.ErrMoreThanEight
	}

	for _, bin := range bd.binary.Binary {
		if bin == shared.One || bin == shared.Zero {
			continue
		}
		return false, shared.ErrInvalidBinary
	}

	return true, nil
}
