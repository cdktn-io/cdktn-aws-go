package awsce


// Experimental.
type TfCostCategory_RuleRuleOrTagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#key TfCostCategory#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#match_options TfCostCategory#match_options}.
	// Experimental.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#values TfCostCategory#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

