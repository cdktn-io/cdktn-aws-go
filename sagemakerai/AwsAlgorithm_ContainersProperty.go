package sagemakerai


// Experimental.
type AwsAlgorithm_ContainersProperty struct {
	// additional_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#additional_s3_data_source AwsAlgorithm#additional_s3_data_source}
	// Experimental.
	AdditionalS3DataSource interface{} `field:"optional" json:"additionalS3DataSource" yaml:"additionalS3DataSource"`
	// base_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#base_model AwsAlgorithm#base_model}
	// Experimental.
	BaseModel interface{} `field:"optional" json:"baseModel" yaml:"baseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#container_hostname AwsAlgorithm#container_hostname}.
	// Experimental.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#environment AwsAlgorithm#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework AwsAlgorithm#framework}.
	// Experimental.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework_version AwsAlgorithm#framework_version}.
	// Experimental.
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image AwsAlgorithm#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image_digest AwsAlgorithm#image_digest}.
	// Experimental.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_checkpoint AwsAlgorithm#is_checkpoint}.
	// Experimental.
	IsCheckpoint interface{} `field:"optional" json:"isCheckpoint" yaml:"isCheckpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_etag AwsAlgorithm#model_data_etag}.
	// Experimental.
	ModelDataEtag *string `field:"optional" json:"modelDataEtag" yaml:"modelDataEtag"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_source AwsAlgorithm#model_data_source}
	// Experimental.
	ModelDataSource interface{} `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_url AwsAlgorithm#model_data_url}.
	// Experimental.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// model_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_input AwsAlgorithm#model_input}
	// Experimental.
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#nearest_model_name AwsAlgorithm#nearest_model_name}.
	// Experimental.
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#product_id AwsAlgorithm#product_id}.
	// Experimental.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

