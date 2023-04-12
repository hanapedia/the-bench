package factory

import (
	"github.com/hanapedia/the-bench/service-unit/internal/domain/core"
	"github.com/hanapedia/the-bench/service-unit/internal/infrastructure/server_adapter/rest"
)

func RestServerAdapterFactory(serverAdapters *[]*core.ServerAdapter, handler *core.Handler) {
	idx := -1
	for i, serverAdapter := range *serverAdapters {
		if restServerAdapter, ok := (*serverAdapter).(rest.RestServerAdapter); ok {
			restServerAdapter.Register(handler)
			idx = i
			break
		}
	}

	if idx < 0 {
		restServerAdapter := rest.NewRestServerAdapter()
		restServerAdapter.Register(handler)
	}
}
