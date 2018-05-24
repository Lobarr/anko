package packages

import "github.com/jinzhu/now"

func init() {
	Packages["now"] = map[string]interface{}{
		"BeginningOfDay":      now.BeginningOfDay,
		"BeginningOfHour":     now.BeginningOfHour,
		"BeginningOfMinute":   now.BeginningOfMinute,
		"BeginningOfMonth":    now.BeginningOfMonth,
		"BeginningOfQuarter":  now.BeginningOfQuarter,
		"BeginningOfWeek":     now.BeginningOfWeek,
		"BeginningOfYear":     now.BeginningOfYear,
		"Between":             now.Between,
		"EndOfDay":            now.EndOfDay,
		"EndOfHour":           now.EndOfHour,
		"EndOfMinute":         now.EndOfMinute,
		"EndOfMonth":          now.EndOfMonth,
		"EndOfQuarter":        now.EndOfQuarter,
		"EndOfSunday":         now.EndOfSunday,
		"EndOfWeek":           now.EndOfWeek,
		"EndOfYear":           now.EndOfYear,
		"Monday":              now.Monday,
		"MustParse":           now.MustParse,
		"MustParseInLocation": now.MustParseInLocation,
		"Parse":               now.Parse,
		"ParseInLocation":     now.ParseInLocation,
		"Sunday":              now.Sunday,
		"New":                 now.New,
	}
	PackageTypes["now"] = map[string]interface{}{
		"Now": now.Now{},
	}
}
