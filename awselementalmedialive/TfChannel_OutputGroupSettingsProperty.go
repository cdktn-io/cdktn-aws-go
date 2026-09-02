package awselementalmedialive


// Experimental.
type TfChannel_OutputGroupSettingsProperty struct {
	// archive_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_group_settings TfChannel#archive_group_settings}
	// Experimental.
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// frame_capture_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_group_settings TfChannel#frame_capture_group_settings}
	// Experimental.
	FrameCaptureGroupSettings *TfChannel_FrameCaptureGroupSettingsProperty `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// hls_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_group_settings TfChannel#hls_group_settings}
	// Experimental.
	HlsGroupSettings *TfChannel_HlsGroupSettingsProperty `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// media_package_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_group_settings TfChannel#media_package_group_settings}
	// Experimental.
	MediaPackageGroupSettings *TfChannel_MediaPackageGroupSettingsProperty `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// ms_smooth_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_group_settings TfChannel#ms_smooth_group_settings}
	// Experimental.
	MsSmoothGroupSettings *TfChannel_MsSmoothGroupSettingsProperty `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// multiplex_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_group_settings TfChannel#multiplex_group_settings}
	// Experimental.
	MultiplexGroupSettings *TfChannel_MultiplexGroupSettingsProperty `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// rtmp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_group_settings TfChannel#rtmp_group_settings}
	// Experimental.
	RtmpGroupSettings *TfChannel_RtmpGroupSettingsProperty `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// udp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_group_settings TfChannel#udp_group_settings}
	// Experimental.
	UdpGroupSettings *TfChannel_UdpGroupSettingsProperty `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

