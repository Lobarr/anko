package packages

import "github.com/cpmech/gosl/io"

func init() {
	Packages["gosl/io"] = map[string]interface{}{
		"Atob":                   io.Atob,
		"Atof":                   io.Atof,
		"Atoi":                   io.Atoi,
		"Btoa":                   io.Btoa,
		"Btoi":                   io.Btoi,
		"DblSf":                  io.DblSf,
		"ExtractStrPair":         io.ExtractStrPair,
		"Itob":                   io.Itob,
		"JoinKeys":               io.JoinKeys,
		"JoinKeys3":              io.JoinKeys3,
		"JoinKeys4":              io.JoinKeys4,
		"JoinKeysPre":            io.JoinKeysPre,
		"Keycode":                io.Keycode,
		"Keycodes":               io.Keycodes,
		"Pf":                     io.Pf,
		"Sf":                     io.Sf,
		"SplitFloats":            io.SplitFloats,
		"SplitInts":              io.SplitInts,
		"SplitKeys":              io.SplitKeys,
		"SplitKeys3":             io.SplitKeys3,
		"SplitKeys4":             io.SplitKeys4,
		"SplitSpacesQuoted ":     io.SplitSpacesQuoted,
		"SplitWithinParentheses": io.SplitWithinParentheses,
		"StrSf":                  io.StrSf,
		"StrSpaces":              io.StrSpaces,
		"StrThickLine":           io.StrThickLine,
		"StrThinLine":            io.StrThinLine,
		"TexNum":                 io.TexNum,
		"ToFcnConvertNum":        ToFcnConvertNum,
		"ToReadLinesCallback":    ToReadLinesCallback,
	}
	PackageTypes["gosl/io"] = map[string]interface{}{
		"Report": io.Report{},
	}
}

func ToFcnConvertNum(f func(int, float64) string) io.FcnConvertNum {
	return io.FcnConvertNum(f)
}

func ToReadLinesCallback(f func(int, string) bool) io.ReadLinesCallback {
	return io.ReadLinesCallback(f)
}
