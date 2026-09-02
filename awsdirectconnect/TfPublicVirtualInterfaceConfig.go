package awsdirectconnect

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPublicVirtualInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#address_family TfPublicVirtualInterface#address_family}.
	// Experimental.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#bgp_asn TfPublicVirtualInterface#bgp_asn}.
	// Experimental.
	BgpAsn *float64 `field:"required" json:"bgpAsn" yaml:"bgpAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#connection_id TfPublicVirtualInterface#connection_id}.
	// Experimental.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#name TfPublicVirtualInterface#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#route_filter_prefixes TfPublicVirtualInterface#route_filter_prefixes}.
	// Experimental.
	RouteFilterPrefixes *[]*string `field:"required" json:"routeFilterPrefixes" yaml:"routeFilterPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#vlan TfPublicVirtualInterface#vlan}.
	// Experimental.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#amazon_address TfPublicVirtualInterface#amazon_address}.
	// Experimental.
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#bgp_auth_key TfPublicVirtualInterface#bgp_auth_key}.
	// Experimental.
	BgpAuthKey *string `field:"optional" json:"bgpAuthKey" yaml:"bgpAuthKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#customer_address TfPublicVirtualInterface#customer_address}.
	// Experimental.
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#id TfPublicVirtualInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#region TfPublicVirtualInterface#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#tags TfPublicVirtualInterface#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#tags_all TfPublicVirtualInterface#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dx_public_virtual_interface#timeouts TfPublicVirtualInterface#timeouts}
	// Experimental.
	Timeouts *TfPublicVirtualInterface_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

