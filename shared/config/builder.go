package config

import (
	"errors"
	"fmt"
)

var (
	ErrNoSources = errors.New("no config sources provided")
)

type Builder struct {
	sources []Source
}

func NewBuidler() Builder {
	return Builder{
		sources: make([]Source, 0),
	}
}

func (b *Builder) AddSource(source Source) {
	b.sources = append(b.sources, source)
}

func (b *Builder) Build() error {
	if len(b.sources) == 0 {
		return ErrNoSources
	}

	for _, source := range b.sources {
		if len(source.Path) == 0 {
			return fmt.Errorf("%w: %s", ErrUnknownPath, source.Path)
		}

		switch source.Kind {
		case SourceKindEnv:

		default:
			return ErrUnknownSource
		}
	}

	return nil
}
