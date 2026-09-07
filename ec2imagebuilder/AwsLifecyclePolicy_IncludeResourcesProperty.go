package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_IncludeResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#amis AwsLifecyclePolicy#amis}.
	// Experimental.
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#containers AwsLifecyclePolicy#containers}.
	// Experimental.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#snapshots AwsLifecyclePolicy#snapshots}.
	// Experimental.
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

