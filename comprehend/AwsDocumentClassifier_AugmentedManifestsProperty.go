package comprehend


// Experimental.
type AwsDocumentClassifier_AugmentedManifestsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#attribute_names AwsDocumentClassifier#attribute_names}.
	// Experimental.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#s3_uri AwsDocumentClassifier#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#annotation_data_s3_uri AwsDocumentClassifier#annotation_data_s3_uri}.
	// Experimental.
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#document_type AwsDocumentClassifier#document_type}.
	// Experimental.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#source_documents_s3_uri AwsDocumentClassifier#source_documents_s3_uri}.
	// Experimental.
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_document_classifier#split AwsDocumentClassifier#split}.
	// Experimental.
	Split *string `field:"optional" json:"split" yaml:"split"`
}

