package packages

import (
	"io/ioutil"
)

func init() {
	Packages["io/ioutil"] = map[string]interface{}{
		"ReadFile":  ioutil.ReadFile,
		"WriteFile": ioutil.WriteFile,
	}
}
