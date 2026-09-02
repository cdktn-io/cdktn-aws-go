package awssagemakerai


// Experimental.
type TfDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// batch_transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#batch_transform_input TfDataQualityJobDefinition#batch_transform_input}
	// Experimental.
	BatchTransformInput *TfDataQualityJobDefinition_BatchTransformInputProperty `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// endpoint_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_data_quality_job_definition#endpoint_input TfDataQualityJobDefinition#endpoint_input}
	// Experimental.
	EndpointInput *TfDataQualityJobDefinition_EndpointInputProperty `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

