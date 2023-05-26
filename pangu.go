package valuer

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newParser,
	)
}
