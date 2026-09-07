package accountmanagement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPrimaryContactConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#address_line_1 AwsPrimaryContact#address_line_1}.
	// Experimental.
	AddressLine1 *string `field:"required" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#city AwsPrimaryContact#city}.
	// Experimental.
	City *string `field:"required" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#country_code AwsPrimaryContact#country_code}.
	// Experimental.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#full_name AwsPrimaryContact#full_name}.
	// Experimental.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#phone_number AwsPrimaryContact#phone_number}.
	// Experimental.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#postal_code AwsPrimaryContact#postal_code}.
	// Experimental.
	PostalCode *string `field:"required" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#account_id AwsPrimaryContact#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#address_line_2 AwsPrimaryContact#address_line_2}.
	// Experimental.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#address_line_3 AwsPrimaryContact#address_line_3}.
	// Experimental.
	AddressLine3 *string `field:"optional" json:"addressLine3" yaml:"addressLine3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#company_name AwsPrimaryContact#company_name}.
	// Experimental.
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#district_or_county AwsPrimaryContact#district_or_county}.
	// Experimental.
	DistrictOrCounty *string `field:"optional" json:"districtOrCounty" yaml:"districtOrCounty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#id AwsPrimaryContact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#state_or_region AwsPrimaryContact#state_or_region}.
	// Experimental.
	StateOrRegion *string `field:"optional" json:"stateOrRegion" yaml:"stateOrRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/account_primary_contact#website_url AwsPrimaryContact#website_url}.
	// Experimental.
	WebsiteUrl *string `field:"optional" json:"websiteUrl" yaml:"websiteUrl"`
}

