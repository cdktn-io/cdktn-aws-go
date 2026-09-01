package awsvpc


// Experimental.
type DataAwsNetworkAcls_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_acls#name DataAwsNetworkAcls#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_acls#values DataAwsNetworkAcls#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

