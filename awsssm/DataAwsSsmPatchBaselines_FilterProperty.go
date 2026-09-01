package awsssm


// Experimental.
type DataAwsSsmPatchBaselines_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_patch_baselines#key DataAwsSsmPatchBaselines#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssm_patch_baselines#values DataAwsSsmPatchBaselines#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

