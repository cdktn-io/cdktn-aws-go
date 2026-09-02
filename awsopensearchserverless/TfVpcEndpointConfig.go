package awsopensearchserverless

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVpcEndpointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the interface endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#name TfVpcEndpoint#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or more subnet IDs from which you'll access OpenSearch Serverless. Up to 6 subnets can be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#subnet_ids TfVpcEndpoint#subnet_ids}
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// ID of the VPC from which you'll access OpenSearch Serverless.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#vpc_id TfVpcEndpoint#vpc_id}
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#region TfVpcEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// One or more security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	//
	// Up to 5 security groups can be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#security_group_ids TfVpcEndpoint#security_group_ids}
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_vpc_endpoint#timeouts TfVpcEndpoint#timeouts}
	// Experimental.
	Timeouts *TfVpcEndpoint_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

