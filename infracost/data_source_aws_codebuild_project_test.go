package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsCodebuildProject(t *testing.T) {
	name := "data.infracost_aws_codebuild_project.my_codebuild_project"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsCodebuildProjectConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_build_minutes", 10000),
				),
			},
		},
	})
}

func testAwsCodebuildProjectConfig() string {
	return `
		data "infracost_aws_codebuild_project" "my_codebuild_project" {
			resources = list("my_codebuild_project_1", "my_codebuild_project_2")

			monthly_build_minutes {
				value = 10000
			}
		}
	`
}
