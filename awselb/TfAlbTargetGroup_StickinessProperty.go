package awselb


// Experimental.
type TfAlbTargetGroup_StickinessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#type TfAlbTargetGroup#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#cookie_duration TfAlbTargetGroup#cookie_duration}.
	// Experimental.
	CookieDuration *float64 `field:"optional" json:"cookieDuration" yaml:"cookieDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#cookie_name TfAlbTargetGroup#cookie_name}.
	// Experimental.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#enabled TfAlbTargetGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

