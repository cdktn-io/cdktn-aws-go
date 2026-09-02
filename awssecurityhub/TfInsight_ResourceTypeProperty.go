package awssecurityhub


// Experimental.
type TfInsight_ResourceTypeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#comparison TfInsight#comparison}.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#value TfInsight#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

