package awselementalmedialive


// Experimental.
type AwsMedialiveMultiplexProgram_VideoSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#constant_bitrate AwsMedialiveMultiplexProgram#constant_bitrate}.
	// Experimental.
	ConstantBitrate *float64 `field:"optional" json:"constantBitrate" yaml:"constantBitrate"`
	// statmux_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#statmux_settings AwsMedialiveMultiplexProgram#statmux_settings}
	// Experimental.
	StatmuxSettings interface{} `field:"optional" json:"statmuxSettings" yaml:"statmuxSettings"`
}

