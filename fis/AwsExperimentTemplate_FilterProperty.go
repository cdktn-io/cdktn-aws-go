package fis


// Experimental.
type AwsExperimentTemplate_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#path AwsExperimentTemplate#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#values AwsExperimentTemplate#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

