package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#csv AwsKinesisAnalyticsApplication#csv}
	// Experimental.
	Csv *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#json AwsKinesisAnalyticsApplication#json}
	// Experimental.
	Json *AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty `field:"optional" json:"json" yaml:"json"`
}

