package kinesisanalytics


// Experimental.
type AwsApplication_InputsSchemaRecordFormatMappingParametersProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#csv AwsApplication#csv}
	// Experimental.
	Csv *AwsApplication_InputsSchemaRecordFormatMappingParametersCsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#json AwsApplication#json}
	// Experimental.
	Json *AwsApplication_InputsSchemaRecordFormatMappingParametersJsonProperty `field:"optional" json:"json" yaml:"json"`
}

