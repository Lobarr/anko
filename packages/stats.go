package packages

import (
	"github.com/montanaflynn/stats"
)

func init() {
	Packages["stats"] = map[string]interface{}{
		"ChebyshevDistance":                 stats.ChebyshevDistance,
		"Correlation":                       stats.Correlation,
		"Covariance":                        stats.Covariance,
		"CovariancePopulation":              stats.CovariancePopulation,
		"EuclideanDistance":                 stats.EuclideanDistance,
		"GeometricMean":                     stats.GeometricMean,
		"HarmonicMean":                      stats.HarmonicMean,
		"InterQuartileRange":                stats.InterQuartileRange,
		"ManhattanDistance":                 stats.ManhattanDistance,
		"Max":                               stats.Max,
		"Mean":                              stats.Mean,
		"Median":                            stats.Median,
		"MedianAbsoluteDeviation":           stats.MedianAbsoluteDeviation,
		"MedianAbsoluteDeviationPopulation": stats.MedianAbsoluteDeviationPopulation,
		"Midhinge":                          stats.Midhinge,
		"Min":                               stats.Min,
		"MinkowskiDistance":                 stats.MinkowskiDistance,
		"Mode":                              stats.Mode,
		"Pearson":                           stats.Pearson,
		"Percentile":                        stats.Percentile,
		"PercentileNearestRank":             stats.PercentileNearestRank,
		"PopulationVariance":                stats.PopulationVariance,
		"Round":                             stats.Round,
		"Sample":                            stats.Sample,
		"SampleVariance":                    stats.SampleVariance,
		"StandardDeviation":                 stats.StandardDeviation,
		"StandardDeviationPopulation":       stats.StandardDeviationPopulation,
		"StandardDeviationSample":           stats.StandardDeviationSample,
		"StdDevP":                           stats.StdDevP,
		"StdDevS":                           stats.StdDevS,
		"Sum":                               stats.Sum,
		"Trimean":                           stats.Trimean,
		"VarP":                              stats.VarP,
		"VarS":                              stats.VarS,
		"Variance":                          stats.Variance,
		"Float64Data":                       Float64Data,
		"ToFloat64Data":                     ToFloat64Data,
		"LoadRawData":                       stats.LoadRawData,
		"QuartileOutliers":                  stats.QuartileOutliers,
		"Series":                            Series,
		"ToSeries":                          ToSeries,
	}

	PackageTypes["stats"] = map[string]interface{}{
		"Coordinate": stats.Coordinate{},
		"Outliers":   stats.Outliers{},
		"Quartiles":  stats.Quartiles{},
	}
}

func Float64Data() stats.Float64Data {
	var temp stats.Float64Data
	return temp
}

func ToFloat64Data(f []float64) stats.Float64Data {
	return stats.Float64Data(f)
}

func Series() stats.Series {
	var temp stats.Series
	return temp
}

func ToSeries(s []stats.Coordinate) stats.Series {
	return stats.Series(s)
}
