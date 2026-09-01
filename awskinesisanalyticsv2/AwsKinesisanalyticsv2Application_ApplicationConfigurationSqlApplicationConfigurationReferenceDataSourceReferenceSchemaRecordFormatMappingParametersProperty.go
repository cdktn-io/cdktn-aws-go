package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersProperty struct {
	// csv_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#csv_mapping_parameters AwsKinesisanalyticsv2Application#csv_mapping_parameters}
	// Experimental.
	CsvMappingParameters *AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersCsvMappingParametersProperty `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// json_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#json_mapping_parameters AwsKinesisanalyticsv2Application#json_mapping_parameters}
	// Experimental.
	JsonMappingParameters *AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersProperty `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

