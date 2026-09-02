package awselb


// Experimental.
type TfTargetGroup_StickinessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#type TfTargetGroup#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#cookie_duration TfTargetGroup#cookie_duration}.
	// Experimental.
	CookieDuration *float64 `field:"optional" json:"cookieDuration" yaml:"cookieDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#cookie_name TfTargetGroup#cookie_name}.
	// Experimental.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#enabled TfTargetGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

