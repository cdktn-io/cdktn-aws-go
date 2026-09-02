package awsopensearch


// Experimental.
type TfDomain_DurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#unit TfDomain#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#value TfDomain#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

