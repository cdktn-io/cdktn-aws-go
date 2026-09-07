package elementalmedialive


// Experimental.
type AwsChannel_OutputSettingsProperty struct {
	// archive_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_output_settings AwsChannel#archive_output_settings}
	// Experimental.
	ArchiveOutputSettings *AwsChannel_ArchiveOutputSettingsProperty `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// frame_capture_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_output_settings AwsChannel#frame_capture_output_settings}
	// Experimental.
	FrameCaptureOutputSettings *AwsChannel_FrameCaptureOutputSettingsProperty `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// hls_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_output_settings AwsChannel#hls_output_settings}
	// Experimental.
	HlsOutputSettings *AwsChannel_HlsOutputSettingsProperty `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// media_package_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_output_settings AwsChannel#media_package_output_settings}
	// Experimental.
	MediaPackageOutputSettings *AwsChannel_MediaPackageOutputSettingsProperty `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// ms_smooth_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_output_settings AwsChannel#ms_smooth_output_settings}
	// Experimental.
	MsSmoothOutputSettings *AwsChannel_MsSmoothOutputSettingsProperty `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// multiplex_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_output_settings AwsChannel#multiplex_output_settings}
	// Experimental.
	MultiplexOutputSettings *AwsChannel_MultiplexOutputSettingsProperty `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// rtmp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_output_settings AwsChannel#rtmp_output_settings}
	// Experimental.
	RtmpOutputSettings *AwsChannel_RtmpOutputSettingsProperty `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// udp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_output_settings AwsChannel#udp_output_settings}
	// Experimental.
	UdpOutputSettings *AwsChannel_UdpOutputSettingsProperty `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}

