package awssagemakerai


// Experimental.
type TfFeatureGroup_S3StorageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#s3_uri TfFeatureGroup#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#kms_key_id TfFeatureGroup#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#resolved_output_s3_uri TfFeatureGroup#resolved_output_s3_uri}.
	// Experimental.
	ResolvedOutputS3Uri *string `field:"optional" json:"resolvedOutputS3Uri" yaml:"resolvedOutputS3Uri"`
}

