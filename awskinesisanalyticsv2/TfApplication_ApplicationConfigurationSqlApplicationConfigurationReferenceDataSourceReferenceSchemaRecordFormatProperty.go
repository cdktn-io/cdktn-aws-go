package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatProperty struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#mapping_parameters TfApplication#mapping_parameters}
	// Experimental.
	MappingParameters *TfApplication_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersProperty `field:"required" json:"mappingParameters" yaml:"mappingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format_type TfApplication#record_format_type}.
	// Experimental.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

