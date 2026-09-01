package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty struct {
	// efs_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#efs_access_point AwsBedrockagentcoreHarness#efs_access_point}
	// Experimental.
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// s3_files_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#s3_files_access_point AwsBedrockagentcoreHarness#s3_files_access_point}
	// Experimental.
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// session_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#session_storage AwsBedrockagentcoreHarness#session_storage}
	// Experimental.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

