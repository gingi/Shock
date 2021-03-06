package filter

import (
	"github.com/MG-RAST/Shock/store/filter/anonymize"
	"github.com/MG-RAST/Shock/store/filter/fq2fa"
	"io"
)

type FilterFunc func(io.ReadCloser) io.ReadCloser

var (
	filters = map[string]FilterFunc{
		"anonymize": anonymize.NewReader,
		"fq2fa":     fq2fa.NewReader,
	}
)

func Has(f string) bool {
	if _, has := filters[f]; has {
		return true
	}
	return false
}

func Filter(f string) FilterFunc {
	return filters[f]
}

func NewReader(f string, file io.ReadCloser) io.ReadCloser {
	return filters[f](file)
}
