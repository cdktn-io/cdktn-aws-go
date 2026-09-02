package awsfinspace


// Experimental.
type TfKxCluster_CapacityConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#node_count TfKxCluster#node_count}.
	// Experimental.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#node_type TfKxCluster#node_type}.
	// Experimental.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
}

