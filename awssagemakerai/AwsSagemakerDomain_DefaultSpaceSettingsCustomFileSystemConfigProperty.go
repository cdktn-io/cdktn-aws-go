package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultSpaceSettingsCustomFileSystemConfigProperty struct {
	// efs_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#efs_file_system_config AwsSagemakerDomain#efs_file_system_config}
	// Experimental.
	EfsFileSystemConfig *AwsSagemakerDomain_DefaultSpaceSettingsCustomFileSystemConfigEfsFileSystemConfigProperty `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
}

