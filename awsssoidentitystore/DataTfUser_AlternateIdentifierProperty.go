package awsssoidentitystore


// Experimental.
type DataTfUser_AlternateIdentifierProperty struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#external_id DataTfUser#external_id}
	// Experimental.
	ExternalId *DataTfUser_ExternalIdProperty `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_user#unique_attribute DataTfUser#unique_attribute}
	// Experimental.
	UniqueAttribute *DataTfUser_UniqueAttributeProperty `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

