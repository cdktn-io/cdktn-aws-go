package awsdynamodb


// Experimental.
type TfGlobalSecondaryIndex_ProjectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#projection_type TfGlobalSecondaryIndex#projection_type}.
	// Experimental.
	ProjectionType *string `field:"required" json:"projectionType" yaml:"projectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#non_key_attributes TfGlobalSecondaryIndex#non_key_attributes}.
	// Experimental.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
}

