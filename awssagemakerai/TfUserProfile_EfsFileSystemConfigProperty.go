package awssagemakerai


// Experimental.
type TfUserProfile_EfsFileSystemConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#file_system_id TfUserProfile#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#file_system_path TfUserProfile#file_system_path}.
	// Experimental.
	FileSystemPath *string `field:"optional" json:"fileSystemPath" yaml:"fileSystemPath"`
}

