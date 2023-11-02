package valuer

import (
	"github.com/pangum/pangu"
	"github.com/pangum/valuer/internal/plugin"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.New,
	).Build().Build().Apply()
}
