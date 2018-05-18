package packages

import "testing"

func init() {
	Packages["testing"] = map[string]interface{}{}
	PackageTypes["testing"] = map[string]interface{}{
		"T": testing.T{},
	}
}
