package awsdirectconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayAssociationProposalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#associated_gateway_id TfGatewayAssociationProposal#associated_gateway_id}.
	// Experimental.
	AssociatedGatewayId *string `field:"required" json:"associatedGatewayId" yaml:"associatedGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#dx_gateway_id TfGatewayAssociationProposal#dx_gateway_id}.
	// Experimental.
	DxGatewayId *string `field:"required" json:"dxGatewayId" yaml:"dxGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#dx_gateway_owner_account_id TfGatewayAssociationProposal#dx_gateway_owner_account_id}.
	// Experimental.
	DxGatewayOwnerAccountId *string `field:"required" json:"dxGatewayOwnerAccountId" yaml:"dxGatewayOwnerAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#allowed_prefixes TfGatewayAssociationProposal#allowed_prefixes}.
	// Experimental.
	AllowedPrefixes *[]*string `field:"optional" json:"allowedPrefixes" yaml:"allowedPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#id TfGatewayAssociationProposal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_gateway_association_proposal#region TfGatewayAssociationProposal#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

