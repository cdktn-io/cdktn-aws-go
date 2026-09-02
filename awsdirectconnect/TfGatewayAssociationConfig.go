package awsdirectconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#dx_gateway_id TfGatewayAssociation#dx_gateway_id}.
	// Experimental.
	DxGatewayId *string `field:"required" json:"dxGatewayId" yaml:"dxGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#allowed_prefixes TfGatewayAssociation#allowed_prefixes}.
	// Experimental.
	AllowedPrefixes *[]*string `field:"optional" json:"allowedPrefixes" yaml:"allowedPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#associated_gateway_id TfGatewayAssociation#associated_gateway_id}.
	// Experimental.
	AssociatedGatewayId *string `field:"optional" json:"associatedGatewayId" yaml:"associatedGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#associated_gateway_owner_account_id TfGatewayAssociation#associated_gateway_owner_account_id}.
	// Experimental.
	AssociatedGatewayOwnerAccountId *string `field:"optional" json:"associatedGatewayOwnerAccountId" yaml:"associatedGatewayOwnerAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#id TfGatewayAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#proposal_id TfGatewayAssociation#proposal_id}.
	// Experimental.
	ProposalId *string `field:"optional" json:"proposalId" yaml:"proposalId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#region TfGatewayAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association#timeouts TfGatewayAssociation#timeouts}
	// Experimental.
	Timeouts *TfGatewayAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

