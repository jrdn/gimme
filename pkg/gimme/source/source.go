package source

import (
	"errors"

	"github.com/jrdn/gimme/pkg/gimme/config"
)

var ErrSourcePullFailed = errors.New("failed to pull source")

type SourceStatus string

type Source interface {
	Status() SourceStatus
	Pull() error
	GetConfig() (*config.Config, error)
	GetInstallDir() string
}
