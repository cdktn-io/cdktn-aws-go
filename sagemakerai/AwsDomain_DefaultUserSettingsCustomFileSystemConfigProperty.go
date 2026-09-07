package sagemakerai


// Experimental.
type AwsDomain_DefaultUserSettingsCustomFileSystemConfigProperty struct {
	// efs_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#efs_file_system_config AwsDomain#efs_file_system_config}
	// Experimental.
	EfsFileSystemConfig *AwsDomain_DefaultUserSettingsCustomFileSystemConfigEfsFileSystemConfigProperty `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
}

