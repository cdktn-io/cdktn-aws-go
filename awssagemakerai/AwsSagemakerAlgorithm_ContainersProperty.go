package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_ContainersProperty struct {
	// additional_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#additional_s3_data_source AwsSagemakerAlgorithm#additional_s3_data_source}
	// Experimental.
	AdditionalS3DataSource interface{} `field:"optional" json:"additionalS3DataSource" yaml:"additionalS3DataSource"`
	// base_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#base_model AwsSagemakerAlgorithm#base_model}
	// Experimental.
	BaseModel interface{} `field:"optional" json:"baseModel" yaml:"baseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#container_hostname AwsSagemakerAlgorithm#container_hostname}.
	// Experimental.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#environment AwsSagemakerAlgorithm#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework AwsSagemakerAlgorithm#framework}.
	// Experimental.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#framework_version AwsSagemakerAlgorithm#framework_version}.
	// Experimental.
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image AwsSagemakerAlgorithm#image}.
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#image_digest AwsSagemakerAlgorithm#image_digest}.
	// Experimental.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_checkpoint AwsSagemakerAlgorithm#is_checkpoint}.
	// Experimental.
	IsCheckpoint interface{} `field:"optional" json:"isCheckpoint" yaml:"isCheckpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_etag AwsSagemakerAlgorithm#model_data_etag}.
	// Experimental.
	ModelDataEtag *string `field:"optional" json:"modelDataEtag" yaml:"modelDataEtag"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_source AwsSagemakerAlgorithm#model_data_source}
	// Experimental.
	ModelDataSource interface{} `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_data_url AwsSagemakerAlgorithm#model_data_url}.
	// Experimental.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// model_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#model_input AwsSagemakerAlgorithm#model_input}
	// Experimental.
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#nearest_model_name AwsSagemakerAlgorithm#nearest_model_name}.
	// Experimental.
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#product_id AwsSagemakerAlgorithm#product_id}.
	// Experimental.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

