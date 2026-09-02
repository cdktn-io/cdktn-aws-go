package awscomprehend


// Experimental.
type TfDocumentClassifier_InputDataConfigProperty struct {
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#augmented_manifests TfDocumentClassifier#augmented_manifests}
	// Experimental.
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#data_format TfDocumentClassifier#data_format}.
	// Experimental.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#label_delimiter TfDocumentClassifier#label_delimiter}.
	// Experimental.
	LabelDelimiter *string `field:"optional" json:"labelDelimiter" yaml:"labelDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#s3_uri TfDocumentClassifier#s3_uri}.
	// Experimental.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#test_s3_uri TfDocumentClassifier#test_s3_uri}.
	// Experimental.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

