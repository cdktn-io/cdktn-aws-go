package elementalmedialive


// Experimental.
type AwsChannel_UdpOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings AwsChannel#container_settings}
	// Experimental.
	ContainerSettings *AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty `field:"required" json:"containerSettings" yaml:"containerSettings"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buffer_msec AwsChannel#buffer_msec}.
	// Experimental.
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// fec_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fec_output_settings AwsChannel#fec_output_settings}
	// Experimental.
	FecOutputSettings *AwsChannel_FecOutputSettingsProperty `field:"optional" json:"fecOutputSettings" yaml:"fecOutputSettings"`
}

