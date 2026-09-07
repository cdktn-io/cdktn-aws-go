package ssoidentitystore


// Experimental.
type AwsUser_AddressesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#country AwsUser#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#formatted AwsUser#formatted}.
	// Experimental.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#locality AwsUser#locality}.
	// Experimental.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#postal_code AwsUser#postal_code}.
	// Experimental.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#primary AwsUser#primary}.
	// Experimental.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#region AwsUser#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#street_address AwsUser#street_address}.
	// Experimental.
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#type AwsUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

