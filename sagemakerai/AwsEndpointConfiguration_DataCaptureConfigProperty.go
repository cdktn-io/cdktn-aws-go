package sagemakerai


// Experimental.
type AwsEndpointConfiguration_DataCaptureConfigProperty struct {
	// capture_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capture_options AwsEndpointConfiguration#capture_options}
	// Experimental.
	CaptureOptions interface{} `field:"required" json:"captureOptions" yaml:"captureOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#destination_s3_uri AwsEndpointConfiguration#destination_s3_uri}.
	// Experimental.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#initial_sampling_percentage AwsEndpointConfiguration#initial_sampling_percentage}.
	// Experimental.
	InitialSamplingPercentage *float64 `field:"required" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// capture_content_type_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capture_content_type_header AwsEndpointConfiguration#capture_content_type_header}
	// Experimental.
	CaptureContentTypeHeader *AwsEndpointConfiguration_CaptureContentTypeHeaderProperty `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#enable_capture AwsEndpointConfiguration#enable_capture}.
	// Experimental.
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id AwsEndpointConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

