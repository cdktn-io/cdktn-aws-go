package awssagemakerai


// Experimental.
type TfEndpointConfiguration_DataCaptureConfigProperty struct {
	// capture_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capture_options TfEndpointConfiguration#capture_options}
	// Experimental.
	CaptureOptions interface{} `field:"required" json:"captureOptions" yaml:"captureOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#destination_s3_uri TfEndpointConfiguration#destination_s3_uri}.
	// Experimental.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_sampling_percentage TfEndpointConfiguration#initial_sampling_percentage}.
	// Experimental.
	InitialSamplingPercentage *float64 `field:"required" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// capture_content_type_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capture_content_type_header TfEndpointConfiguration#capture_content_type_header}
	// Experimental.
	CaptureContentTypeHeader *TfEndpointConfiguration_CaptureContentTypeHeaderProperty `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#enable_capture TfEndpointConfiguration#enable_capture}.
	// Experimental.
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id TfEndpointConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

