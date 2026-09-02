package awscomprehend


// Experimental.
type TfDocumentClassifier_OutputDataConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#s3_uri TfDocumentClassifier#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#kms_key_id TfDocumentClassifier#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

