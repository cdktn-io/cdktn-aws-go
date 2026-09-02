package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_AmisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#is_public TfLifecyclePolicy#is_public}.
	// Experimental.
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// last_launched block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#last_launched TfLifecyclePolicy#last_launched}
	// Experimental.
	LastLaunched interface{} `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#regions TfLifecyclePolicy#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#shared_accounts TfLifecyclePolicy#shared_accounts}.
	// Experimental.
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map TfLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

