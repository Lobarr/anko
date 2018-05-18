package packages

import "testing"

func init() {
	PackageTypes["testing"] = map[string]interface{}{
		"T": testing.T{},
	}
}
