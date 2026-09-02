package awskinesisanalytics


// Experimental.
type TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#csv TfApplication#csv}
	// Experimental.
	Csv *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#json TfApplication#json}
	// Experimental.
	Json *TfApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty `field:"optional" json:"json" yaml:"json"`
}

