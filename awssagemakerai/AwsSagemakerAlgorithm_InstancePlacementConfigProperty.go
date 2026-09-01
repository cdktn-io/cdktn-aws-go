package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_InstancePlacementConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#enable_multiple_jobs AwsSagemakerAlgorithm#enable_multiple_jobs}.
	// Experimental.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#placement_specifications AwsSagemakerAlgorithm#placement_specifications}
	// Experimental.
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

