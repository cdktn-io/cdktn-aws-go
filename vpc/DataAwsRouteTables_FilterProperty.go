package vpc


// Experimental.
type DataAwsRouteTables_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route_tables#name DataAwsRouteTables#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route_tables#values DataAwsRouteTables#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

