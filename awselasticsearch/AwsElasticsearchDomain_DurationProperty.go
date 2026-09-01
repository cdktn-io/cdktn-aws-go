package awselasticsearch


// Experimental.
type AwsElasticsearchDomain_DurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#unit AwsElasticsearchDomain#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#value AwsElasticsearchDomain#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

