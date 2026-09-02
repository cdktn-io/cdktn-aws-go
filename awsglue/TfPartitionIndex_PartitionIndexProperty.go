package awsglue


// Experimental.
type TfPartitionIndex_PartitionIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#index_name TfPartitionIndex#index_name}.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#keys TfPartitionIndex#keys}.
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

