package awselementalmedialive


// Experimental.
type TfChannel_HlsCdnSettingsProperty struct {
	// hls_akamai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_akamai_settings TfChannel#hls_akamai_settings}
	// Experimental.
	HlsAkamaiSettings *TfChannel_HlsAkamaiSettingsProperty `field:"optional" json:"hlsAkamaiSettings" yaml:"hlsAkamaiSettings"`
	// hls_basic_put_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_basic_put_settings TfChannel#hls_basic_put_settings}
	// Experimental.
	HlsBasicPutSettings *TfChannel_HlsBasicPutSettingsProperty `field:"optional" json:"hlsBasicPutSettings" yaml:"hlsBasicPutSettings"`
	// hls_media_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_media_store_settings TfChannel#hls_media_store_settings}
	// Experimental.
	HlsMediaStoreSettings *TfChannel_HlsMediaStoreSettingsProperty `field:"optional" json:"hlsMediaStoreSettings" yaml:"hlsMediaStoreSettings"`
	// hls_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_s3_settings TfChannel#hls_s3_settings}
	// Experimental.
	HlsS3Settings *TfChannel_HlsS3SettingsProperty `field:"optional" json:"hlsS3Settings" yaml:"hlsS3Settings"`
	// hls_webdav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_webdav_settings TfChannel#hls_webdav_settings}
	// Experimental.
	HlsWebdavSettings *TfChannel_HlsWebdavSettingsProperty `field:"optional" json:"hlsWebdavSettings" yaml:"hlsWebdavSettings"`
}

