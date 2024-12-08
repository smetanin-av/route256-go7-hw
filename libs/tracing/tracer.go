package tracing

import (
	"github.com/uber/jaeger-client-go/config"
)

func InitGlobal(service string) error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(service)
	if err != nil {
		return err
	}

	return nil
}
