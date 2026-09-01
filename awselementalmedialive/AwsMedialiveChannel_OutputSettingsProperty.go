package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_OutputSettingsProperty struct {
	// archive_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_output_settings AwsMedialiveChannel#archive_output_settings}
	// Experimental.
	ArchiveOutputSettings *AwsMedialiveChannel_ArchiveOutputSettingsProperty `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// frame_capture_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_output_settings AwsMedialiveChannel#frame_capture_output_settings}
	// Experimental.
	FrameCaptureOutputSettings *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// hls_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_output_settings AwsMedialiveChannel#hls_output_settings}
	// Experimental.
	HlsOutputSettings *AwsMedialiveChannel_HlsOutputSettingsProperty `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// media_package_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_output_settings AwsMedialiveChannel#media_package_output_settings}
	// Experimental.
	MediaPackageOutputSettings *AwsMedialiveChannel_MediaPackageOutputSettingsProperty `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// ms_smooth_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_output_settings AwsMedialiveChannel#ms_smooth_output_settings}
	// Experimental.
	MsSmoothOutputSettings *AwsMedialiveChannel_MsSmoothOutputSettingsProperty `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// multiplex_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_output_settings AwsMedialiveChannel#multiplex_output_settings}
	// Experimental.
	MultiplexOutputSettings *AwsMedialiveChannel_MultiplexOutputSettingsProperty `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// rtmp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_output_settings AwsMedialiveChannel#rtmp_output_settings}
	// Experimental.
	RtmpOutputSettings *AwsMedialiveChannel_RtmpOutputSettingsProperty `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// udp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_output_settings AwsMedialiveChannel#udp_output_settings}
	// Experimental.
	UdpOutputSettings *AwsMedialiveChannel_UdpOutputSettingsProperty `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}

