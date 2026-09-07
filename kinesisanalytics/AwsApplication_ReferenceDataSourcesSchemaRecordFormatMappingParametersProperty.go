package kinesisanalytics


// Experimental.
type AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersProperty struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#csv AwsApplication#csv}
	// Experimental.
	Csv *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersCsvProperty `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#json AwsApplication#json}
	// Experimental.
	Json *AwsApplication_ReferenceDataSourcesSchemaRecordFormatMappingParametersJsonProperty `field:"optional" json:"json" yaml:"json"`
}

