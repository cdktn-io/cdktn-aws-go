package sagemakerai


// Experimental.
type AwsEndpointConfiguration_NotificationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#error_topic AwsEndpointConfiguration#error_topic}.
	// Experimental.
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#include_inference_response_in AwsEndpointConfiguration#include_inference_response_in}.
	// Experimental.
	IncludeInferenceResponseIn *[]*string `field:"optional" json:"includeInferenceResponseIn" yaml:"includeInferenceResponseIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#success_topic AwsEndpointConfiguration#success_topic}.
	// Experimental.
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

