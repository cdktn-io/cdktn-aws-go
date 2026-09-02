package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNetworkInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#subnet_id TfNetworkInterface#subnet_id}.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#attachment TfNetworkInterface#attachment}
	// Experimental.
	Attachment interface{} `field:"optional" json:"attachment" yaml:"attachment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#description TfNetworkInterface#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#enable_primary_ipv6 TfNetworkInterface#enable_primary_ipv6}.
	// Experimental.
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// ena_srd_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ena_srd_specification TfNetworkInterface#ena_srd_specification}
	// Experimental.
	EnaSrdSpecification *TfNetworkInterface_EnaSrdSpecificationProperty `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#id TfNetworkInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#interface_type TfNetworkInterface#interface_type}.
	// Experimental.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv4_prefix_count TfNetworkInterface#ipv4_prefix_count}.
	// Experimental.
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv4_prefixes TfNetworkInterface#ipv4_prefixes}.
	// Experimental.
	Ipv4Prefixes *[]*string `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_address_count TfNetworkInterface#ipv6_address_count}.
	// Experimental.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_addresses TfNetworkInterface#ipv6_addresses}.
	// Experimental.
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_address_list TfNetworkInterface#ipv6_address_list}.
	// Experimental.
	Ipv6AddressList *[]*string `field:"optional" json:"ipv6AddressList" yaml:"ipv6AddressList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_address_list_enabled TfNetworkInterface#ipv6_address_list_enabled}.
	// Experimental.
	Ipv6AddressListEnabled interface{} `field:"optional" json:"ipv6AddressListEnabled" yaml:"ipv6AddressListEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_prefix_count TfNetworkInterface#ipv6_prefix_count}.
	// Experimental.
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ipv6_prefixes TfNetworkInterface#ipv6_prefixes}.
	// Experimental.
	Ipv6Prefixes *[]*string `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#private_ip TfNetworkInterface#private_ip}.
	// Experimental.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#private_ip_list TfNetworkInterface#private_ip_list}.
	// Experimental.
	PrivateIpList *[]*string `field:"optional" json:"privateIpList" yaml:"privateIpList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#private_ip_list_enabled TfNetworkInterface#private_ip_list_enabled}.
	// Experimental.
	PrivateIpListEnabled interface{} `field:"optional" json:"privateIpListEnabled" yaml:"privateIpListEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#private_ips TfNetworkInterface#private_ips}.
	// Experimental.
	PrivateIps *[]*string `field:"optional" json:"privateIps" yaml:"privateIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#private_ips_count TfNetworkInterface#private_ips_count}.
	// Experimental.
	PrivateIpsCount *float64 `field:"optional" json:"privateIpsCount" yaml:"privateIpsCount"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#region TfNetworkInterface#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#security_groups TfNetworkInterface#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#source_dest_check TfNetworkInterface#source_dest_check}.
	// Experimental.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#tags TfNetworkInterface#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#tags_all TfNetworkInterface#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

