package awsopensearch


// Experimental.
type TfDomainPolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_policy#delete TfDomainPolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain_policy#update TfDomainPolicy#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

