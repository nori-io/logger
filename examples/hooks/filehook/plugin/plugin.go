package main

import (
	"bytes"
	"context"
	"os"

	"github.com/nori-io/common/v4/pkg/domain/config"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/plugin"
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
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/logger/Hook",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori.io",
			URI:  "https://nori.io/",
		},
		Core: meta.Core{
			VersionConstraint: "=0.2.0",
		},
		Dependencies: []meta.Dependency{},
		Description: meta.Description{
			Name:        "Test File Hook Plugin",
			Description: "This is a tets plugin for development and testing purpose",
		},
		Interface: logger.HookInterface,
		License: []meta.License{
			{
				Title: "GPLv3",
				Type:  "GPLv3",
				URI:   "https://www.gnu.org/licenses/",
			},
		},
		Tags: []string{"hook", "file_hook"},
	}
}

func (p service) Start(ctx context.Context, registry plugin.Registry) error {
	var err error
	p.instance.writer, err = os.OpenFile(p.instance.Path(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func (p service) Stop(ctx context.Context, registry plugin.Registry) error {
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
