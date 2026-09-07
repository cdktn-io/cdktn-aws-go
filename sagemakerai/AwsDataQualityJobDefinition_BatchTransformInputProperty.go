package sagemakerai


// Experimental.
type AwsDataQualityJobDefinition_BatchTransformInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#data_captured_destination_s3_uri AwsDataQualityJobDefinition#data_captured_destination_s3_uri}.
	// Experimental.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// dataset_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#dataset_format AwsDataQualityJobDefinition#dataset_format}
	// Experimental.
	DatasetFormat *AwsDataQualityJobDefinition_DatasetFormatProperty `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#local_path AwsDataQualityJobDefinition#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#s3_data_distribution_type AwsDataQualityJobDefinition#s3_data_distribution_type}.
	// Experimental.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#s3_input_mode AwsDataQualityJobDefinition#s3_input_mode}.
	// Experimental.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

