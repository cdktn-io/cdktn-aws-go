package vpc


// Experimental.
type DataAwsRouteTable_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route_table#name DataAwsRouteTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route_table#values DataAwsRouteTable#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

