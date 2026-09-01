package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_EndpointInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#endpoint_name AwsSagemakerDataQualityJobDefinition#endpoint_name}.
	// Experimental.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#local_path AwsSagemakerDataQualityJobDefinition#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#s3_data_distribution_type AwsSagemakerDataQualityJobDefinition#s3_data_distribution_type}.
	// Experimental.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#s3_input_mode AwsSagemakerDataQualityJobDefinition#s3_input_mode}.
	// Experimental.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

