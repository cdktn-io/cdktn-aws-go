package awsssm


// Experimental.
type TfPatchBaseline_GlobalFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#key TfPatchBaseline#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#values TfPatchBaseline#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

