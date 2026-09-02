package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#role_arn TfHyperParameterTuningJob#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// algorithm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#algorithm_specification TfHyperParameterTuningJob#algorithm_specification}
	// Experimental.
	AlgorithmSpecification interface{} `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// checkpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#checkpoint_config TfHyperParameterTuningJob#checkpoint_config}
	// Experimental.
	CheckpointConfig interface{} `field:"optional" json:"checkpointConfig" yaml:"checkpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#definition_name TfHyperParameterTuningJob#definition_name}.
	// Experimental.
	DefinitionName *string `field:"optional" json:"definitionName" yaml:"definitionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_inter_container_traffic_encryption TfHyperParameterTuningJob#enable_inter_container_traffic_encryption}.
	// Experimental.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_managed_spot_training TfHyperParameterTuningJob#enable_managed_spot_training}.
	// Experimental.
	EnableManagedSpotTraining interface{} `field:"optional" json:"enableManagedSpotTraining" yaml:"enableManagedSpotTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_network_isolation TfHyperParameterTuningJob#enable_network_isolation}.
	// Experimental.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#environment TfHyperParameterTuningJob#environment}.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// hyper_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#hyper_parameter_ranges TfHyperParameterTuningJob#hyper_parameter_ranges}
	// Experimental.
	HyperParameterRanges interface{} `field:"optional" json:"hyperParameterRanges" yaml:"hyperParameterRanges"`
	// hyper_parameter_tuning_resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#hyper_parameter_tuning_resource_config TfHyperParameterTuningJob#hyper_parameter_tuning_resource_config}
	// Experimental.
	HyperParameterTuningResourceConfig interface{} `field:"optional" json:"hyperParameterTuningResourceConfig" yaml:"hyperParameterTuningResourceConfig"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#input_data_config TfHyperParameterTuningJob#input_data_config}
	// Experimental.
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#output_data_config TfHyperParameterTuningJob#output_data_config}
	// Experimental.
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#resource_config TfHyperParameterTuningJob#resource_config}
	// Experimental.
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#retry_strategy TfHyperParameterTuningJob#retry_strategy}.
	// Experimental.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#static_hyper_parameters TfHyperParameterTuningJob#static_hyper_parameters}.
	// Experimental.
	StaticHyperParameters *map[string]*string `field:"optional" json:"staticHyperParameters" yaml:"staticHyperParameters"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#stopping_condition TfHyperParameterTuningJob#stopping_condition}
	// Experimental.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// tuning_objective block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tuning_objective TfHyperParameterTuningJob#tuning_objective}
	// Experimental.
	TuningObjective interface{} `field:"optional" json:"tuningObjective" yaml:"tuningObjective"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#vpc_config TfHyperParameterTuningJob#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

