package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_OutputGroupsProperty struct {
	// output_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_group_settings AwsMedialiveChannel#output_group_settings}
	// Experimental.
	OutputGroupSettings *AwsMedialiveChannel_OutputGroupSettingsProperty `field:"required" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#outputs AwsMedialiveChannel#outputs}
	// Experimental.
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsMedialiveChannel#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

