package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_column_delimiter TfApplication#record_column_delimiter}.
	// Experimental.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_row_delimiter TfApplication#record_row_delimiter}.
	// Experimental.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

