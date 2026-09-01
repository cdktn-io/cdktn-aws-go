package awseventbridge


// Experimental.
type AwsCloudwatchEventPermission_ConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_permission#key AwsCloudwatchEventPermission#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_permission#type AwsCloudwatchEventPermission#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_permission#value AwsCloudwatchEventPermission#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

