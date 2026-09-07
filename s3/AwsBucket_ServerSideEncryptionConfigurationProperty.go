package s3


// Experimental.
type AwsBucket_ServerSideEncryptionConfigurationProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rule AwsBucket#rule}
	// Experimental.
	Rule *AwsBucket_ServerSideEncryptionConfigurationRuleProperty `field:"required" json:"rule" yaml:"rule"`
}

