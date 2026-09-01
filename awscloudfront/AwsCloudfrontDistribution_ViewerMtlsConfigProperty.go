package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_ViewerMtlsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#mode AwsCloudfrontDistribution#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// trust_store_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trust_store_config AwsCloudfrontDistribution#trust_store_config}
	// Experimental.
	TrustStoreConfig *AwsCloudfrontDistribution_TrustStoreConfigProperty `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

