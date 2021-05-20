package main

import (
	"bytes"
	"context"
	"os"

	"github.com/nori-io/common/v5/pkg/domain/config"
	enum "github.com/nori-io/common/v5/pkg/domain/enum/meta"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-io/common/v5/pkg/domain/meta"
	"github.com/nori-io/common/v5/pkg/domain/plugin"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	metadata "github.com/nori-io/common/v5/pkg/meta"
)

type service struct {
	instance *instance
}

type instance struct {
	Path   config.String
	writer *os.File
}

var (
	Plugin plugin.Plugin = &service{}
)

func (p *service) Init(_ context.Context, config config.Config, log logger.FieldLogger) error {
	p.instance = &instance{}
	p.instance.Path = config.String("hooks.filehook.path", "file to collect log messages")
	return nil
}

func (p *service) Instance() interface{} {
	return p.instance
}

func (p service) Meta() meta.Meta {
	return &metadata.Meta{
		ID: metadata.ID{
			ID:      "nori/logger/Hook",
			Version: "1.0.0",
		},
		Author: metadata.Author{
			Name: "Nori.io",
			URL:  "https://nori.io/",
		},
		Dependencies: []meta.Dependency{},
		Description: metadata.Description{
			Title:       "Test File Hook Plugin",
			Description: "This is a tets plugin for development and testing purpose",
		},
		Interface: logger.HookInterface,
		License: []meta.License{
			metadata.License{
				Title: "GPLv3",
				Type:  enum.GPLv3,
				URL:   "https://www.gnu.org/licenses/",
			},
		},
		Tags: []string{"hook", "file_hook"},
	}
}

func (p service) Start(ctx context.Context, registry registry.Registry) error {
	var err error
	p.instance.writer, err = os.OpenFile(p.instance.Path(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func (p service) Stop(ctx context.Context, registry registry.Registry) error {
	return p.instance.writer.Close()
}

// Hook
func (i *instance) Levels() []logger.Level {
	return []logger.Level{logger.LevelFatal, logger.LevelPanic, logger.LevelNotice, logger.LevelCritical, logger.LevelError,
		logger.LevelWarning, logger.LevelInfo}
}

func (i *instance) Fire(e logger.Entry, field ...logger.Field) error {
	if e.Level == logger.LevelDebug {
		return nil
	}

	b := bytes.Buffer{}
	out, _ := e.Formatter.Format(e, field...)
	b.Write(out)
	_, err := i.writer.Write(b.Bytes())
	return err
}
