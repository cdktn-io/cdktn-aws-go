package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_multiple_jobs TfHyperParameterTuningJob#enable_multiple_jobs}.
	// Experimental.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#placement_specifications TfHyperParameterTuningJob#placement_specifications}
	// Experimental.
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

