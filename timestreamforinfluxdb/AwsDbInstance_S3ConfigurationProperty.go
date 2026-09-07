package timestreamforinfluxdb


// Experimental.
type AwsDbInstance_S3ConfigurationProperty struct {
	// The name of the S3 bucket to deliver logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_instance#bucket_name AwsDbInstance#bucket_name}
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Indicates whether log delivery to the S3 bucket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_instance#enabled AwsDbInstance#enabled}
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

