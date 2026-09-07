package vpc


// Experimental.
type DataAwsNetworkInterfaces_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_interfaces#name DataAwsNetworkInterfaces#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/network_interfaces#values DataAwsNetworkInterfaces#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

