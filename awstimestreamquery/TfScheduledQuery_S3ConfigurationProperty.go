package awstimestreamquery


// Experimental.
type TfScheduledQuery_S3ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#bucket_name TfScheduledQuery#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#encryption_option TfScheduledQuery#encryption_option}.
	// Experimental.
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#object_key_prefix TfScheduledQuery#object_key_prefix}.
	// Experimental.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

