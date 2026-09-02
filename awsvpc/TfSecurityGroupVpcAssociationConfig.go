package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSecurityGroupVpcAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_security_group_vpc_association#security_group_id TfSecurityGroupVpcAssociation#security_group_id}.
	// Experimental.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_security_group_vpc_association#vpc_id TfSecurityGroupVpcAssociation#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_security_group_vpc_association#region TfSecurityGroupVpcAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_security_group_vpc_association#timeouts TfSecurityGroupVpcAssociation#timeouts}
	// Experimental.
	Timeouts *TfSecurityGroupVpcAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

