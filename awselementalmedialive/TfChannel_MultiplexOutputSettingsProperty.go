package awselementalmedialive


// Experimental.
type TfChannel_MultiplexOutputSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination TfChannel#destination}
	// Experimental.
	Destination *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsMultiplexOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
}

