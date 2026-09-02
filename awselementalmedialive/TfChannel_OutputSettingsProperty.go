package awselementalmedialive


// Experimental.
type TfChannel_OutputSettingsProperty struct {
	// archive_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#archive_output_settings TfChannel#archive_output_settings}
	// Experimental.
	ArchiveOutputSettings *TfChannel_ArchiveOutputSettingsProperty `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// frame_capture_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_output_settings TfChannel#frame_capture_output_settings}
	// Experimental.
	FrameCaptureOutputSettings *TfChannel_FrameCaptureOutputSettingsProperty `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// hls_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_output_settings TfChannel#hls_output_settings}
	// Experimental.
	HlsOutputSettings *TfChannel_HlsOutputSettingsProperty `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// media_package_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#media_package_output_settings TfChannel#media_package_output_settings}
	// Experimental.
	MediaPackageOutputSettings *TfChannel_MediaPackageOutputSettingsProperty `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// ms_smooth_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ms_smooth_output_settings TfChannel#ms_smooth_output_settings}
	// Experimental.
	MsSmoothOutputSettings *TfChannel_MsSmoothOutputSettingsProperty `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// multiplex_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_output_settings TfChannel#multiplex_output_settings}
	// Experimental.
	MultiplexOutputSettings *TfChannel_MultiplexOutputSettingsProperty `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// rtmp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_output_settings TfChannel#rtmp_output_settings}
	// Experimental.
	RtmpOutputSettings *TfChannel_RtmpOutputSettingsProperty `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// udp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#udp_output_settings TfChannel#udp_output_settings}
	// Experimental.
	UdpOutputSettings *TfChannel_UdpOutputSettingsProperty `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}

