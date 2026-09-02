package awsfsx


// Experimental.
type TfS3AccessPointAttachment_S3AccessPointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#policy TfS3AccessPointAttachment#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#vpc_configuration TfS3AccessPointAttachment#vpc_configuration}
	// Experimental.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

