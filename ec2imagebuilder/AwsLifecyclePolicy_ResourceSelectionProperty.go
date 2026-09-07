package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_ResourceSelectionProperty struct {
	// recipe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#recipe AwsLifecyclePolicy#recipe}
	// Experimental.
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map AwsLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

