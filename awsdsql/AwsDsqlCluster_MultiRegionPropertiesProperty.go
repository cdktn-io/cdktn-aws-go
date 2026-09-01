package awsdsql


// Experimental.
type AwsDsqlCluster_MultiRegionPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dsql_cluster#clusters AwsDsqlCluster#clusters}.
	// Experimental.
	Clusters *[]*string `field:"optional" json:"clusters" yaml:"clusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dsql_cluster#witness_region AwsDsqlCluster#witness_region}.
	// Experimental.
	WitnessRegion *string `field:"optional" json:"witnessRegion" yaml:"witnessRegion"`
}

