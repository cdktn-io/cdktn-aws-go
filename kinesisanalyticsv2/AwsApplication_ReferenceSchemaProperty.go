package kinesisanalyticsv2


// Experimental.
type AwsApplication_ReferenceSchemaProperty struct {
	// record_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_column AwsApplication#record_column}
	// Experimental.
	RecordColumn interface{} `field:"required" json:"recordColumn" yaml:"recordColumn"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format AwsApplication#record_format}
	// Experimental.
	RecordFormat *AwsApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatProperty `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_encoding AwsApplication#record_encoding}.
	// Experimental.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

