package awsec2imagebuilder


// Experimental.
type AwsImagebuilderLifecyclePolicy_AmisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#is_public AwsImagebuilderLifecyclePolicy#is_public}.
	// Experimental.
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// last_launched block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#last_launched AwsImagebuilderLifecyclePolicy#last_launched}
	// Experimental.
	LastLaunched interface{} `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#regions AwsImagebuilderLifecyclePolicy#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#shared_accounts AwsImagebuilderLifecyclePolicy#shared_accounts}.
	// Experimental.
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#tag_map AwsImagebuilderLifecyclePolicy#tag_map}.
	// Experimental.
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

