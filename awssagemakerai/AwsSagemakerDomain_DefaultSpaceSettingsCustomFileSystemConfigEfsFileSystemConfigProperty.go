package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultSpaceSettingsCustomFileSystemConfigEfsFileSystemConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#file_system_id AwsSagemakerDomain#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#file_system_path AwsSagemakerDomain#file_system_path}.
	// Experimental.
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
}

