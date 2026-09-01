package awsec2


// Experimental.
type AwsEc2Fleet_CapacityRebalanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#replacement_strategy AwsEc2Fleet#replacement_strategy}.
	// Experimental.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#termination_delay AwsEc2Fleet#termination_delay}.
	// Experimental.
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
}

