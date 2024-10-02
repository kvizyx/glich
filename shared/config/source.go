package config

import "errors"

type SourceKind int

const (
	SourceKindEnv SourceKind = iota
)

var (
	ErrUnknownSource = errors.New("unknown config source")
	ErrUnknownPath   = errors.New("unknown config path")
)

type Source struct {
	Kind SourceKind
	Path string
}
