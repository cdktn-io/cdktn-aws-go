package awsssoidentitystore


// Experimental.
type TfUser_AddressesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#country TfUser#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#formatted TfUser#formatted}.
	// Experimental.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#locality TfUser#locality}.
	// Experimental.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#postal_code TfUser#postal_code}.
	// Experimental.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#primary TfUser#primary}.
	// Experimental.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#region TfUser#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#street_address TfUser#street_address}.
	// Experimental.
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#type TfUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

