package awsconnectcustomerprofiles

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#domain_name TfProfile#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#account_number TfProfile#account_number}.
	// Experimental.
	AccountNumber *string `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#additional_information TfProfile#additional_information}.
	// Experimental.
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#address TfProfile#address}
	// Experimental.
	Address *TfProfile_AddressProperty `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#attributes TfProfile#attributes}.
	// Experimental.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// billing_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#billing_address TfProfile#billing_address}
	// Experimental.
	BillingAddress *TfProfile_BillingAddressProperty `field:"optional" json:"billingAddress" yaml:"billingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#birth_date TfProfile#birth_date}.
	// Experimental.
	BirthDate *string `field:"optional" json:"birthDate" yaml:"birthDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#business_email_address TfProfile#business_email_address}.
	// Experimental.
	BusinessEmailAddress *string `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#business_name TfProfile#business_name}.
	// Experimental.
	BusinessName *string `field:"optional" json:"businessName" yaml:"businessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#business_phone_number TfProfile#business_phone_number}.
	// Experimental.
	BusinessPhoneNumber *string `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#email_address TfProfile#email_address}.
	// Experimental.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#first_name TfProfile#first_name}.
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#gender_string TfProfile#gender_string}.
	// Experimental.
	GenderString *string `field:"optional" json:"genderString" yaml:"genderString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#home_phone_number TfProfile#home_phone_number}.
	// Experimental.
	HomePhoneNumber *string `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#id TfProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#last_name TfProfile#last_name}.
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// mailing_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#mailing_address TfProfile#mailing_address}
	// Experimental.
	MailingAddress *TfProfile_MailingAddressProperty `field:"optional" json:"mailingAddress" yaml:"mailingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#middle_name TfProfile#middle_name}.
	// Experimental.
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#mobile_phone_number TfProfile#mobile_phone_number}.
	// Experimental.
	MobilePhoneNumber *string `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#party_type_string TfProfile#party_type_string}.
	// Experimental.
	PartyTypeString *string `field:"optional" json:"partyTypeString" yaml:"partyTypeString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#personal_email_address TfProfile#personal_email_address}.
	// Experimental.
	PersonalEmailAddress *string `field:"optional" json:"personalEmailAddress" yaml:"personalEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#phone_number TfProfile#phone_number}.
	// Experimental.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#region TfProfile#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// shipping_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_profile#shipping_address TfProfile#shipping_address}
	// Experimental.
	ShippingAddress *TfProfile_ShippingAddressProperty `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
}

