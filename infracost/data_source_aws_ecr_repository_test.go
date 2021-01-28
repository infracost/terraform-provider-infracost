package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsECRRepository(t *testing.T) {
	name := "data.infracost_aws_ecr_repository.ecr_repository"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsECRRepositoryConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "storage_gb", 100000),
				),
			},
		},
	})
}

func testAwsECRRepositoryConfig() string {
	return `
		data "infracost_aws_ecr_repository" "ecr_repository" {
			resources = list("ecr_repo_1", "ecr_repo_2")

			storage_gb {
				value = 100000
			}
		}
	`
}
