package awssagemakerai


// Experimental.
type AwsSagemakerDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#batch_transform_input AwsSagemakerDataQualityJobDefinition#batch_transform_input}
	// Experimental.
	BatchTransformInput *AwsSagemakerDataQualityJobDefinition_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#endpoint_input AwsSagemakerDataQualityJobDefinition#endpoint_input}
	// Experimental.
	EndpointInput *AwsSagemakerDataQualityJobDefinition_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

