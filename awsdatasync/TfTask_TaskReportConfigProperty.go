package awsdatasync


// Experimental.
type TfTask_TaskReportConfigProperty struct {
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#s3_destination TfTask#s3_destination}
	// Experimental.
	S3Destination *TfTask_S3DestinationProperty `field:"required" json:"s3Destination" yaml:"s3Destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#output_type TfTask#output_type}.
	// Experimental.
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#report_level TfTask#report_level}.
	// Experimental.
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
	// report_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#report_overrides TfTask#report_overrides}
	// Experimental.
	ReportOverrides *TfTask_ReportOverridesProperty `field:"optional" json:"reportOverrides" yaml:"reportOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#s3_object_versioning TfTask#s3_object_versioning}.
	// Experimental.
	S3ObjectVersioning *string `field:"optional" json:"s3ObjectVersioning" yaml:"s3ObjectVersioning"`
}

