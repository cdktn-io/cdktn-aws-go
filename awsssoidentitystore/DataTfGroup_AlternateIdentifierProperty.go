package awsssoidentitystore


// Experimental.
type DataTfGroup_AlternateIdentifierProperty struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#external_id DataTfGroup#external_id}
	// Experimental.
	ExternalId *DataTfGroup_ExternalIdProperty `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#unique_attribute DataTfGroup#unique_attribute}
	// Experimental.
	UniqueAttribute *DataTfGroup_UniqueAttributeProperty `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

