package sagemakerai


// Experimental.
type AwsTrainingJob_FileSystemDataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#directory_path AwsTrainingJob#directory_path}.
	// Experimental.
	DirectoryPath *string `field:"required" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_access_mode AwsTrainingJob#file_system_access_mode}.
	// Experimental.
	FileSystemAccessMode *string `field:"required" json:"fileSystemAccessMode" yaml:"fileSystemAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_id AwsTrainingJob#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#file_system_type AwsTrainingJob#file_system_type}.
	// Experimental.
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
}

