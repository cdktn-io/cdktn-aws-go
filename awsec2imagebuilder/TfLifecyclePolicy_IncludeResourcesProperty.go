package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_IncludeResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#amis TfLifecyclePolicy#amis}.
	// Experimental.
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#containers TfLifecyclePolicy#containers}.
	// Experimental.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#snapshots TfLifecyclePolicy#snapshots}.
	// Experimental.
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

