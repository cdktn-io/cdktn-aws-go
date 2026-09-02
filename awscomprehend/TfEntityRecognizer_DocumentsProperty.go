package awscomprehend


// Experimental.
type TfEntityRecognizer_DocumentsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#s3_uri TfEntityRecognizer#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#input_format TfEntityRecognizer#input_format}.
	// Experimental.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/comprehend_entity_recognizer#test_s3_uri TfEntityRecognizer#test_s3_uri}.
	// Experimental.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

