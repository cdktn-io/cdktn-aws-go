package osis


// Experimental.
type AwsPipeline_VpcOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#subnet_ids AwsPipeline#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#security_group_ids AwsPipeline#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#vpc_endpoint_management AwsPipeline#vpc_endpoint_management}.
	// Experimental.
	VpcEndpointManagement *string `field:"optional" json:"vpcEndpointManagement" yaml:"vpcEndpointManagement"`
}

