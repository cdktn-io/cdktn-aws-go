package sagemakerai


// Experimental.
type AwsModel_PrimaryContainerModelDataSourceS3DataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#compression_type AwsModel#compression_type}.
	// Experimental.
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#s3_data_type AwsModel#s3_data_type}.
	// Experimental.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#s3_uri AwsModel#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#model_access_config AwsModel#model_access_config}
	// Experimental.
	ModelAccessConfig *AwsModel_PrimaryContainerModelDataSourceS3DataSourceModelAccessConfigProperty `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
}

