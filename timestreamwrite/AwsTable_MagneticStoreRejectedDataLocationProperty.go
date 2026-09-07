package timestreamwrite


// Experimental.
type AwsTable_MagneticStoreRejectedDataLocationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#s3_configuration AwsTable#s3_configuration}
	// Experimental.
	S3Configuration *AwsTable_S3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

