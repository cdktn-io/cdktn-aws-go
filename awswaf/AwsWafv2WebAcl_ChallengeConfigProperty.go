package awswaf


// Experimental.
type AwsWafv2WebAcl_ChallengeConfigProperty struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#immunity_time_property AwsWafv2WebAcl#immunity_time_property}
	// Experimental.
	ImmunityTimeProperty *AwsWafv2WebAcl_ChallengeConfigImmunityTimePropertyProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

