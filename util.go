package viper

import (
	"log/slog"
)

type ConfigParseError struct {
	err error
}

func (pe ConfigParseError) Error() string { _ = "STUB: not implemented"; return "" }

func (pe ConfigParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func toCaseInsensitiveValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

func copyAndInsensitiviseMap(m map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func insensitiviseVal(val any) any { _ = "STUB: not implemented"; return *new(any) }

func insensitiviseMap(m map[string]any) { _ = "STUB: not implemented"; return }

func insensitiveArray(a []any) { _ = "STUB: not implemented"; return }

func absPathify(logger *slog.Logger, inPath string) string { _ = "STUB: not implemented"; return "" }

func userHomeDir() string { _ = "STUB: not implemented"; return "" }

func safeMul(a, b uint) uint { _ = "STUB: not implemented"; return 0 }

func parseSizeInBytes(sizeStr string) uint { _ = "STUB: not implemented"; return 0 }

func deepSearch(m map[string]any, path []string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
