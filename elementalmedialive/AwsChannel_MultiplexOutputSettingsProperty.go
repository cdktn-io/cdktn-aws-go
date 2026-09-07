package elementalmedialive


// Experimental.
type AwsChannel_MultiplexOutputSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsMultiplexOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
}

