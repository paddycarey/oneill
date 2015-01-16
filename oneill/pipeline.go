package oneill

import (
	"github.com/rehabstudio/oneill/logger"
)

type Processor func([]*SiteConfig) []*SiteConfig

type ProcessorPipeline struct {
	processors []Processor
}

func (p *ProcessorPipeline) AddProcessor(processor Processor) {
	p.processors = append(p.processors, processor)
}

func (p *ProcessorPipeline) RunPipeline() []*SiteConfig {
	logger.LogDebug("Running processor pipeline")
	var siteDefinitions []*SiteConfig
	for _, processor := range p.processors {
		siteDefinitions = processor(siteDefinitions)
	}
	return siteDefinitions
}
