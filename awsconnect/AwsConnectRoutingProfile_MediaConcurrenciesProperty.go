package awsconnect


// Experimental.
type AwsConnectRoutingProfile_MediaConcurrenciesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#channel AwsConnectRoutingProfile#channel}.
	// Experimental.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#concurrency AwsConnectRoutingProfile#concurrency}.
	// Experimental.
	Concurrency *float64 `field:"required" json:"concurrency" yaml:"concurrency"`
	// cross_channel_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#cross_channel_behavior AwsConnectRoutingProfile#cross_channel_behavior}
	// Experimental.
	CrossChannelBehavior *AwsConnectRoutingProfile_CrossChannelBehaviorProperty `field:"optional" json:"crossChannelBehavior" yaml:"crossChannelBehavior"`
}

