package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentProperty struct {
	// filesystem_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#filesystem_configuration AwsBedrockagentcoreHarness#filesystem_configuration}
	// Experimental.
	FilesystemConfiguration interface{} `field:"optional" json:"filesystemConfiguration" yaml:"filesystemConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#lifecycle_configuration AwsBedrockagentcoreHarness#lifecycle_configuration}.
	// Experimental.
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#network_configuration AwsBedrockagentcoreHarness#network_configuration}
	// Experimental.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
}

