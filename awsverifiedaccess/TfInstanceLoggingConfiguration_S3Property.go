package awsverifiedaccess


// Experimental.
type TfInstanceLoggingConfiguration_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#enabled TfInstanceLoggingConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#bucket_name TfInstanceLoggingConfiguration#bucket_name}.
	// Experimental.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#bucket_owner TfInstanceLoggingConfiguration#bucket_owner}.
	// Experimental.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_instance_logging_configuration#prefix TfInstanceLoggingConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

