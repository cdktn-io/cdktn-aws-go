package awsvpclattice


// Experimental.
type AwsVpclatticeListenerRule_MatchProperty struct {
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#http_match AwsVpclatticeListenerRule#http_match}
	// Experimental.
	HttpMatch *AwsVpclatticeListenerRule_HttpMatchProperty `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

