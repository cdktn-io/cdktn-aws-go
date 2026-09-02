package awskinesisanalyticsv2


// Experimental.
type TfApplication_VpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#security_group_ids TfApplication#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#subnet_ids TfApplication#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

