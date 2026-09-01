package awssagemakerai


// Experimental.
type AwsSagemakerEndpointConfiguration_AsyncInferenceConfigProperty struct {
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#output_config AwsSagemakerEndpointConfiguration#output_config}
	// Experimental.
	OutputConfig *AwsSagemakerEndpointConfiguration_OutputConfigProperty `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#client_config AwsSagemakerEndpointConfiguration#client_config}
	// Experimental.
	ClientConfig *AwsSagemakerEndpointConfiguration_ClientConfigProperty `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

