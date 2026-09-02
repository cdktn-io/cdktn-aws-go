package awssagemakerai


// Experimental.
type TfFlowDefinition_OutputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#s3_output_path TfFlowDefinition#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#kms_key_id TfFlowDefinition#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

