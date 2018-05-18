package packages

import (
	"github.com/cpmech/gosl/fun"
	"github.com/cpmech/gosl/la"
)

func init() {
	Packages["gosl/fun"] = map[string]interface{}{
		"Atan2p":               fun.Atan2p,
		"Atan2pDeg":            fun.Atan2pDeg,
		"Beta":                 fun.Beta,
		"Binomial":             fun.Binomial,
		"Boxcar":               fun.Boxcar,
		"CarlsonRc":            fun.CarlsonRc,
		"CarlsonRd":            fun.CarlsonRd,
		"CarlsonRf":            fun.CarlsonRf,
		"CarlsonRj":            fun.CarlsonRj,
		"ChebyshevT":           fun.ChebyshevT,
		"ChebyshevTdiff1":      fun.ChebyshevTdiff1,
		"ChebyshevTdiff2":      fun.ChebyshevTdiff2,
		"ChebyshevXgauss":      fun.ChebyshevXgauss,
		"ChebyshevXlob":        fun.ChebyshevXlob,
		"Dft1d":                fun.Dft1d,
		"Elliptic1":            fun.Elliptic1,
		"Elliptic2":            fun.Elliptic2,
		"Elliptic3":            fun.Elliptic3,
		"ExpMix":               fun.ExpMix,
		"ExpPix":               fun.ExpPix,
		"Factorial100":         fun.Factorial100,
		"Factorial22":          fun.Factorial22,
		"Hat":                  fun.Hat,
		"HatD1":                fun.HatD1,
		"Heav":                 fun.Heav,
		"ImagPowN":             fun.ImagPowN,
		"ImagXpowN":            fun.ImagXpowN,
		"Logistic":             fun.Logistic,
		"LogisticD1":           fun.LogisticD1,
		"ModBesselI0":          fun.ModBesselI0,
		"ModBesselI1":          fun.ModBesselI1,
		"ModBesselIn":          fun.ModBesselIn,
		"ModBesselK0":          fun.ModBesselK0,
		"ModBesselK1":          fun.ModBesselK1,
		"ModBesselKn":          fun.ModBesselKn,
		"NegOnePowN":           fun.NegOnePowN,
		"PlotLagInterpI":       fun.PlotLagInterpI,
		"PlotLagInterpL":       fun.PlotLagInterpL,
		"PlotLagInterpW":       fun.PlotLagInterpW,
		"Pow2":                 fun.Pow2,
		"Pow3":                 fun.Pow3,
		"PowP":                 fun.PowP,
		"Ramp":                 fun.Ramp,
		"Rbinomial":            fun.Rbinomial,
		"Rect":                 fun.Rect,
		"Sabs":                 fun.Sabs,
		"SabsD1":               fun.SabsD1,
		"SabsD2":               fun.SabsD2,
		"Sign":                 fun.Sign,
		"Sinc":                 fun.Sinc,
		"Sramp":                fun.Sramp,
		"SrampD1":              fun.SrampD1,
		"SrampD2":              fun.SrampD2,
		"SuqCos":               fun.SuqCos,
		"SuqSin":               fun.SuqSin,
		"UintBinomial":         fun.UintBinomial,
		"NewAxis":              fun.NewAxis,
		"NewBiLinear":          fun.NewBiLinear,
		"NewChebyInterp":       fun.NewChebyInterp,
		"NewDataInterp":        fun.NewDataInterp,
		"NewFourierInterp":     fun.NewFourierInterp,
		"NewGeneralOrthoPoly":  fun.NewGeneralOrthoPoly,
		"NewInterpCubic":       fun.NewInterpCubic,
		"NewInterpQuad":        fun.NewInterpQuad,
		"NewLagIntSet":         fun.NewLagIntSet,
		"NewLagrangeInterp":    fun.NewLagrangeInterp,
		"NewSinusoidBasis":     fun.NewSinusoidBasis,
		"NewSinusoidEssential": fun.NewSinusoidEssential,
		"ToInterpType":         ToInterpType,
		"ToSs":                 ToSs,
		"ToSss":                ToSss,
		"ToSv":                 ToSv,
		"ToSvs":                ToSvs,
		"ToTt":                 ToTt,
		"ToTv":                 ToTv,
		"ToVs":                 ToVs,
		"ToVss":                ToVss,
		"ToVv":                 ToVv,
		"ToVvss":               ToVvss,
		"ToVvvss":              ToVvvss,
	}
	PackageTypes["gosl/fun"] = map[string]interface{}{
		"Axis":             fun.Axis{},
		"BiLinear":         fun.BiLinear{},
		"ChebyInterp":      fun.ChebyInterp{},
		"DataInterp":       fun.DataInterp{},
		"FourierInterp":    fun.FourierInterp{},
		"GeneralOrthoPoly": fun.GeneralOrthoPoly{},
		"InterpCubic":      fun.InterpCubic{},
		"InterpQuad":       fun.InterpQuad{},
		"LagIntSet":        fun.LagIntSet{},
		"LagrangeInterp":   fun.LagrangeInterp{},
		"Sinusoid":         fun.Sinusoid{},
	}
}

func ToInterpType(i int) fun.InterpType {
	return fun.InterpType(i)
}

func ToSs(f func(float64) float64) fun.Ss {
	return fun.Ss(f)
}

func ToSss(f func(float64, float64) float64) fun.Sss {
	return fun.Sss(f)
}

func ToSv(f func(la.Vector) float64) fun.Sv {
	return fun.Sv(f)
}

func ToSvs(f func(la.Vector, float64) float64) fun.Svs {
	return fun.Svs(f)
}

func ToTt(f func(*la.Triplet, *la.Triplet)) fun.Tt {
	return fun.Tt(f)
}

func ToTv(f func(*la.Triplet, la.Vector)) fun.Tv {
	return fun.Tv(f)
}

func ToVs(f func(la.Vector, float64)) fun.Vs {
	return fun.Vs(f)
}

func ToVss(f func(la.Vector, float64, float64)) fun.Vss {
	return fun.Vss(f)
}

func ToVv(f func(la.Vector, la.Vector)) fun.Vv {
	return fun.Vv(f)
}

func ToVvss(f func(la.Vector, la.Vector, float64, float64)) fun.Vvss {
	return fun.Vvss(f)
}

func ToVvvss(f func(la.Vector, la.Vector, la.Vector, float64, float64)) fun.Vvvss {
	return fun.Vvvss(f)
}
