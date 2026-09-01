package awsvpclattice


// Experimental.
type AwsVpclatticeListenerRule_HttpMatchProperty struct {
	// header_matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#header_matches AwsVpclatticeListenerRule#header_matches}
	// Experimental.
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#method AwsVpclatticeListenerRule#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#path_match AwsVpclatticeListenerRule#path_match}
	// Experimental.
	PathMatch *AwsVpclatticeListenerRule_PathMatchProperty `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

