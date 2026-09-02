package awssesv2


// Experimental.
type TfConfigurationSetEventDestination_DimensionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#default_dimension_value TfConfigurationSetEventDestination#default_dimension_value}.
	// Experimental.
	DefaultDimensionValue *string `field:"required" json:"defaultDimensionValue" yaml:"defaultDimensionValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#dimension_name TfConfigurationSetEventDestination#dimension_name}.
	// Experimental.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#dimension_value_source TfConfigurationSetEventDestination#dimension_value_source}.
	// Experimental.
	DimensionValueSource *string `field:"required" json:"dimensionValueSource" yaml:"dimensionValueSource"`
}

