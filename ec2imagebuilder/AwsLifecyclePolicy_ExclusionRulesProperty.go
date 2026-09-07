package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_ExclusionRulesProperty struct {
	// amis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#amis AwsLifecyclePolicy#amis}
	// Experimental.
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map AwsLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

