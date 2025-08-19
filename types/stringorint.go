package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type StringOrInt int

func (d *StringOrInt) UnmarshalJSON(b []byte) error {
	str := strings.Replace(string(b), "\"", "", -1)
	if str == "" {
		*d = StringOrInt(0)
		return nil
	}

	if !isFloat.MatchString(str) {
		return fmt.Errorf("failed to match %s: %s", isFloat.String(), str)
	}

	parsed, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*d = StringOrInt(math.Trunc(parsed)) // truncate to make an int
	return nil
}
