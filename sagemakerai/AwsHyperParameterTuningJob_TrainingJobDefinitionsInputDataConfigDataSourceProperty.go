package sagemakerai


// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigDataSourceProperty struct {
	// file_system_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#file_system_data_source AwsHyperParameterTuningJob#file_system_data_source}
	// Experimental.
	FileSystemDataSource interface{} `field:"optional" json:"fileSystemDataSource" yaml:"fileSystemDataSource"`
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_data_source AwsHyperParameterTuningJob#s3_data_source}
	// Experimental.
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

