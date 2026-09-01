package awssagemakerai


// Experimental.
type AwsSagemakerHyperParameterTuningJob_WarmStartConfigProperty struct {
	// parent_hyper_parameter_tuning_jobs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#parent_hyper_parameter_tuning_jobs AwsSagemakerHyperParameterTuningJob#parent_hyper_parameter_tuning_jobs}
	// Experimental.
	ParentHyperParameterTuningJobs interface{} `field:"optional" json:"parentHyperParameterTuningJobs" yaml:"parentHyperParameterTuningJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#warm_start_type AwsSagemakerHyperParameterTuningJob#warm_start_type}.
	// Experimental.
	WarmStartType *string `field:"optional" json:"warmStartType" yaml:"warmStartType"`
}

