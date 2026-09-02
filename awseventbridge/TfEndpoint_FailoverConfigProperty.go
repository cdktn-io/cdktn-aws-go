package awseventbridge


// Experimental.
type TfEndpoint_FailoverConfigProperty struct {
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#primary TfEndpoint#primary}
	// Experimental.
	Primary *TfEndpoint_PrimaryProperty `field:"required" json:"primary" yaml:"primary"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#secondary TfEndpoint#secondary}
	// Experimental.
	Secondary *TfEndpoint_SecondaryProperty `field:"required" json:"secondary" yaml:"secondary"`
}

