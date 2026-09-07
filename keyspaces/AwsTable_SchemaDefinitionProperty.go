package keyspaces


// Experimental.
type AwsTable_SchemaDefinitionProperty struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#column AwsTable#column}
	// Experimental.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#partition_key AwsTable#partition_key}
	// Experimental.
	PartitionKey interface{} `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// clustering_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#clustering_key AwsTable#clustering_key}
	// Experimental.
	ClusteringKey interface{} `field:"optional" json:"clusteringKey" yaml:"clusteringKey"`
	// static_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#static_column AwsTable#static_column}
	// Experimental.
	StaticColumn interface{} `field:"optional" json:"staticColumn" yaml:"staticColumn"`
}

