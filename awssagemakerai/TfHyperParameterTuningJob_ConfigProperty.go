package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy TfHyperParameterTuningJob#strategy}.
	// Experimental.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// objective block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#objective TfHyperParameterTuningJob#objective}
	// Experimental.
	Objective interface{} `field:"optional" json:"objective" yaml:"objective"`
	// parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#parameter_ranges TfHyperParameterTuningJob#parameter_ranges}
	// Experimental.
	ParameterRanges interface{} `field:"optional" json:"parameterRanges" yaml:"parameterRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#random_seed TfHyperParameterTuningJob#random_seed}.
	// Experimental.
	RandomSeed *float64 `field:"optional" json:"randomSeed" yaml:"randomSeed"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#resource_limits TfHyperParameterTuningJob#resource_limits}
	// Experimental.
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// strategy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy_config TfHyperParameterTuningJob#strategy_config}
	// Experimental.
	StrategyConfig interface{} `field:"optional" json:"strategyConfig" yaml:"strategyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_job_early_stopping_type TfHyperParameterTuningJob#training_job_early_stopping_type}.
	// Experimental.
	TrainingJobEarlyStoppingType *string `field:"optional" json:"trainingJobEarlyStoppingType" yaml:"trainingJobEarlyStoppingType"`
	// tuning_job_completion_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tuning_job_completion_criteria TfHyperParameterTuningJob#tuning_job_completion_criteria}
	// Experimental.
	TuningJobCompletionCriteria interface{} `field:"optional" json:"tuningJobCompletionCriteria" yaml:"tuningJobCompletionCriteria"`
}

