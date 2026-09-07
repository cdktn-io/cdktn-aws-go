package finspace


// Experimental.
type AwsKxCluster_DatabaseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#database_name AwsKxCluster#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// cache_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#cache_configurations AwsKxCluster#cache_configurations}
	// Experimental.
	CacheConfigurations interface{} `field:"optional" json:"cacheConfigurations" yaml:"cacheConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#changeset_id AwsKxCluster#changeset_id}.
	// Experimental.
	ChangesetId *string `field:"optional" json:"changesetId" yaml:"changesetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#dataview_name AwsKxCluster#dataview_name}.
	// Experimental.
	DataviewName *string `field:"optional" json:"dataviewName" yaml:"dataviewName"`
}

