package finspace


// Experimental.
type AwsKxEnvironment_PortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#from AwsKxEnvironment#from}.
	// Experimental.
	From *float64 `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#to AwsKxEnvironment#to}.
	// Experimental.
	To *float64 `field:"required" json:"to" yaml:"to"`
}

