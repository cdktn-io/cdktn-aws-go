package awsec2


// Experimental.
type TfInstance_InstanceMarketOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#market_type TfInstance#market_type}.
	// Experimental.
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// spot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#spot_options TfInstance#spot_options}
	// Experimental.
	SpotOptions *TfInstance_SpotOptionsProperty `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

