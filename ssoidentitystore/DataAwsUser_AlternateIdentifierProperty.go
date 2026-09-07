package ssoidentitystore


// Experimental.
type DataAwsUser_AlternateIdentifierProperty struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#external_id DataAwsUser#external_id}
	// Experimental.
	ExternalId *DataAwsUser_ExternalIdProperty `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#unique_attribute DataAwsUser#unique_attribute}
	// Experimental.
	UniqueAttribute *DataAwsUser_UniqueAttributeProperty `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

