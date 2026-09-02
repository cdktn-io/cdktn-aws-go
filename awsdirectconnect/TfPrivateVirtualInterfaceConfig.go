package awsdirectconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPrivateVirtualInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#address_family TfPrivateVirtualInterface#address_family}.
	// Experimental.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#connection_id TfPrivateVirtualInterface#connection_id}.
	// Experimental.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#name TfPrivateVirtualInterface#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#vlan TfPrivateVirtualInterface#vlan}.
	// Experimental.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#amazon_address TfPrivateVirtualInterface#amazon_address}.
	// Experimental.
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#bgp_asn TfPrivateVirtualInterface#bgp_asn}.
	// Experimental.
	BgpAsn *float64 `field:"optional" json:"bgpAsn" yaml:"bgpAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#bgp_asn_long TfPrivateVirtualInterface#bgp_asn_long}.
	// Experimental.
	BgpAsnLong *string `field:"optional" json:"bgpAsnLong" yaml:"bgpAsnLong"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#bgp_auth_key TfPrivateVirtualInterface#bgp_auth_key}.
	// Experimental.
	BgpAuthKey *string `field:"optional" json:"bgpAuthKey" yaml:"bgpAuthKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#customer_address TfPrivateVirtualInterface#customer_address}.
	// Experimental.
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#dx_gateway_id TfPrivateVirtualInterface#dx_gateway_id}.
	// Experimental.
	DxGatewayId *string `field:"optional" json:"dxGatewayId" yaml:"dxGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#id TfPrivateVirtualInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#mtu TfPrivateVirtualInterface#mtu}.
	// Experimental.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#region TfPrivateVirtualInterface#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#sitelink_enabled TfPrivateVirtualInterface#sitelink_enabled}.
	// Experimental.
	SitelinkEnabled interface{} `field:"optional" json:"sitelinkEnabled" yaml:"sitelinkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#tags TfPrivateVirtualInterface#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#tags_all TfPrivateVirtualInterface#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#timeouts TfPrivateVirtualInterface#timeouts}
	// Experimental.
	Timeouts *TfPrivateVirtualInterface_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_private_virtual_interface#vpn_gateway_id TfPrivateVirtualInterface#vpn_gateway_id}.
	// Experimental.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

