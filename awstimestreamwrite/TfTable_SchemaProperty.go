package awstimestreamwrite


// Experimental.
type TfTable_SchemaProperty struct {
	// composite_partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#composite_partition_key TfTable#composite_partition_key}
	// Experimental.
	CompositePartitionKey *TfTable_CompositePartitionKeyProperty `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}

