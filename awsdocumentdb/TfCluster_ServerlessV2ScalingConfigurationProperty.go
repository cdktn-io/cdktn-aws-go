package awsdocumentdb


// Experimental.
type TfCluster_ServerlessV2ScalingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#max_capacity TfCluster#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#min_capacity TfCluster#min_capacity}.
	// Experimental.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

