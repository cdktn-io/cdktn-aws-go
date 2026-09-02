package awsnetworkmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectPeerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#connect_attachment_id TfConnectPeer#connect_attachment_id}.
	// Experimental.
	ConnectAttachmentId *string `field:"required" json:"connectAttachmentId" yaml:"connectAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#peer_address TfConnectPeer#peer_address}.
	// Experimental.
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// bgp_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#bgp_options TfConnectPeer#bgp_options}
	// Experimental.
	BgpOptions *TfConnectPeer_BgpOptionsProperty `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#core_network_address TfConnectPeer#core_network_address}.
	// Experimental.
	CoreNetworkAddress *string `field:"optional" json:"coreNetworkAddress" yaml:"coreNetworkAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#id TfConnectPeer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#inside_cidr_blocks TfConnectPeer#inside_cidr_blocks}.
	// Experimental.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#subnet_arn TfConnectPeer#subnet_arn}.
	// Experimental.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#tags TfConnectPeer#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#tags_all TfConnectPeer#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer#timeouts TfConnectPeer#timeouts}
	// Experimental.
	Timeouts *TfConnectPeer_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

