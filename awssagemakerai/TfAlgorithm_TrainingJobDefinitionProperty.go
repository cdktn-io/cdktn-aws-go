package awssagemakerai


// Experimental.
type TfAlgorithm_TrainingJobDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_input_mode TfAlgorithm#training_input_mode}.
	// Experimental.
	TrainingInputMode *string `field:"required" json:"trainingInputMode" yaml:"trainingInputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hyper_parameters TfAlgorithm#hyper_parameters}.
	// Experimental.
	HyperParameters *map[string]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#input_data_config TfAlgorithm#input_data_config}
	// Experimental.
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#output_data_config TfAlgorithm#output_data_config}
	// Experimental.
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#resource_config TfAlgorithm#resource_config}
	// Experimental.
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#stopping_condition TfAlgorithm#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

