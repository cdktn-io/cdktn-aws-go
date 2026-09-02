package awsbackup


// Experimental.
type TfReportPlan_ReportDeliveryChannelProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#s3_bucket_name TfReportPlan#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#formats TfReportPlan#formats}.
	// Experimental.
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#s3_key_prefix TfReportPlan#s3_key_prefix}.
	// Experimental.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

