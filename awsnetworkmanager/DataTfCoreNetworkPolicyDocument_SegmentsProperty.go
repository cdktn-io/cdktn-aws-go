package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_SegmentsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#name DataTfCoreNetworkPolicyDocument#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#allow_filter DataTfCoreNetworkPolicyDocument#allow_filter}.
	// Experimental.
	AllowFilter *[]*string `field:"optional" json:"allowFilter" yaml:"allowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#deny_filter DataTfCoreNetworkPolicyDocument#deny_filter}.
	// Experimental.
	DenyFilter *[]*string `field:"optional" json:"denyFilter" yaml:"denyFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#description DataTfCoreNetworkPolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_locations DataTfCoreNetworkPolicyDocument#edge_locations}.
	// Experimental.
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#isolate_attachments DataTfCoreNetworkPolicyDocument#isolate_attachments}.
	// Experimental.
	IsolateAttachments interface{} `field:"optional" json:"isolateAttachments" yaml:"isolateAttachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#require_attachment_acceptance DataTfCoreNetworkPolicyDocument#require_attachment_acceptance}.
	// Experimental.
	RequireAttachmentAcceptance interface{} `field:"optional" json:"requireAttachmentAcceptance" yaml:"requireAttachmentAcceptance"`
}

