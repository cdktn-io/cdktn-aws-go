package ec2


// Experimental.
type AwsInstance_InstanceMarketOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#market_type AwsInstance#market_type}.
	// Experimental.
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// spot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#spot_options AwsInstance#spot_options}
	// Experimental.
	SpotOptions *AwsInstance_SpotOptionsProperty `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

