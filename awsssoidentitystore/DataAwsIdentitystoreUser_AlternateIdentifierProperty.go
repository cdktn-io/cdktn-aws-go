package awsssoidentitystore


// Experimental.
type DataAwsIdentitystoreUser_AlternateIdentifierProperty struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#external_id DataAwsIdentitystoreUser#external_id}
	// Experimental.
	ExternalId *DataAwsIdentitystoreUser_ExternalIdProperty `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#unique_attribute DataAwsIdentitystoreUser#unique_attribute}
	// Experimental.
	UniqueAttribute *DataAwsIdentitystoreUser_UniqueAttributeProperty `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

