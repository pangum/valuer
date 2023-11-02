package plugin

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/valuer"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(http *http.Client, logger log.Logger) (parser *valuer.Parser) {
	builder := valuer.New()
	builder.Logger(logger)
	builder.Http(http)
	parser = builder.Build()

	return
}
