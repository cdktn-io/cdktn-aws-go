package awsec2


// Experimental.
type AwsEc2Fleet_MemoryMibProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#min AwsEc2Fleet#min}.
	// Experimental.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#max AwsEc2Fleet#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

