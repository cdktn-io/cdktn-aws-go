package awsvpclattice


// Experimental.
type TfListenerRule_MatchProperty struct {
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#http_match TfListenerRule#http_match}
	// Experimental.
	HttpMatch *TfListenerRule_HttpMatchProperty `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

