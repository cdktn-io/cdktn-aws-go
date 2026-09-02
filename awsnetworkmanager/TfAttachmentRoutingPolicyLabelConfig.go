package awsnetworkmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAttachmentRoutingPolicyLabelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_attachment_routing_policy_label#attachment_id TfAttachmentRoutingPolicyLabel#attachment_id}.
	// Experimental.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_attachment_routing_policy_label#core_network_id TfAttachmentRoutingPolicyLabel#core_network_id}.
	// Experimental.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_attachment_routing_policy_label#routing_policy_label TfAttachmentRoutingPolicyLabel#routing_policy_label}.
	// Experimental.
	RoutingPolicyLabel *string `field:"required" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
}

