package awssagemakerai


// Experimental.
type TfAlgorithm_TransformJobDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#batch_strategy TfAlgorithm#batch_strategy}.
	// Experimental.
	BatchStrategy *string `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#environment TfAlgorithm#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_concurrent_transforms TfAlgorithm#max_concurrent_transforms}.
	// Experimental.
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_payload_in_mb TfAlgorithm#max_payload_in_mb}.
	// Experimental.
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#transform_input TfAlgorithm#transform_input}
	// Experimental.
	TransformInput interface{} `field:"optional" json:"transformInput" yaml:"transformInput"`
	// transform_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#transform_output TfAlgorithm#transform_output}
	// Experimental.
	TransformOutput interface{} `field:"optional" json:"transformOutput" yaml:"transformOutput"`
	// transform_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#transform_resources TfAlgorithm#transform_resources}
	// Experimental.
	TransformResources interface{} `field:"optional" json:"transformResources" yaml:"transformResources"`
}

