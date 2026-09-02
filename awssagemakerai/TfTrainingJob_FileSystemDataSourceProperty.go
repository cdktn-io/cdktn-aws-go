package awssagemakerai


// Experimental.
type TfTrainingJob_FileSystemDataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#directory_path TfTrainingJob#directory_path}.
	// Experimental.
	DirectoryPath *string `field:"required" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_access_mode TfTrainingJob#file_system_access_mode}.
	// Experimental.
	FileSystemAccessMode *string `field:"required" json:"fileSystemAccessMode" yaml:"fileSystemAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_id TfTrainingJob#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_type TfTrainingJob#file_system_type}.
	// Experimental.
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
}

