package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_InputsSchemaProperty struct {
	// record_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_columns AwsKinesisAnalyticsApplication#record_columns}
	// Experimental.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_format AwsKinesisAnalyticsApplication#record_format}
	// Experimental.
	RecordFormat *AwsKinesisAnalyticsApplication_InputsSchemaRecordFormatProperty `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_encoding AwsKinesisAnalyticsApplication#record_encoding}.
	// Experimental.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

