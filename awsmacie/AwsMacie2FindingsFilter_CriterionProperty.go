package awsmacie


// Experimental.
type AwsMacie2FindingsFilter_CriterionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#field AwsMacie2FindingsFilter#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#eq AwsMacie2FindingsFilter#eq}.
	// Experimental.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#eq_exact_match AwsMacie2FindingsFilter#eq_exact_match}.
	// Experimental.
	EqExactMatch *[]*string `field:"optional" json:"eqExactMatch" yaml:"eqExactMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#gt AwsMacie2FindingsFilter#gt}.
	// Experimental.
	Gt *string `field:"optional" json:"gt" yaml:"gt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#gte AwsMacie2FindingsFilter#gte}.
	// Experimental.
	Gte *string `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#lt AwsMacie2FindingsFilter#lt}.
	// Experimental.
	Lt *string `field:"optional" json:"lt" yaml:"lt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#lte AwsMacie2FindingsFilter#lte}.
	// Experimental.
	Lte *string `field:"optional" json:"lte" yaml:"lte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_findings_filter#neq AwsMacie2FindingsFilter#neq}.
	// Experimental.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

