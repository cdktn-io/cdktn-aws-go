package awssagemakerai


// Experimental.
type TfAlgorithm_TrainingSpecificationAdditionalS3DataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_data_type TfAlgorithm#s3_data_type}.
	// Experimental.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_uri TfAlgorithm#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#compression_type TfAlgorithm#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#etag TfAlgorithm#etag}.
	// Experimental.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
}

