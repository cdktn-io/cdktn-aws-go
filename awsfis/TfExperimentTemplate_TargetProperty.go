package awsfis


// Experimental.
type TfExperimentTemplate_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#name TfExperimentTemplate#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#resource_type TfExperimentTemplate#resource_type}.
	// Experimental.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#selection_mode TfExperimentTemplate#selection_mode}.
	// Experimental.
	SelectionMode *string `field:"required" json:"selectionMode" yaml:"selectionMode"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#filter TfExperimentTemplate#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#parameters TfExperimentTemplate#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#resource_arns TfExperimentTemplate#resource_arns}.
	// Experimental.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// resource_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fis_experiment_template#resource_tag TfExperimentTemplate#resource_tag}
	// Experimental.
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
}

