package elementalmedialive


// Experimental.
type AwsChannel_OutputGroupSettingsProperty struct {
	// archive_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_group_settings AwsChannel#archive_group_settings}
	// Experimental.
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// frame_capture_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_group_settings AwsChannel#frame_capture_group_settings}
	// Experimental.
	FrameCaptureGroupSettings *AwsChannel_FrameCaptureGroupSettingsProperty `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// hls_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_group_settings AwsChannel#hls_group_settings}
	// Experimental.
	HlsGroupSettings *AwsChannel_HlsGroupSettingsProperty `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// media_package_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_group_settings AwsChannel#media_package_group_settings}
	// Experimental.
	MediaPackageGroupSettings *AwsChannel_MediaPackageGroupSettingsProperty `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// ms_smooth_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_group_settings AwsChannel#ms_smooth_group_settings}
	// Experimental.
	MsSmoothGroupSettings *AwsChannel_MsSmoothGroupSettingsProperty `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// multiplex_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_group_settings AwsChannel#multiplex_group_settings}
	// Experimental.
	MultiplexGroupSettings *AwsChannel_MultiplexGroupSettingsProperty `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// rtmp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_group_settings AwsChannel#rtmp_group_settings}
	// Experimental.
	RtmpGroupSettings *AwsChannel_RtmpGroupSettingsProperty `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// udp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_group_settings AwsChannel#udp_group_settings}
	// Experimental.
	UdpGroupSettings *AwsChannel_UdpGroupSettingsProperty `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

