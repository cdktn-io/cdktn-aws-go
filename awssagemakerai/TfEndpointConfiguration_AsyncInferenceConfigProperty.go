package awssagemakerai


// Experimental.
type TfEndpointConfiguration_AsyncInferenceConfigProperty struct {
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#output_config TfEndpointConfiguration#output_config}
	// Experimental.
	OutputConfig *TfEndpointConfiguration_OutputConfigProperty `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#client_config TfEndpointConfiguration#client_config}
	// Experimental.
	ClientConfig *TfEndpointConfiguration_ClientConfigProperty `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

