package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_RecipeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#name AwsLifecyclePolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#semantic_version AwsLifecyclePolicy#semantic_version}.
	// Experimental.
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}

