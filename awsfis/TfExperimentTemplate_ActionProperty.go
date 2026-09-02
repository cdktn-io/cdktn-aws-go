package awsfis


// Experimental.
type TfExperimentTemplate_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#action_id TfExperimentTemplate#action_id}.
	// Experimental.
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#name TfExperimentTemplate#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#description TfExperimentTemplate#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#parameter TfExperimentTemplate#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#start_after TfExperimentTemplate#start_after}.
	// Experimental.
	StartAfter *[]*string `field:"optional" json:"startAfter" yaml:"startAfter"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#target TfExperimentTemplate#target}
	// Experimental.
	Target *TfExperimentTemplate_ActionTargetProperty `field:"optional" json:"target" yaml:"target"`
}

