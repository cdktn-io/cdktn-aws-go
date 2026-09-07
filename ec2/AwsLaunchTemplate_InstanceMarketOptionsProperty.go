package ec2


// Experimental.
type AwsLaunchTemplate_InstanceMarketOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#market_type AwsLaunchTemplate#market_type}.
	// Experimental.
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// spot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#spot_options AwsLaunchTemplate#spot_options}
	// Experimental.
	SpotOptions *AwsLaunchTemplate_SpotOptionsProperty `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

