package awseventbridge


// Experimental.
type AwsCloudwatchEventBus_LogConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_bus#include_detail AwsCloudwatchEventBus#include_detail}.
	// Experimental.
	IncludeDetail *string `field:"optional" json:"includeDetail" yaml:"includeDetail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_bus#level AwsCloudwatchEventBus#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

