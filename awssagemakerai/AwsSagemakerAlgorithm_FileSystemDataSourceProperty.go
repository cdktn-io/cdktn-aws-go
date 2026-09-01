package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_FileSystemDataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#directory_path AwsSagemakerAlgorithm#directory_path}.
	// Experimental.
	DirectoryPath *string `field:"required" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#file_system_access_mode AwsSagemakerAlgorithm#file_system_access_mode}.
	// Experimental.
	FileSystemAccessMode *string `field:"required" json:"fileSystemAccessMode" yaml:"fileSystemAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#file_system_id AwsSagemakerAlgorithm#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#file_system_type AwsSagemakerAlgorithm#file_system_type}.
	// Experimental.
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
}

