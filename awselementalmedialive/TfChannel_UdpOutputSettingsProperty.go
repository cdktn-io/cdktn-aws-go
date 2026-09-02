package awselementalmedialive


// Experimental.
type TfChannel_UdpOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings TfChannel#container_settings}
	// Experimental.
	ContainerSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty `field:"required" json:"containerSettings" yaml:"containerSettings"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination TfChannel#destination}
	// Experimental.
	Destination *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buffer_msec TfChannel#buffer_msec}.
	// Experimental.
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// fec_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fec_output_settings TfChannel#fec_output_settings}
	// Experimental.
	FecOutputSettings *TfChannel_FecOutputSettingsProperty `field:"optional" json:"fecOutputSettings" yaml:"fecOutputSettings"`
}

