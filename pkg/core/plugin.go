package core

import (
	"log"

	"github.com/mwantia/nautilus/pkg/shared"
)

type DebugProcessor struct {
}

func NewImpl() *DebugProcessor {
	return &DebugProcessor{}
}

func (p *DebugProcessor) Name() (string, error) {
	return "debug", nil
}

func (p *DebugProcessor) Process(ctx *shared.NautilusPipelineContext) (*shared.NautilusPipelineContext, error) {
	log.Println("Processing external debug plugin...")
	return ctx, nil
}

func (p *DebugProcessor) Configure() error {
	log.Println("Configuring external debug plugin...")
	return nil
}

func (p *DebugProcessor) Health() error {
	log.Println("Checking health for external debug plugin...")
	return nil
}
