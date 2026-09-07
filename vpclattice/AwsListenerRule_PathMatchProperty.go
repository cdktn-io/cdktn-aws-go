package vpclattice


// Experimental.
type AwsListenerRule_PathMatchProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#match AwsListenerRule#match}
	// Experimental.
	Match *AwsListenerRule_MatchHttpMatchPathMatchMatchProperty `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#case_sensitive AwsListenerRule#case_sensitive}.
	// Experimental.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

