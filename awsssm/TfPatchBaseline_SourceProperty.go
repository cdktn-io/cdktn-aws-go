package awsssm


// Experimental.
type TfPatchBaseline_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#configuration TfPatchBaseline#configuration}.
	// Experimental.
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#name TfPatchBaseline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#products TfPatchBaseline#products}.
	// Experimental.
	Products *[]*string `field:"required" json:"products" yaml:"products"`
}

