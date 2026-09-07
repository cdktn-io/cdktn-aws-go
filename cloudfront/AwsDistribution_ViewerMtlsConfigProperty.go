package cloudfront


// Experimental.
type AwsDistribution_ViewerMtlsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#mode AwsDistribution#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// trust_store_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trust_store_config AwsDistribution#trust_store_config}
	// Experimental.
	TrustStoreConfig *AwsDistribution_TrustStoreConfigProperty `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

