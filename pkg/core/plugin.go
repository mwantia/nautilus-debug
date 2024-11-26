package core

import (
	"github.com/mwantia/nautilus/pkg/log"
	"github.com/mwantia/nautilus/pkg/plugin"
)

type DebugProcessor struct {
	Config *DebugConfig
	Logger *log.Logger
}

func NewImpl() *DebugProcessor {
	return &DebugProcessor{
		Logger: log.NewLogger("debug"),
	}
}

func (p *DebugProcessor) Name() (string, error) {
	return "debug", nil
}

func (p *DebugProcessor) GetCapabilities() (plugin.PipelineProcessorCapability, error) {
	return plugin.PipelineProcessorCapability{
		Types: []plugin.PipelineProcessorCapabilityType{
			plugin.None,
		},
	}, nil
}

func (p *DebugProcessor) Configure(data map[string]interface{}) error {
	cfg, err := GetConfig(data)
	if err != nil {
		p.Logger.Warn("Error converting mapstructure", "error", err)
	}

	p.Config = cfg

	p.Logger.Info("Configuring external debug plugin...")
	p.Logger.Debug("External Debug plugin", "name", cfg.Name)

	return nil
}

func (p *DebugProcessor) Process(data *plugin.PipelineContextData) (*plugin.PipelineContextData, error) {
	p.Logger.Info("Processing external debug plugin...")
	return data, nil
}

func (p *DebugProcessor) Health() error {
	p.Logger.Info("Checking health for external debug plugin...")
	return nil
}
