package infracost

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"infracost": testAccProvider,
	}
}

func testCheckResourceAttrValue(name string, key string, expectedVal float64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Resource not found: %s", r)
		}
		return testCheckResourceAttrTypeSetFloatVal(r.Primary, key, "value", expectedVal)
	}
}

func testCheckResourceAttrTypeSetFloatVal(is *terraform.InstanceState, key string, typeSetAttr string, expected float64) error {
	var found, val = false, 0.0
	var err error

	for k, v := range is.Attributes {
		if strings.HasPrefix(k, fmt.Sprintf("%s.", key)) {
			if strings.HasSuffix(k, fmt.Sprintf(".%s", typeSetAttr)) {
				found = true
				val, err = strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("Could not parse %s value to float: %s", typeSetAttr, v)
				}
			}
		}
	}

	if !found {
		return fmt.Errorf("No %s value found", typeSetAttr)
	}
	if val != expected {
		return fmt.Errorf("Unexpected %s value: %v, expected: %v", typeSetAttr, val, expected)
	}
	return nil
}
