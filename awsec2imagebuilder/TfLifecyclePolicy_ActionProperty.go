package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_ActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#type TfLifecyclePolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// include_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#include_resources TfLifecyclePolicy#include_resources}
	// Experimental.
	IncludeResources interface{} `field:"optional" json:"includeResources" yaml:"includeResources"`
}

