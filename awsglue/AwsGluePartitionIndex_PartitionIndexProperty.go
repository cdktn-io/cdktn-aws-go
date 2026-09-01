package awsglue


// Experimental.
type AwsGluePartitionIndex_PartitionIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#index_name AwsGluePartitionIndex#index_name}.
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition_index#keys AwsGluePartitionIndex#keys}.
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

