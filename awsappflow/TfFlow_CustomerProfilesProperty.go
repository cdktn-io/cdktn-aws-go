package awsappflow


// Experimental.
type TfFlow_CustomerProfilesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#domain_name TfFlow#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#object_type_name TfFlow#object_type_name}.
	// Experimental.
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
}

