package awsssm


// Experimental.
type AwsSsmPatchBaseline_ApprovalRuleProperty struct {
	// patch_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#patch_filter AwsSsmPatchBaseline#patch_filter}
	// Experimental.
	PatchFilter interface{} `field:"required" json:"patchFilter" yaml:"patchFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approve_after_days AwsSsmPatchBaseline#approve_after_days}.
	// Experimental.
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#approve_until_date AwsSsmPatchBaseline#approve_until_date}.
	// Experimental.
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#compliance_level AwsSsmPatchBaseline#compliance_level}.
	// Experimental.
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#enable_non_security AwsSsmPatchBaseline#enable_non_security}.
	// Experimental.
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
}

