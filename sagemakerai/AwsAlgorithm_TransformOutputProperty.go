package sagemakerai


// Experimental.
type AwsAlgorithm_TransformOutputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_output_path AwsAlgorithm#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#accept AwsAlgorithm#accept}.
	// Experimental.
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#assemble_with AwsAlgorithm#assemble_with}.
	// Experimental.
	AssembleWith *string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#kms_key_id AwsAlgorithm#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

