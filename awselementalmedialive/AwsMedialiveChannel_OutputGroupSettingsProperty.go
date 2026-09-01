package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_OutputGroupSettingsProperty struct {
	// archive_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_group_settings AwsMedialiveChannel#archive_group_settings}
	// Experimental.
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// frame_capture_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_group_settings AwsMedialiveChannel#frame_capture_group_settings}
	// Experimental.
	FrameCaptureGroupSettings *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// hls_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_group_settings AwsMedialiveChannel#hls_group_settings}
	// Experimental.
	HlsGroupSettings *AwsMedialiveChannel_HlsGroupSettingsProperty `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// media_package_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_group_settings AwsMedialiveChannel#media_package_group_settings}
	// Experimental.
	MediaPackageGroupSettings *AwsMedialiveChannel_MediaPackageGroupSettingsProperty `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// ms_smooth_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_group_settings AwsMedialiveChannel#ms_smooth_group_settings}
	// Experimental.
	MsSmoothGroupSettings *AwsMedialiveChannel_MsSmoothGroupSettingsProperty `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// multiplex_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_group_settings AwsMedialiveChannel#multiplex_group_settings}
	// Experimental.
	MultiplexGroupSettings *AwsMedialiveChannel_MultiplexGroupSettingsProperty `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// rtmp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_group_settings AwsMedialiveChannel#rtmp_group_settings}
	// Experimental.
	RtmpGroupSettings *AwsMedialiveChannel_RtmpGroupSettingsProperty `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// udp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_group_settings AwsMedialiveChannel#udp_group_settings}
	// Experimental.
	UdpGroupSettings *AwsMedialiveChannel_UdpGroupSettingsProperty `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

