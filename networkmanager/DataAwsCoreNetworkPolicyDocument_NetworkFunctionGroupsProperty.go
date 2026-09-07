package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_NetworkFunctionGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#name DataAwsCoreNetworkPolicyDocument#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#require_attachment_acceptance DataAwsCoreNetworkPolicyDocument#require_attachment_acceptance}.
	// Experimental.
	RequireAttachmentAcceptance interface{} `field:"required" json:"requireAttachmentAcceptance" yaml:"requireAttachmentAcceptance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#description DataAwsCoreNetworkPolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

