package kinesisanalytics


// Experimental.
type AwsApplication_ReferenceDataSourcesSchemaProperty struct {
	// record_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_columns AwsApplication#record_columns}
	// Experimental.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_format AwsApplication#record_format}
	// Experimental.
	RecordFormat *AwsApplication_ReferenceDataSourcesSchemaRecordFormatProperty `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#record_encoding AwsApplication#record_encoding}.
	// Experimental.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

