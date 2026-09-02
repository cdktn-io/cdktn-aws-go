package awsvpclattice


// Experimental.
type TfListenerRule_MatchHttpMatchHeaderMatchesMatchProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#contains TfListenerRule#contains}.
	// Experimental.
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#exact TfListenerRule#exact}.
	// Experimental.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#prefix TfListenerRule#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

