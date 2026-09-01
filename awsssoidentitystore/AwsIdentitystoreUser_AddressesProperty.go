package awsssoidentitystore


// Experimental.
type AwsIdentitystoreUser_AddressesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#country AwsIdentitystoreUser#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#formatted AwsIdentitystoreUser#formatted}.
	// Experimental.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#locality AwsIdentitystoreUser#locality}.
	// Experimental.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#postal_code AwsIdentitystoreUser#postal_code}.
	// Experimental.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#primary AwsIdentitystoreUser#primary}.
	// Experimental.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#region AwsIdentitystoreUser#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#street_address AwsIdentitystoreUser#street_address}.
	// Experimental.
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#type AwsIdentitystoreUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

