package awsdatasync


// Experimental.
type TfLocationEfs_Ec2ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_efs#security_group_arns TfLocationEfs#security_group_arns}.
	// Experimental.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_location_efs#subnet_arn TfLocationEfs#subnet_arn}.
	// Experimental.
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

