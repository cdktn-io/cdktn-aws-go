package awsrolesanywhere


// Experimental.
type AwsRolesanywhereTrustAnchor_NotificationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#channel AwsRolesanywhereTrustAnchor#channel}.
	// Experimental.
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#enabled AwsRolesanywhereTrustAnchor#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#event AwsRolesanywhereTrustAnchor#event}.
	// Experimental.
	Event *string `field:"optional" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#threshold AwsRolesanywhereTrustAnchor#threshold}.
	// Experimental.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

