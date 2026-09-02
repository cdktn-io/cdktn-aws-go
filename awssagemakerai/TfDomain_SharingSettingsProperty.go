package awssagemakerai


// Experimental.
type TfDomain_SharingSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#notebook_output_option TfDomain#notebook_output_option}.
	// Experimental.
	NotebookOutputOption *string `field:"optional" json:"notebookOutputOption" yaml:"notebookOutputOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#s3_kms_key_id TfDomain#s3_kms_key_id}.
	// Experimental.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#s3_output_path TfDomain#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

