package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_column_delimiter AwsKinesisAnalyticsApplication#record_column_delimiter}.
	// Experimental.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_row_delimiter AwsKinesisAnalyticsApplication#record_row_delimiter}.
	// Experimental.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

