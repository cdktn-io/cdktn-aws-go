package awsce


// Experimental.
type TfCostCategory_SplitChargeRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#method TfCostCategory#method}.
	// Experimental.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#source TfCostCategory#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#targets TfCostCategory#targets}.
	// Experimental.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#parameter TfCostCategory#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

