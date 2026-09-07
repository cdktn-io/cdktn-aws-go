package vpclattice


// Experimental.
type AwsListenerRule_HttpMatchProperty struct {
	// header_matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#header_matches AwsListenerRule#header_matches}
	// Experimental.
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#method AwsListenerRule#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#path_match AwsListenerRule#path_match}
	// Experimental.
	PathMatch *AwsListenerRule_PathMatchProperty `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

