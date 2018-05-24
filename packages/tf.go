package packages

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

//TODO: add Input

func init() {
	Packages["tf"] = map[string]interface{}{
		"Version":        tf.Version,
		"NewGraph":       tf.NewGraph,
		"LoadSavedModel": tf.LoadSavedModel, //FIXME: make this work with gizo
		"NewSession":     tf.NewSession,
		"MakeShape":      tf.MakeShape,
		"ScalarShape":    tf.ScalarShape,
		"NewTensor":      tf.NewTensor,
		"ReadTensor":     tf.ReadTensor,
	}
	PackageTypes["tf"] = map[string]interface{}{
		"Float":      tf.Float,
		"Double":     tf.Double,
		"Int32":      tf.Int32,
		"UInt32":     tf.Uint32,
		"Uint8":      tf.Uint8,
		"Int16":      tf.Int16,
		"Int8":       tf.Int8,
		"String":     tf.String,
		"Complex64":  tf.Complex64,
		"Complex":    tf.Complex,
		"Int64":      tf.Int64,
		"Uint64":     tf.Uint64,
		"Bool":       tf.Bool,
		"Qint8":      tf.Qint8,
		"Quint8":     tf.Quint8,
		"Qint32":     tf.Qint32,
		"Bfloat16":   tf.Bfloat16,
		"Qint16":     tf.Qint16,
		"Quint16":    tf.Quint16,
		"Uint16":     tf.Uint16,
		"Complex128": tf.Complex128,
		"Half":       tf.Half,
		"Device":     tf.Device{},
		"Graph":      tf.Graph{},
		"OpSpec":     tf.OpSpec{},
		"Operation":  tf.Operation{},
		"Output":     tf.Output{},
		// "Input":      tf.Input{},
		"PartialRun":     tf.PartialRun{},
		"SavedModel":     tf.SavedModel{},
		"Session":        tf.Session{},
		"SessionOptions": tf.SessionOptions{},
		"Shape":          tf.Shape{},
		"Tensor":         tf.Tensor{},
	}
}

func OutputList(t []tf.Output) tf.OutputList {
	return tf.OutputList(t)
}
