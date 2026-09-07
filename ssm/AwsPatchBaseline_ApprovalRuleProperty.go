package ssm


// Experimental.
type AwsPatchBaseline_ApprovalRuleProperty struct {
	// patch_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#patch_filter AwsPatchBaseline#patch_filter}
	// Experimental.
	PatchFilter interface{} `field:"required" json:"patchFilter" yaml:"patchFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approve_after_days AwsPatchBaseline#approve_after_days}.
	// Experimental.
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approve_until_date AwsPatchBaseline#approve_until_date}.
	// Experimental.
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#compliance_level AwsPatchBaseline#compliance_level}.
	// Experimental.
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#enable_non_security AwsPatchBaseline#enable_non_security}.
	// Experimental.
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
}

