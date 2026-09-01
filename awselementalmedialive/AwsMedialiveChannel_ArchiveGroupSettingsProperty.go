package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_ArchiveGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsMedialiveChannel#destination}
	// Experimental.
	Destination *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputGroupSettingsArchiveGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// archive_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_cdn_settings AwsMedialiveChannel#archive_cdn_settings}
	// Experimental.
	ArchiveCdnSettings *AwsMedialiveChannel_ArchiveCdnSettingsProperty `field:"optional" json:"archiveCdnSettings" yaml:"archiveCdnSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rollover_interval AwsMedialiveChannel#rollover_interval}.
	// Experimental.
	RolloverInterval *float64 `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

