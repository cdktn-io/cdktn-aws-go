package awsosis


// Experimental.
type TfPipeline_VpcOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#subnet_ids TfPipeline#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#security_group_ids TfPipeline#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#vpc_endpoint_management TfPipeline#vpc_endpoint_management}.
	// Experimental.
	VpcEndpointManagement *string `field:"optional" json:"vpcEndpointManagement" yaml:"vpcEndpointManagement"`
}

