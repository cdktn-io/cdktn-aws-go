package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_ExclusionRulesProperty struct {
	// amis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#amis TfLifecyclePolicy#amis}
	// Experimental.
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map TfLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

