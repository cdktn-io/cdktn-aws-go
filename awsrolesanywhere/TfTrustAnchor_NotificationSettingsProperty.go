package awsrolesanywhere


// Experimental.
type TfTrustAnchor_NotificationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#channel TfTrustAnchor#channel}.
	// Experimental.
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#enabled TfTrustAnchor#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#event TfTrustAnchor#event}.
	// Experimental.
	Event *string `field:"optional" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#threshold TfTrustAnchor#threshold}.
	// Experimental.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

