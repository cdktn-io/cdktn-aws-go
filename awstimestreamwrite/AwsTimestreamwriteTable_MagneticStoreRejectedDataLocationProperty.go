package awstimestreamwrite


// Experimental.
type AwsTimestreamwriteTable_MagneticStoreRejectedDataLocationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#s3_configuration AwsTimestreamwriteTable#s3_configuration}
	// Experimental.
	S3Configuration *AwsTimestreamwriteTable_S3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

