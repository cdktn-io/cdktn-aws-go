package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_ResourceSelectionProperty struct {
	// recipe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#recipe TfLifecyclePolicy#recipe}
	// Experimental.
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map TfLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

