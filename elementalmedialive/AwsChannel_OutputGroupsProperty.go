package elementalmedialive


// Experimental.
type AwsChannel_OutputGroupsProperty struct {
	// output_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_group_settings AwsChannel#output_group_settings}
	// Experimental.
	OutputGroupSettings *AwsChannel_OutputGroupSettingsProperty `field:"required" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#outputs AwsChannel#outputs}
	// Experimental.
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsChannel#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

