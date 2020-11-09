package infracost

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsApiGatewayV2Api() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceRead,
        Schema: map[string]*schema.Schema{
            "resources":                resourcesSchema(),
            "average_message_size":     usageSchema(),
            "monthly_requests":         usageSchema(),
            "monthly_messages":         usageSchema(),
            "request_size":             usageSchema(),
        },
    }
}
