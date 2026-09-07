package s3


// Experimental.
type AwsBucketAnalyticsConfiguration_DataExportProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#destination AwsBucketAnalyticsConfiguration#destination}
	// Experimental.
	Destination *AwsBucketAnalyticsConfiguration_DestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#output_schema_version AwsBucketAnalyticsConfiguration#output_schema_version}.
	// Experimental.
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

