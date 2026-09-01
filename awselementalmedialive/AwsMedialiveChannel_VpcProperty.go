package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_VpcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#public_address_allocation_ids AwsMedialiveChannel#public_address_allocation_ids}.
	// Experimental.
	PublicAddressAllocationIds *[]*string `field:"required" json:"publicAddressAllocationIds" yaml:"publicAddressAllocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#subnet_ids AwsMedialiveChannel#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#security_group_ids AwsMedialiveChannel#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

