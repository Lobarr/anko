package packages

import (
	"fmt"
)

func init() {
	Packages["fmt"] = map[string]interface{}{
		"Errorf":   fmt.Errorf,
		"Fprint":   fmt.Fprint,
		"Fprintf":  fmt.Fprintf,
		"Fprintln": fmt.Fprintln,
		"Print":    fmt.Print,
		"Printf":   fmt.Printf,
		"Println":  fmt.Println,
		"Sprint":   fmt.Sprint,
		"Sprintf":  fmt.Sprintf,
		"Sprintln": fmt.Sprintln,
	}
}
