package networkmanager


// Experimental.
type AwsDevice_AwsLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#subnet_arn AwsDevice#subnet_arn}.
	// Experimental.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#zone AwsDevice#zone}.
	// Experimental.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

