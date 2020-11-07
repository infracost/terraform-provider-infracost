# infracost_aws_codebuild_project

Provides estimated usage data for an AWS Codebuild Project.

## Example Usage

```hcl
resource "aws_codebuild_project" "build_project" {
  name           = "test-project-cache"
  description    = "test_codebuild_project_cache"
  build_timeout  = "5"
  queued_timeout = "5"

  service_role = "arn:aws:iam:12345678::role/project"

  artifacts {
    type = "NO_ARTIFACTS"
  }

  cache {
    type  = "LOCAL"
    modes = ["LOCAL_DOCKER_LAYER_CACHE", "LOCAL_SOURCE_CACHE"]
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:1.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

  }

  source {
    type            = "GITHUB"
    location        = "https://github.com/infracost/infracost.git"
    git_clone_depth = 1
  }

  tags = {
    Environment = "Test"
  }
}

data "infracost_aws_codebuild_project" "build_cost" {
  resources = list(aws_codebuild_project.build_project.id)

  monthly_build_minutes {
    value = 10000
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the Codebuild Projects to apply the estimated usage.
* `monthly_build_minutes` - (Required) The estimated total duration of monthly codebuild execution usage in minutes.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.

