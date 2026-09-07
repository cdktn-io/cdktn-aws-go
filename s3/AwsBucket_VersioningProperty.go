package s3


// Experimental.
type AwsBucket_VersioningProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#enabled AwsBucket#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#mfa_delete AwsBucket#mfa_delete}.
	// Experimental.
	MfaDelete interface{} `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

