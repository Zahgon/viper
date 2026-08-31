package viper

import "github.com/spf13/pflag"

type FlagValueSet interface {
	VisitAll(fn func(FlagValue))
}

type FlagValue interface {
	HasChanged() bool
	Name() string
	ValueString() string
	ValueType() string
}

type pflagValueSet struct {
	flags *pflag.FlagSet
}

func (p pflagValueSet) VisitAll(fn func(flag FlagValue)) { _ = "STUB: not implemented"; return }

type pflagValue struct {
	flag *pflag.Flag
}

func (p pflagValue) HasChanged() bool { _ = "STUB: not implemented"; return false }

func (p pflagValue) Name() string { _ = "STUB: not implemented"; return "" }

func (p pflagValue) ValueString() string { _ = "STUB: not implemented"; return "" }

func (p pflagValue) ValueType() string { _ = "STUB: not implemented"; return "" }
