package s3


// Experimental.
type AwsBucketVersioning_VersioningConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_versioning#status AwsBucketVersioning#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_versioning#mfa_delete AwsBucketVersioning#mfa_delete}.
	// Experimental.
	MfaDelete *string `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

