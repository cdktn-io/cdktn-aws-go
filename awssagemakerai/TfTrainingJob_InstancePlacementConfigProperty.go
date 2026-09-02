package awssagemakerai


// Experimental.
type TfTrainingJob_InstancePlacementConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_multiple_jobs TfTrainingJob#enable_multiple_jobs}.
	// Experimental.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#placement_specifications TfTrainingJob#placement_specifications}
	// Experimental.
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

