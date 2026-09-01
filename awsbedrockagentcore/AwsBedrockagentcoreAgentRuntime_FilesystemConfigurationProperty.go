package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_FilesystemConfigurationProperty struct {
	// efs_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#efs_access_point AwsBedrockagentcoreAgentRuntime#efs_access_point}
	// Experimental.
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// s3_files_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#s3_files_access_point AwsBedrockagentcoreAgentRuntime#s3_files_access_point}
	// Experimental.
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// session_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#session_storage AwsBedrockagentcoreAgentRuntime#session_storage}
	// Experimental.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

