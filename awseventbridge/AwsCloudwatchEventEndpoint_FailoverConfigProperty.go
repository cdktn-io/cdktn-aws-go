package awseventbridge


// Experimental.
type AwsCloudwatchEventEndpoint_FailoverConfigProperty struct {
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#primary AwsCloudwatchEventEndpoint#primary}
	// Experimental.
	Primary *AwsCloudwatchEventEndpoint_PrimaryProperty `field:"required" json:"primary" yaml:"primary"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#secondary AwsCloudwatchEventEndpoint#secondary}
	// Experimental.
	Secondary *AwsCloudwatchEventEndpoint_SecondaryProperty `field:"required" json:"secondary" yaml:"secondary"`
}

