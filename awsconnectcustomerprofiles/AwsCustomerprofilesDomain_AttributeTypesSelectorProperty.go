package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_AttributeTypesSelectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#attribute_matching_model AwsCustomerprofilesDomain#attribute_matching_model}.
	// Experimental.
	AttributeMatchingModel *string `field:"required" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#address AwsCustomerprofilesDomain#address}.
	// Experimental.
	Address *[]*string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#email_address AwsCustomerprofilesDomain#email_address}.
	// Experimental.
	EmailAddress *[]*string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#phone_number AwsCustomerprofilesDomain#phone_number}.
	// Experimental.
	PhoneNumber *[]*string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

