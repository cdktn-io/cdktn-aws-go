package awselasticsearch


// Experimental.
type TfDomain_DurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#unit TfDomain#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#value TfDomain#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

