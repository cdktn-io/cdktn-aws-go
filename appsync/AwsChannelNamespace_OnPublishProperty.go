package appsync


// Experimental.
type AwsChannelNamespace_OnPublishProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#behavior AwsChannelNamespace#behavior}.
	// Experimental.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// integration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#integration AwsChannelNamespace#integration}
	// Experimental.
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
}

