package awsopensearch


// Experimental.
type TfVpcEndpoint_VpcOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_vpc_endpoint#subnet_ids TfVpcEndpoint#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_vpc_endpoint#security_group_ids TfVpcEndpoint#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

