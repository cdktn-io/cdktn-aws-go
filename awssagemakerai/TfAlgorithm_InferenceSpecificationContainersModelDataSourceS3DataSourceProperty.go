package awssagemakerai


// Experimental.
type TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#compression_type TfAlgorithm#compression_type}.
	// Experimental.
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_data_type TfAlgorithm#s3_data_type}.
	// Experimental.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_uri TfAlgorithm#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#etag TfAlgorithm#etag}.
	// Experimental.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// hub_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_access_config TfAlgorithm#hub_access_config}
	// Experimental.
	HubAccessConfig interface{} `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#manifest_etag TfAlgorithm#manifest_etag}.
	// Experimental.
	ManifestEtag *string `field:"optional" json:"manifestEtag" yaml:"manifestEtag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#manifest_s3_uri TfAlgorithm#manifest_s3_uri}.
	// Experimental.
	ManifestS3Uri *string `field:"optional" json:"manifestS3Uri" yaml:"manifestS3Uri"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_access_config TfAlgorithm#model_access_config}
	// Experimental.
	ModelAccessConfig interface{} `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
}

