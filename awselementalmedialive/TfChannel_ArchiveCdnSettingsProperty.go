package awselementalmedialive


// Experimental.
type TfChannel_ArchiveCdnSettingsProperty struct {
	// archive_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_s3_settings TfChannel#archive_s3_settings}
	// Experimental.
	ArchiveS3Settings *TfChannel_ArchiveS3SettingsProperty `field:"optional" json:"archiveS3Settings" yaml:"archiveS3Settings"`
}

