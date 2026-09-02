package awskinesisanalytics


// Experimental.
type TfApplication_InputsSchemaRecordFormatMappingParametersProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#csv TfApplication#csv}
	// Experimental.
	Csv *TfApplication_InputsSchemaRecordFormatMappingParametersCsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#json TfApplication#json}
	// Experimental.
	Json *TfApplication_InputsSchemaRecordFormatMappingParametersJsonProperty `field:"optional" json:"json" yaml:"json"`
}

