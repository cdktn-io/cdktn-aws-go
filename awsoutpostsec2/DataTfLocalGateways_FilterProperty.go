package awsoutpostsec2


// Experimental.
type DataTfLocalGateways_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateways#name DataTfLocalGateways#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateways#values DataTfLocalGateways#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

