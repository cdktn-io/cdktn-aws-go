package awssagemakerai


// Experimental.
type TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_data_type TfAlgorithm#s3_data_type}.
	// Experimental.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_uri TfAlgorithm#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#attribute_names TfAlgorithm#attribute_names}.
	// Experimental.
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// hub_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_access_config TfAlgorithm#hub_access_config}
	// Experimental.
	HubAccessConfig interface{} `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_group_names TfAlgorithm#instance_group_names}.
	// Experimental.
	InstanceGroupNames *[]*string `field:"optional" json:"instanceGroupNames" yaml:"instanceGroupNames"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_access_config TfAlgorithm#model_access_config}
	// Experimental.
	ModelAccessConfig interface{} `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#s3_data_distribution_type TfAlgorithm#s3_data_distribution_type}.
	// Experimental.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
}

