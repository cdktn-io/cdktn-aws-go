package awssagemakerai


// Experimental.
type TfModel_PrimaryContainerProperty struct {
	// additional_model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#additional_model_data_source TfModel#additional_model_data_source}
	// Experimental.
	AdditionalModelDataSource interface{} `field:"optional" json:"additionalModelDataSource" yaml:"additionalModelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#container_hostname TfModel#container_hostname}.
	// Experimental.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#environment TfModel#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#image TfModel#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#image_config TfModel#image_config}
	// Experimental.
	ImageConfig *TfModel_PrimaryContainerImageConfigProperty `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#inference_specification_name TfModel#inference_specification_name}.
	// Experimental.
	InferenceSpecificationName *string `field:"optional" json:"inferenceSpecificationName" yaml:"inferenceSpecificationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#mode TfModel#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#model_data_source TfModel#model_data_source}
	// Experimental.
	ModelDataSource *TfModel_PrimaryContainerModelDataSourceProperty `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#model_data_url TfModel#model_data_url}.
	// Experimental.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#model_package_name TfModel#model_package_name}.
	// Experimental.
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// multi_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#multi_model_config TfModel#multi_model_config}
	// Experimental.
	MultiModelConfig *TfModel_PrimaryContainerMultiModelConfigProperty `field:"optional" json:"multiModelConfig" yaml:"multiModelConfig"`
}

