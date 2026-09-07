package sagemakerai


// Experimental.
type AwsEndpointConfiguration_AsyncInferenceConfigProperty struct {
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#output_config AwsEndpointConfiguration#output_config}
	// Experimental.
	OutputConfig *AwsEndpointConfiguration_OutputConfigProperty `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#client_config AwsEndpointConfiguration#client_config}
	// Experimental.
	ClientConfig *AwsEndpointConfiguration_ClientConfigProperty `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

