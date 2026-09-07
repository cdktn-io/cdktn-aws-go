package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#type AwsLifecyclePolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// include_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#include_resources AwsLifecyclePolicy#include_resources}
	// Experimental.
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
}

