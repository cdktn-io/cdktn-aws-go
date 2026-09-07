package fsx


// Experimental.
type AwsS3AccessPointAttachment_S3AccessPointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#policy AwsS3AccessPointAttachment#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#vpc_configuration AwsS3AccessPointAttachment#vpc_configuration}
	// Experimental.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

