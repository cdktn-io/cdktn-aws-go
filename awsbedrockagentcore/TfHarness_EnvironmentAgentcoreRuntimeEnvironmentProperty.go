package awsbedrockagentcore


// Experimental.
type TfHarness_EnvironmentAgentcoreRuntimeEnvironmentProperty struct {
	// filesystem_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#filesystem_configuration TfHarness#filesystem_configuration}
	// Experimental.
	FilesystemConfiguration interface{} `field:"optional" json:"filesystemConfiguration" yaml:"filesystemConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#lifecycle_configuration TfHarness#lifecycle_configuration}.
	// Experimental.
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#network_configuration TfHarness#network_configuration}
	// Experimental.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
}

