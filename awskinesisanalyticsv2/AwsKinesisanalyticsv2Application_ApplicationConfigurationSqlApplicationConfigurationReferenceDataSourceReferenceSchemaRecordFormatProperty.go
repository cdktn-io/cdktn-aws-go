package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatProperty struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#mapping_parameters AwsKinesisanalyticsv2Application#mapping_parameters}
	// Experimental.
	MappingParameters *AwsKinesisanalyticsv2Application_ApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersProperty `field:"required" json:"mappingParameters" yaml:"mappingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format_type AwsKinesisanalyticsv2Application#record_format_type}.
	// Experimental.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

