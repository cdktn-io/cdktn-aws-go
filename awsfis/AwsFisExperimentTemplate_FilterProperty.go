package awsfis


// Experimental.
type AwsFisExperimentTemplate_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#path AwsFisExperimentTemplate#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#values AwsFisExperimentTemplate#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

