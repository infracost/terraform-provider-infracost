# infracost_aws_ecr_repository

Provides estimated usage data for an AWS ECR repository Requests.

## Example Usage

```hcl
resource "aws_ecr_repository" "repo" {
  name = "ecr-repo"
}

data "infracost_aws_ecr_repository" "costs" {
  resources = list(aws_ecr_repository.repo.id)

  storage_gb {
    value = 1
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the ECR Repository to apply the estimated usage.
* `storage_gb` - (Optional) The estimated total size of ECR repository in GB.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.

