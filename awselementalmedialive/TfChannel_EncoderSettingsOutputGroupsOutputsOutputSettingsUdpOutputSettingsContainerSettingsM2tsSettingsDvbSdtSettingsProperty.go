package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_sdt TfChannel#output_sdt}.
	// Experimental.
	OutputSdt *string `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rep_interval TfChannel#rep_interval}.
	// Experimental.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#service_name TfChannel#service_name}.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#service_provider_name TfChannel#service_provider_name}.
	// Experimental.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

