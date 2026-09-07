package sagemakerai


// Experimental.
type AwsDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#batch_transform_input AwsDataQualityJobDefinition#batch_transform_input}
	// Experimental.
	BatchTransformInput *AwsDataQualityJobDefinition_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#endpoint_input AwsDataQualityJobDefinition#endpoint_input}
	// Experimental.
	EndpointInput *AwsDataQualityJobDefinition_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

