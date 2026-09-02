package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_RecipeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#name TfLifecyclePolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#semantic_version TfLifecyclePolicy#semantic_version}.
	// Experimental.
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}

