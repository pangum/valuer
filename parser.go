package starter

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
	"github.com/goexl/valuer"
)

// Parser 解析器
type Parser = valuer.Parser

func newParser(http *resty.Client, logger simaqian.Logger) (parser *Parser) {
	builder := valuer.New()
	builder.Logger(logger)
	builder.Http(http)
	parser = builder.Build()

	return
}
