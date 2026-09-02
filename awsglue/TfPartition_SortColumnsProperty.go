package awsglue


// Experimental.
type TfPartition_SortColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#column TfPartition#column}.
	// Experimental.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#sort_order TfPartition#sort_order}.
	// Experimental.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

