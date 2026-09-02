package awsnetworkmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPrefixListAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#core_network_id TfPrefixListAssociation#core_network_id}.
	// Experimental.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#prefix_list_alias TfPrefixListAssociation#prefix_list_alias}.
	// Experimental.
	PrefixListAlias *string `field:"required" json:"prefixListAlias" yaml:"prefixListAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#prefix_list_arn TfPrefixListAssociation#prefix_list_arn}.
	// Experimental.
	PrefixListArn *string `field:"required" json:"prefixListArn" yaml:"prefixListArn"`
}

