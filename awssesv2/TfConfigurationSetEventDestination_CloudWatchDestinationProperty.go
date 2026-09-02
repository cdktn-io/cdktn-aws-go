package awssesv2


// Experimental.
type TfConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// dimension_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#dimension_configuration TfConfigurationSetEventDestination#dimension_configuration}
	// Experimental.
	DimensionConfiguration interface{} `field:"required" json:"dimensionConfiguration" yaml:"dimensionConfiguration"`
}

