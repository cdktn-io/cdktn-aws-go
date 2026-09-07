package keyspaces


// Experimental.
type AwsTable_ClusteringKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#name AwsTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#order_by AwsTable#order_by}.
	// Experimental.
	OrderBy *string `field:"required" json:"orderBy" yaml:"orderBy"`
}

