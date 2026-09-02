package awss3


// Experimental.
type TfBucket_ServerSideEncryptionConfigurationProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rule TfBucket#rule}
	// Experimental.
	Rule *TfBucket_ServerSideEncryptionConfigurationRuleProperty `field:"required" json:"rule" yaml:"rule"`
}

