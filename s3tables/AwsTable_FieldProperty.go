package s3tables


// Experimental.
type AwsTable_FieldProperty struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#name AwsTable#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The field type. S3 Tables supports all Apache Iceberg primitive types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#type AwsTable#type}
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A Boolean value that specifies whether values are required for each row in this field. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#required AwsTable#required}
	// Experimental.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

