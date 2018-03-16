package validate

import (
	"fmt"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/hashicorp/terraform/helper/schema"
	"regexp"
)

func ListIntBetween(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		values, ok := i.([]int)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be []int", k))
			return
		}

		for v := range values {
			if v < min || v > max {
				es = append(es, fmt.Errorf("expected %s values to be in the range (%d - %d), got %d", k, min, max, v))
			}
		}

		return
	}
}

func Rfc3339Time(v interface{}, k string) (ws []string, errors []error) {
	if _, err := date.ParseTime(time.RFC3339, v.(string)); err != nil {
		errors = append(errors, fmt.Errorf("%q has the invalid RFC3339 date format %q: %+v", k, v, err))
	}

	return
}

func Iso8601Duration(v interface{}, k string) (ws []string, errors []error) {
	expression := `^(R\d*\/)?P(?:\d+(?:\.\d+)?Y)?(?:\d+(?:\.\d+)?M)?(?:\d+(?:\.\d+)?W)?(?:\d+(?:\.\d+)?D)?(?:T(?:\d+(?:\.\d+)?H)?(?:\d+(?:\.\d+)?M)?(?:\d+(?:\.\d+)?S)?)?$`
	if ok := regexp.MustCompile(expression).MatchString(v.(string)); !ok {
		errors = append(errors, fmt.Errorf("%q has the invalid ISO8601 duration format %q", k, v))
	}

	return
}
