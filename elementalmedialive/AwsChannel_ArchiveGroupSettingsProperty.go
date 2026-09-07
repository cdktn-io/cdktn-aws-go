package elementalmedialive


// Experimental.
type AwsChannel_ArchiveGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsArchiveGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// archive_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_cdn_settings AwsChannel#archive_cdn_settings}
	// Experimental.
	ArchiveCdnSettings *AwsChannel_ArchiveCdnSettingsProperty `field:"optional" json:"archiveCdnSettings" yaml:"archiveCdnSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rollover_interval AwsChannel#rollover_interval}.
	// Experimental.
	RolloverInterval *float64 `field:"optional" json:"rolloverInterval" yaml:"rolloverInterval"`
}

