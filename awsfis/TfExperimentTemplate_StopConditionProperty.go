package awsfis


// Experimental.
type TfExperimentTemplate_StopConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#source TfExperimentTemplate#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#value TfExperimentTemplate#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

