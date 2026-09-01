package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_HlsCdnSettingsProperty struct {
	// hls_akamai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_akamai_settings AwsMedialiveChannel#hls_akamai_settings}
	// Experimental.
	HlsAkamaiSettings *AwsMedialiveChannel_HlsAkamaiSettingsProperty `field:"optional" json:"hlsAkamaiSettings" yaml:"hlsAkamaiSettings"`
	// hls_basic_put_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_basic_put_settings AwsMedialiveChannel#hls_basic_put_settings}
	// Experimental.
	HlsBasicPutSettings *AwsMedialiveChannel_HlsBasicPutSettingsProperty `field:"optional" json:"hlsBasicPutSettings" yaml:"hlsBasicPutSettings"`
	// hls_media_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_media_store_settings AwsMedialiveChannel#hls_media_store_settings}
	// Experimental.
	HlsMediaStoreSettings *AwsMedialiveChannel_HlsMediaStoreSettingsProperty `field:"optional" json:"hlsMediaStoreSettings" yaml:"hlsMediaStoreSettings"`
	// hls_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_s3_settings AwsMedialiveChannel#hls_s3_settings}
	// Experimental.
	HlsS3Settings *AwsMedialiveChannel_HlsS3SettingsProperty `field:"optional" json:"hlsS3Settings" yaml:"hlsS3Settings"`
	// hls_webdav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_webdav_settings AwsMedialiveChannel#hls_webdav_settings}
	// Experimental.
	HlsWebdavSettings *AwsMedialiveChannel_HlsWebdavSettingsProperty `field:"optional" json:"hlsWebdavSettings" yaml:"hlsWebdavSettings"`
}

