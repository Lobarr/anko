package packages

import (
	funk "github.com/thoas/go-funk"
)

func init() {
	Packages["funk"] = map[string]interface{}{
		"Contains":    funk.Contains,
		"IndexOf":     funk.IndexOf,
		"LastIndexOf": funk.LastIndexOf,
		"Get":         funk.Get,
		"Keys":        funk.Keys,
		"Values":      funk.Values,
		"Uniq":        funk.Uniq,
		"Sum":         funk.Sum,
		"Reverse":     funk.Reverse,
	}
}
