package eventbridge


// Experimental.
type AwsEndpoint_FailoverConfigProperty struct {
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#primary AwsEndpoint#primary}
	// Experimental.
	Primary *AwsEndpoint_PrimaryProperty `field:"required" json:"primary" yaml:"primary"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#secondary AwsEndpoint#secondary}
	// Experimental.
	Secondary *AwsEndpoint_SecondaryProperty `field:"required" json:"secondary" yaml:"secondary"`
}

