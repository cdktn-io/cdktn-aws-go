package opensearch


// Experimental.
type AwsDomain_DurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#unit AwsDomain#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#value AwsDomain#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

