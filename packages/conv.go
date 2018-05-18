package packages

import (
	conv "github.com/cstockton/go-conv"
)

func init() {
	Packages["conv"] = map[string]interface{}{
		"Bool":     conv.Bool,
		"Duration": conv.Duration,
		"Float32":  conv.Float32,
		"Float64":  conv.Float64,
		"Int":      conv.Int,
		"Int8":     conv.Int8,
		"Int16":    conv.Int16,
		"Int32":    conv.Int32,
		"Int64":    conv.Int64,
		"String":   conv.String,
		"Time":     conv.Time,
		"Uint":     conv.Uint,
		"Uint8":    conv.Uint8,
		"Uint16":   conv.Uint16,
		"Uint32":   conv.Uint32,
		"Uint64":   conv.Uint64,
	}
}
