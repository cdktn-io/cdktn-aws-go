package ec2imagebuilder


// Experimental.
type AwsLifecyclePolicy_LastLaunchedProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#unit AwsLifecyclePolicy#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_lifecycle_policy#value AwsLifecyclePolicy#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

