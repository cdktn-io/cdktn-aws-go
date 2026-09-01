package awssagemakerai


// Experimental.
type AwsSagemakerHyperParameterTuningJob_ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy AwsSagemakerHyperParameterTuningJob#strategy}.
	// Experimental.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// objective block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#objective AwsSagemakerHyperParameterTuningJob#objective}
	// Experimental.
	Objective interface{} `field:"optional" json:"objective" yaml:"objective"`
	// parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#parameter_ranges AwsSagemakerHyperParameterTuningJob#parameter_ranges}
	// Experimental.
	ParameterRanges interface{} `field:"optional" json:"parameterRanges" yaml:"parameterRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#random_seed AwsSagemakerHyperParameterTuningJob#random_seed}.
	// Experimental.
	RandomSeed *float64 `field:"optional" json:"randomSeed" yaml:"randomSeed"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#resource_limits AwsSagemakerHyperParameterTuningJob#resource_limits}
	// Experimental.
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// strategy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy_config AwsSagemakerHyperParameterTuningJob#strategy_config}
	// Experimental.
	StrategyConfig interface{} `field:"optional" json:"strategyConfig" yaml:"strategyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_job_early_stopping_type AwsSagemakerHyperParameterTuningJob#training_job_early_stopping_type}.
	// Experimental.
	TrainingJobEarlyStoppingType *string `field:"optional" json:"trainingJobEarlyStoppingType" yaml:"trainingJobEarlyStoppingType"`
	// tuning_job_completion_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tuning_job_completion_criteria AwsSagemakerHyperParameterTuningJob#tuning_job_completion_criteria}
	// Experimental.
	TuningJobCompletionCriteria interface{} `field:"optional" json:"tuningJobCompletionCriteria" yaml:"tuningJobCompletionCriteria"`
}

