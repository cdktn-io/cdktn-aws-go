package awscomprehend


// Experimental.
type AwsComprehendEntityRecognizer_AugmentedManifestsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#attribute_names AwsComprehendEntityRecognizer#attribute_names}.
	// Experimental.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#s3_uri AwsComprehendEntityRecognizer#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#annotation_data_s3_uri AwsComprehendEntityRecognizer#annotation_data_s3_uri}.
	// Experimental.
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#document_type AwsComprehendEntityRecognizer#document_type}.
	// Experimental.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#source_documents_s3_uri AwsComprehendEntityRecognizer#source_documents_s3_uri}.
	// Experimental.
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#split AwsComprehendEntityRecognizer#split}.
	// Experimental.
	Split *string `field:"optional" json:"split" yaml:"split"`
}

