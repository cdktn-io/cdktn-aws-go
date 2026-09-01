package awstimestreamwrite


// Experimental.
type AwsTimestreamwriteTable_SchemaProperty struct {
	// composite_partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#composite_partition_key AwsTimestreamwriteTable#composite_partition_key}
	// Experimental.
	CompositePartitionKey *AwsTimestreamwriteTable_CompositePartitionKeyProperty `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}

