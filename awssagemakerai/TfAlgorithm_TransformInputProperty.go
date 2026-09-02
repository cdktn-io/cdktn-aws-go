package awssagemakerai


// Experimental.
type TfAlgorithm_TransformInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#compression_type TfAlgorithm#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#content_type TfAlgorithm#content_type}.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#data_source TfAlgorithm#data_source}
	// Experimental.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#split_type TfAlgorithm#split_type}.
	// Experimental.
	SplitType *string `field:"optional" json:"splitType" yaml:"splitType"`
}

