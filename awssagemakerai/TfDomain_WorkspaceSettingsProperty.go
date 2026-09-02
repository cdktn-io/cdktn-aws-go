package awssagemakerai


// Experimental.
type TfDomain_WorkspaceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#s3_artifact_path TfDomain#s3_artifact_path}.
	// Experimental.
	S3ArtifactPath *string `field:"optional" json:"s3ArtifactPath" yaml:"s3ArtifactPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#s3_kms_key_id TfDomain#s3_kms_key_id}.
	// Experimental.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
}

