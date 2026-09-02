package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_ResourceLimitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_parallel_training_jobs TfHyperParameterTuningJob#max_parallel_training_jobs}.
	// Experimental.
	MaxParallelTrainingJobs *float64 `field:"required" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_number_of_training_jobs TfHyperParameterTuningJob#max_number_of_training_jobs}.
	// Experimental.
	MaxNumberOfTrainingJobs *float64 `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_runtime_in_seconds TfHyperParameterTuningJob#max_runtime_in_seconds}.
	// Experimental.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

