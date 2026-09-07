package iam


// Experimental.
type DataAwsPrincipalPolicySimulation_ContextProperty struct {
	// The key name of the context entry, such as "aws:CurrentTime".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation#key DataAwsPrincipalPolicySimulation#key}
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The type that the simulator should use to interpret the strings given in argument "values".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation#type DataAwsPrincipalPolicySimulation#type}
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// One or more values to assign to the context key, given as a string in a syntax appropriate for the selected value type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_principal_policy_simulation#values DataAwsPrincipalPolicySimulation#values}
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

