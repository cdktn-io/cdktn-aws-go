package awsmskconnect


// Experimental.
type TfCustomPlugin_LocationProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_custom_plugin#s3 TfCustomPlugin#s3}
	// Experimental.
	S3 *TfCustomPlugin_S3Property `field:"required" json:"s3" yaml:"s3"`
}

