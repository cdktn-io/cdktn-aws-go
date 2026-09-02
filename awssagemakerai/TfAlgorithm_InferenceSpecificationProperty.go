package awssagemakerai


// Experimental.
type TfAlgorithm_InferenceSpecificationProperty struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#containers TfAlgorithm#containers}
	// Experimental.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_content_types TfAlgorithm#supported_content_types}.
	// Experimental.
	SupportedContentTypes *[]*string `field:"optional" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_realtime_inference_instance_types TfAlgorithm#supported_realtime_inference_instance_types}.
	// Experimental.
	SupportedRealtimeInferenceInstanceTypes *[]*string `field:"optional" json:"supportedRealtimeInferenceInstanceTypes" yaml:"supportedRealtimeInferenceInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_response_mime_types TfAlgorithm#supported_response_mime_types}.
	// Experimental.
	SupportedResponseMimeTypes *[]*string `field:"optional" json:"supportedResponseMimeTypes" yaml:"supportedResponseMimeTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_transform_instance_types TfAlgorithm#supported_transform_instance_types}.
	// Experimental.
	SupportedTransformInstanceTypes *[]*string `field:"optional" json:"supportedTransformInstanceTypes" yaml:"supportedTransformInstanceTypes"`
}

