package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_PolicyDetailProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#action AwsLifecyclePolicy#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// exclusion_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#exclusion_rules AwsLifecyclePolicy#exclusion_rules}
	// Experimental.
	ExclusionRules interface{} `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#filter AwsLifecyclePolicy#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

