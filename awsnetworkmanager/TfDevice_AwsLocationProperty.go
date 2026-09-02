package awsnetworkmanager


// Experimental.
type TfDevice_AwsLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#subnet_arn TfDevice#subnet_arn}.
	// Experimental.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_device#zone TfDevice#zone}.
	// Experimental.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

