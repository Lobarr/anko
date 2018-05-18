package packages

import (
	"math/cmplx"
)

func init() {
	Packages["math/cmplx"] = map[string]interface{}{
		"Abs":   cmplx.Abs,
		"Acos":  cmplx.Acos,
		"Acosh": cmplx.Acosh,
		"Asin":  cmplx.Asin,
		"Asinh": cmplx.Asinh,
		"Atan":  cmplx.Atan,
		"Atanh": cmplx.Atanh,
		"Conj":  cmplx.Conj,
		"Cos":   cmplx.Cos,
		"Cosh":  cmplx.Cosh,
		"Cot":   cmplx.Cot,
		"Exp":   cmplx.Exp,
		"Inf":   cmplx.Inf,
		"IsInf": cmplx.IsInf,
		"IsNaN": cmplx.IsNaN,
		"Log":   cmplx.Log,
		"Log10": cmplx.Log10,
		"NaN":   cmplx.NaN,
		"Phase": cmplx.Phase,
		"Polar": cmplx.Polar,
		"Pow":   cmplx.Pow,
		"Rect":  cmplx.Rect,
		"Sin":   cmplx.Sin,
		"Sinh":  cmplx.Sinh,
		"Sqrt":  cmplx.Sqrt,
		"Tan":   cmplx.Tan,
		"Tanh":  cmplx.Tanh,
	}
}
