package awsec2imagebuilder


// Experimental.
type AwsImagebuilderLifecyclePolicy_PolicyDetailProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#action AwsImagebuilderLifecyclePolicy#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// exclusion_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#exclusion_rules AwsImagebuilderLifecyclePolicy#exclusion_rules}
	// Experimental.
	ExclusionRules interface{} `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#filter AwsImagebuilderLifecyclePolicy#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

