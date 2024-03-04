package cli

import (
	"github.com/adityatresnobudi/bin2dec/app"
	"github.com/adityatresnobudi/bin2dec/entity"
)

func HandleConvert(input string) (int, error) {
	binConverter := app.NewBin2DecConverter(&entity.Input{Binary: input})
	res, err := binConverter.Convert()
	return *res, err
}
