package awsec2imagebuilder


// Experimental.
type TfImageRecipe_SystemsManagerAgentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#uninstall_after_build TfImageRecipe#uninstall_after_build}.
	// Experimental.
	UninstallAfterBuild interface{} `field:"required" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

