package awscodebuild


// Experimental.
type AwsCodebuildReportGroup_ExportConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#type AwsCodebuildReportGroup#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#s3_destination AwsCodebuildReportGroup#s3_destination}
	// Experimental.
	S3Destination *AwsCodebuildReportGroup_S3DestinationProperty `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

