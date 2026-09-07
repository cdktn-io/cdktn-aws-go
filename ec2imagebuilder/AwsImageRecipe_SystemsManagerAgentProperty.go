package ec2imagebuilder


// Experimental.
type AwsImageRecipe_SystemsManagerAgentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#uninstall_after_build AwsImageRecipe#uninstall_after_build}.
	// Experimental.
	UninstallAfterBuild interface{} `field:"required" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

