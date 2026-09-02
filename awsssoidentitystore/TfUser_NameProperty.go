package awsssoidentitystore


// Experimental.
type TfUser_NameProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#family_name TfUser#family_name}.
	// Experimental.
	FamilyName *string `field:"required" json:"familyName" yaml:"familyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#given_name TfUser#given_name}.
	// Experimental.
	GivenName *string `field:"required" json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#formatted TfUser#formatted}.
	// Experimental.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#honorific_prefix TfUser#honorific_prefix}.
	// Experimental.
	HonorificPrefix *string `field:"optional" json:"honorificPrefix" yaml:"honorificPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#honorific_suffix TfUser#honorific_suffix}.
	// Experimental.
	HonorificSuffix *string `field:"optional" json:"honorificSuffix" yaml:"honorificSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/identitystore_user#middle_name TfUser#middle_name}.
	// Experimental.
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
}

