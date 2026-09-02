package awsvpclattice


// Experimental.
type TfListenerRule_HeaderMatchesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#match TfListenerRule#match}
	// Experimental.
	Match *TfListenerRule_MatchHttpMatchHeaderMatchesMatchProperty `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#name TfListenerRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#case_sensitive TfListenerRule#case_sensitive}.
	// Experimental.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

