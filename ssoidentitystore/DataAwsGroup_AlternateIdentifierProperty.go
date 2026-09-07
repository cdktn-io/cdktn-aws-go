package ssoidentitystore


// Experimental.
type DataAwsGroup_AlternateIdentifierProperty struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#external_id DataAwsGroup#external_id}
	// Experimental.
	ExternalId *DataAwsGroup_ExternalIdProperty `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#unique_attribute DataAwsGroup#unique_attribute}
	// Experimental.
	UniqueAttribute *DataAwsGroup_UniqueAttributeProperty `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

