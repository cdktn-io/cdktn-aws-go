package awsec2imagebuilder


// Experimental.
type TfLifecyclePolicy_LastLaunchedProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#unit TfLifecyclePolicy#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#value TfLifecyclePolicy#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

