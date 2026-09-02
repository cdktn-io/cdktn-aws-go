package awsbedrockagentcore


// Experimental.
type TfAgentRuntime_EfsAccessPointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#access_point_arn TfAgentRuntime#access_point_arn}.
	// Experimental.
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#mount_path TfAgentRuntime#mount_path}.
	// Experimental.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
}

