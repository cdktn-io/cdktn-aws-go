package cloudsearch


// Experimental.
type AwsDomainServiceAccessPolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain_service_access_policy#delete AwsDomainServiceAccessPolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain_service_access_policy#update AwsDomainServiceAccessPolicy#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

