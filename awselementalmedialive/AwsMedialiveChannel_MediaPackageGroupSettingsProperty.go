package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_MediaPackageGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsMedialiveChannel#destination}
	// Experimental.
	Destination *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMediaPackageGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
}

