package formatter

import (
	"fmt"
)

func setFields(in interface{}) string {
	return fmt.Sprintf("%+v", in)
}
