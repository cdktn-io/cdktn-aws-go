package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#type TfLifecyclePolicy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#value TfLifecyclePolicy#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#retain_at_least TfLifecyclePolicy#retain_at_least}.
	// Experimental.
	RetainAtLeast *float64 `field:"optional" json:"retainAtLeast" yaml:"retainAtLeast"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#unit TfLifecyclePolicy#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

