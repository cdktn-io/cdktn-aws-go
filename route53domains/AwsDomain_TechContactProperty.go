package route53domains


// Experimental.
type AwsDomain_TechContactProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#address_line_1 AwsDomain#address_line_1}.
	// Experimental.
	AddressLine1 *string `field:"optional" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#address_line_2 AwsDomain#address_line_2}.
	// Experimental.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#city AwsDomain#city}.
	// Experimental.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#contact_type AwsDomain#contact_type}.
	// Experimental.
	ContactType *string `field:"optional" json:"contactType" yaml:"contactType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#country_code AwsDomain#country_code}.
	// Experimental.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#email AwsDomain#email}.
	// Experimental.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// extra_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#extra_param AwsDomain#extra_param}
	// Experimental.
	ExtraParam interface{} `field:"optional" json:"extraParam" yaml:"extraParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#fax AwsDomain#fax}.
	// Experimental.
	Fax *string `field:"optional" json:"fax" yaml:"fax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#first_name AwsDomain#first_name}.
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#last_name AwsDomain#last_name}.
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#organization_name AwsDomain#organization_name}.
	// Experimental.
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#phone_number AwsDomain#phone_number}.
	// Experimental.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#state AwsDomain#state}.
	// Experimental.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#zip_code AwsDomain#zip_code}.
	// Experimental.
	ZipCode *string `field:"optional" json:"zipCode" yaml:"zipCode"`
}

