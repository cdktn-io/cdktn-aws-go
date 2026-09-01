package awsfis


// Experimental.
type AwsFisExperimentTemplate_ExperimentOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#account_targeting AwsFisExperimentTemplate#account_targeting}.
	// Experimental.
	AccountTargeting *string `field:"optional" json:"accountTargeting" yaml:"accountTargeting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#empty_target_resolution_mode AwsFisExperimentTemplate#empty_target_resolution_mode}.
	// Experimental.
	EmptyTargetResolutionMode *string `field:"optional" json:"emptyTargetResolutionMode" yaml:"emptyTargetResolutionMode"`
}

