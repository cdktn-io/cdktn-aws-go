package awssagemakerai


// Experimental.
type TfAlgorithm_ContainersProperty struct {
	// additional_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#additional_s3_data_source TfAlgorithm#additional_s3_data_source}
	// Experimental.
	AdditionalS3DataSource interface{} `field:"optional" json:"additionalS3DataSource" yaml:"additionalS3DataSource"`
	// base_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#base_model TfAlgorithm#base_model}
	// Experimental.
	BaseModel interface{} `field:"optional" json:"baseModel" yaml:"baseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#container_hostname TfAlgorithm#container_hostname}.
	// Experimental.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#environment TfAlgorithm#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework TfAlgorithm#framework}.
	// Experimental.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework_version TfAlgorithm#framework_version}.
	// Experimental.
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image TfAlgorithm#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image_digest TfAlgorithm#image_digest}.
	// Experimental.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_checkpoint TfAlgorithm#is_checkpoint}.
	// Experimental.
	IsCheckpoint interface{} `field:"optional" json:"isCheckpoint" yaml:"isCheckpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_etag TfAlgorithm#model_data_etag}.
	// Experimental.
	ModelDataEtag *string `field:"optional" json:"modelDataEtag" yaml:"modelDataEtag"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_source TfAlgorithm#model_data_source}
	// Experimental.
	ModelDataSource interface{} `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_url TfAlgorithm#model_data_url}.
	// Experimental.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// model_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_input TfAlgorithm#model_input}
	// Experimental.
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#nearest_model_name TfAlgorithm#nearest_model_name}.
	// Experimental.
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#product_id TfAlgorithm#product_id}.
	// Experimental.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

