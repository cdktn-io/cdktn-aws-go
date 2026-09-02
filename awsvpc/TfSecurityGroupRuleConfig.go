package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSecurityGroupRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#from_port TfSecurityGroupRule#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#protocol TfSecurityGroupRule#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#security_group_id TfSecurityGroupRule#security_group_id}.
	// Experimental.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#to_port TfSecurityGroupRule#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#type TfSecurityGroupRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#cidr_blocks TfSecurityGroupRule#cidr_blocks}.
	// Experimental.
	CidrBlocks *[]*string `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#description TfSecurityGroupRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#id TfSecurityGroupRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#ipv6_cidr_blocks TfSecurityGroupRule#ipv6_cidr_blocks}.
	// Experimental.
	Ipv6CidrBlocks *[]*string `field:"optional" json:"ipv6CidrBlocks" yaml:"ipv6CidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#prefix_list_ids TfSecurityGroupRule#prefix_list_ids}.
	// Experimental.
	PrefixListIds *[]*string `field:"optional" json:"prefixListIds" yaml:"prefixListIds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#region TfSecurityGroupRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#self TfSecurityGroupRule#self}.
	// Experimental.
	SelfAttribute interface{} `field:"optional" json:"selfAttribute" yaml:"selfAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#source_security_group_id TfSecurityGroupRule#source_security_group_id}.
	// Experimental.
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/security_group_rule#timeouts TfSecurityGroupRule#timeouts}
	// Experimental.
	Timeouts *TfSecurityGroupRule_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

