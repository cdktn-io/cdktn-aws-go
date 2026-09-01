package awselementalmedialive


// Experimental.
type AwsMedialiveMultiplexProgram_MultiplexProgramSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#preferred_channel_pipeline AwsMedialiveMultiplexProgram#preferred_channel_pipeline}.
	// Experimental.
	PreferredChannelPipeline *string `field:"required" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#program_number AwsMedialiveMultiplexProgram#program_number}.
	// Experimental.
	ProgramNumber *float64 `field:"required" json:"programNumber" yaml:"programNumber"`
	// service_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#service_descriptor AwsMedialiveMultiplexProgram#service_descriptor}
	// Experimental.
	ServiceDescriptor interface{} `field:"optional" json:"serviceDescriptor" yaml:"serviceDescriptor"`
	// video_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#video_settings AwsMedialiveMultiplexProgram#video_settings}
	// Experimental.
	VideoSettings interface{} `field:"optional" json:"videoSettings" yaml:"videoSettings"`
}

