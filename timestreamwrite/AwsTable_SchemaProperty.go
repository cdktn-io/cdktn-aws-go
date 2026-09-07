package timestreamwrite


// Experimental.
type AwsTable_SchemaProperty struct {
	// composite_partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#composite_partition_key AwsTable#composite_partition_key}
	// Experimental.
	CompositePartitionKey *AwsTable_CompositePartitionKeyProperty `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}

