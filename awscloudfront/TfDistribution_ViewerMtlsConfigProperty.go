package awscloudfront


// Experimental.
type TfDistribution_ViewerMtlsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#mode TfDistribution#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// trust_store_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trust_store_config TfDistribution#trust_store_config}
	// Experimental.
	TrustStoreConfig *TfDistribution_TrustStoreConfigProperty `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

