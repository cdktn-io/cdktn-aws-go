package awscloudsearch


// Experimental.
type TfDomainServiceAccessPolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain_service_access_policy#delete TfDomainServiceAccessPolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain_service_access_policy#update TfDomainServiceAccessPolicy#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

