package glue


// Experimental.
type AwsPartition_SortColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#column AwsPartition#column}.
	// Experimental.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#sort_order AwsPartition#sort_order}.
	// Experimental.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

