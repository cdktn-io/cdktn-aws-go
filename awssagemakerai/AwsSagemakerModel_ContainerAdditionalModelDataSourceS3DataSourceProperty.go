package awssagemakerai


// Experimental.
type AwsSagemakerModel_ContainerAdditionalModelDataSourceS3DataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#compression_type AwsSagemakerModel#compression_type}.
	// Experimental.
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#s3_data_type AwsSagemakerModel#s3_data_type}.
	// Experimental.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#s3_uri AwsSagemakerModel#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#model_access_config AwsSagemakerModel#model_access_config}
	// Experimental.
	ModelAccessConfig *AwsSagemakerModel_ContainerAdditionalModelDataSourceS3DataSourceModelAccessConfigProperty `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
}

