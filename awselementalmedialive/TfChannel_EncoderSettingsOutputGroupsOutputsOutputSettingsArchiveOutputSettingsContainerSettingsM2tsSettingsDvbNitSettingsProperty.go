package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#network_id TfChannel#network_id}.
	// Experimental.
	NetworkId *float64 `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#network_name TfChannel#network_name}.
	// Experimental.
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rep_interval TfChannel#rep_interval}.
	// Experimental.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
}

