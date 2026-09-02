package awsvpclattice


// Experimental.
type TfListenerRule_PathMatchProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#match TfListenerRule#match}
	// Experimental.
	Match *TfListenerRule_MatchHttpMatchPathMatchMatchProperty `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#case_sensitive TfListenerRule#case_sensitive}.
	// Experimental.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

