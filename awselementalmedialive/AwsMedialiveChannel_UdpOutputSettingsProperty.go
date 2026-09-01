package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_UdpOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings AwsMedialiveChannel#container_settings}
	// Experimental.
	ContainerSettings *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty `field:"required" json:"containerSettings" yaml:"containerSettings"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsMedialiveChannel#destination}
	// Experimental.
	Destination *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buffer_msec AwsMedialiveChannel#buffer_msec}.
	// Experimental.
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// fec_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fec_output_settings AwsMedialiveChannel#fec_output_settings}
	// Experimental.
	FecOutputSettings *AwsMedialiveChannel_FecOutputSettingsProperty `field:"optional" json:"fecOutputSettings" yaml:"fecOutputSettings"`
}

