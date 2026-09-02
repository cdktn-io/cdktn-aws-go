package awsverifiedaccess


// Experimental.
type TfEndpoint_CidrOptionsPortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#from_port TfEndpoint#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint#to_port TfEndpoint#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

