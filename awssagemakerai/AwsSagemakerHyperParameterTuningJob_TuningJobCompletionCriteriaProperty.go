package awssagemakerai


// Experimental.
type AwsSagemakerHyperParameterTuningJob_TuningJobCompletionCriteriaProperty struct {
	// best_objective_not_improving block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#best_objective_not_improving AwsSagemakerHyperParameterTuningJob#best_objective_not_improving}
	// Experimental.
	BestObjectiveNotImproving interface{} `field:"optional" json:"bestObjectiveNotImproving" yaml:"bestObjectiveNotImproving"`
	// convergence_detected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#convergence_detected AwsSagemakerHyperParameterTuningJob#convergence_detected}
	// Experimental.
	ConvergenceDetected interface{} `field:"optional" json:"convergenceDetected" yaml:"convergenceDetected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#target_objective_metric_value AwsSagemakerHyperParameterTuningJob#target_objective_metric_value}.
	// Experimental.
	TargetObjectiveMetricValue *float64 `field:"optional" json:"targetObjectiveMetricValue" yaml:"targetObjectiveMetricValue"`
}

