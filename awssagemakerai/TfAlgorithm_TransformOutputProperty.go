package awssagemakerai


// Experimental.
type TfAlgorithm_TransformOutputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_output_path TfAlgorithm#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#accept TfAlgorithm#accept}.
	// Experimental.
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#assemble_with TfAlgorithm#assemble_with}.
	// Experimental.
	AssembleWith *string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#kms_key_id TfAlgorithm#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

