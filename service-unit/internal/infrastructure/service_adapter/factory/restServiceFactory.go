package factory

import (
	"errors"
	"fmt"

	"github.com/hanapedia/the-bench/service-unit/internal/domain/core"
	"github.com/hanapedia/the-bench/service-unit/internal/infrastructure/service_adapter/rest"
)

func (serviceAdapterDetails ServiceAdapterDetails) RestServiceAdapterFactory() (core.ServiceAdapter, error) {
	var err error = nil
	var serviceAdapter core.ServiceAdapter
	switch serviceAdapterDetails.action {
	case "read":
    serviceAdapter = rest.RestReadAdapter{URL: fmt.Sprintf("http://%s:8080", serviceAdapterDetails.serviceName)}
	case "write":
		serviceAdapter = rest.RestWriteAdapter{URL: fmt.Sprintf("http://%s:8080", serviceAdapterDetails.serviceName)}
	default:
		err = errors.New("No matching protocol found")
	}
	return serviceAdapter, err
}
