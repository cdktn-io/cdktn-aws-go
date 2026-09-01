package awsssm


// Experimental.
type AwsSsmPatchBaseline_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#configuration AwsSsmPatchBaseline#configuration}.
	// Experimental.
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#name AwsSsmPatchBaseline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_patch_baseline#products AwsSsmPatchBaseline#products}.
	// Experimental.
	Products *[]*string `field:"required" json:"products" yaml:"products"`
}

