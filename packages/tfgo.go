package packages

import (
	"github.com/galeone/tfgo"
)

func init() {
	Packages["tfgo"] = map[string]interface{}{
		"Batchify":  tfgo.Batchify,
		"Cast":      tfgo.Cast,
		"Const":     tfgo.Const,
		"Exec":      tfgo.Exec,
		"IsClose":   tfgo.IsClose,
		"IsFloat":   tfgo.IsFloat,
		"IsInteger": tfgo.IsInteger,
		"MaxValue":  tfgo.MaxValue,
		"MinValue":  tfgo.MinValue,
		"NewRoot":   tfgo.NewRoot,
		"NewScope":  tfgo.NewScope,
		"NewTensor": tfgo.NewTensor,
	}
	PackageTypes["tfgo"] = map[string]interface{}{
		"Model":  tfgo.Model{},
		"Tensor": tfgo.Tensor{},
	}
}
