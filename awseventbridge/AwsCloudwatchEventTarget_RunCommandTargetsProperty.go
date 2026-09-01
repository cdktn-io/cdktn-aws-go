package awseventbridge


// Experimental.
type AwsCloudwatchEventTarget_RunCommandTargetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#key AwsCloudwatchEventTarget#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#values AwsCloudwatchEventTarget#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

