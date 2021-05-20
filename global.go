package logger

import (
	"sync"

	"github.com/nori-io/common/v5/pkg/domain/logger"
)

var (
	_globalMu sync.RWMutex
	_globalL  = New()
)

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
// It's safe for concurrent use.
func L() logger.Logger {
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}

// ReplaceGlobals replaces the global Logger and SugaredLogger, and returns a
// function to restore the original values. It's safe for concurrent use.
func ReplaceGlobals(logger logger.Logger) func() {
	_globalMu.Lock()
	prev := _globalL
	_globalL = logger
	_globalMu.Unlock()
	return func() { ReplaceGlobals(prev) }
}
