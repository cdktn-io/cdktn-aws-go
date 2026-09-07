package waf


// Experimental.
type AwsWebAcl_ChallengeConfigProperty struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#immunity_time_property AwsWebAcl#immunity_time_property}
	// Experimental.
	ImmunityTimeProperty *AwsWebAcl_ChallengeConfigImmunityTimePropertyProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

