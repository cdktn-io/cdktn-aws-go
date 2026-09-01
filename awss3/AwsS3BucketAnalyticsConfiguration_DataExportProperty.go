package awss3


// Experimental.
type AwsS3BucketAnalyticsConfiguration_DataExportProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#destination AwsS3BucketAnalyticsConfiguration#destination}
	// Experimental.
	Destination *AwsS3BucketAnalyticsConfiguration_DestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#output_schema_version AwsS3BucketAnalyticsConfiguration#output_schema_version}.
	// Experimental.
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

