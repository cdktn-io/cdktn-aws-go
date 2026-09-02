package awstimestreamwrite


// Experimental.
type TfTable_MagneticStoreRejectedDataLocationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#s3_configuration TfTable#s3_configuration}
	// Experimental.
	S3Configuration *TfTable_S3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

