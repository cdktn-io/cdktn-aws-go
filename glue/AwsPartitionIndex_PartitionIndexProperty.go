package glue


// Experimental.
type AwsPartitionIndex_PartitionIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#index_name AwsPartitionIndex#index_name}.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#keys AwsPartitionIndex#keys}.
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

