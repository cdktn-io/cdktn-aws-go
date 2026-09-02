package awsoutpostsec2


// Experimental.
type DataTfCoipPool_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pool#name DataTfCoipPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pool#values DataTfCoipPool#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

