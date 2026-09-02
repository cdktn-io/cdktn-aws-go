package awsmsk


// Experimental.
type TfServerlessCluster_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_serverless_cluster#subnet_ids TfServerlessCluster#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_serverless_cluster#security_group_ids TfServerlessCluster#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

