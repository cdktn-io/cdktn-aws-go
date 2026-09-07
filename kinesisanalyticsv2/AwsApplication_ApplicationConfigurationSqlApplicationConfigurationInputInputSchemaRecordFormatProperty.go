package kinesisanalyticsv2


// Experimental.
type AwsApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatProperty struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#mapping_parameters AwsApplication#mapping_parameters}
	// Experimental.
	MappingParameters *AwsApplication_ApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersProperty `field:"required" json:"mappingParameters" yaml:"mappingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#record_format_type AwsApplication#record_format_type}.
	// Experimental.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

