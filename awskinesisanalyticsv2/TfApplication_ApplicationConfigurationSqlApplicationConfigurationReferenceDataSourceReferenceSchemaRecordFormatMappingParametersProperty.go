package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersProperty struct {
	// csv_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#csv_mapping_parameters TfApplication#csv_mapping_parameters}
	// Experimental.
	CsvMappingParameters *TfApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersCsvMappingParametersProperty `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// json_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#json_mapping_parameters TfApplication#json_mapping_parameters}
	// Experimental.
	JsonMappingParameters *TfApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersProperty `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

