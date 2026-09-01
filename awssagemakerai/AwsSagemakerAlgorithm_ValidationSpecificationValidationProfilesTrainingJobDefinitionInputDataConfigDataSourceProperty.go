package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceProperty struct {
	// file_system_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#file_system_data_source AwsSagemakerAlgorithm#file_system_data_source}
	// Experimental.
	FileSystemDataSource interface{} `field:"optional" json:"fileSystemDataSource" yaml:"fileSystemDataSource"`
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_data_source AwsSagemakerAlgorithm#s3_data_source}
	// Experimental.
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

