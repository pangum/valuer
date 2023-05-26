package valuer

import (
	"github.com/goexl/valuer"
	"github.com/pangum/http"
	"github.com/pangum/logging"
)

// Parser 解析器
type Parser = valuer.Parser

func newParser(http *http.Client, logger logging.Logger) (parser *Parser) {
	builder := valuer.New()
	builder.Logger(logger)
	builder.Http(http.Client)
	parser = builder.Build()

	return
}
