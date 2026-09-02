package awscodebuild


// Experimental.
type TfReportGroup_ExportConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#type TfReportGroup#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#s3_destination TfReportGroup#s3_destination}
	// Experimental.
	S3Destination *TfReportGroup_S3DestinationProperty `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

