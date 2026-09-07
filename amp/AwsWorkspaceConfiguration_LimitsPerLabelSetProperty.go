package amp


// Experimental.
type AwsWorkspaceConfiguration_LimitsPerLabelSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#label_set AwsWorkspaceConfiguration#label_set}.
	// Experimental.
	LabelSet *map[string]*string `field:"required" json:"labelSet" yaml:"labelSet"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#limits AwsWorkspaceConfiguration#limits}
	// Experimental.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
}

