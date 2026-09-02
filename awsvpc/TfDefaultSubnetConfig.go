package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDefaultSubnetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#availability_zone TfDefaultSubnet#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#assign_ipv6_address_on_creation TfDefaultSubnet#assign_ipv6_address_on_creation}.
	// Experimental.
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#customer_owned_ipv4_pool TfDefaultSubnet#customer_owned_ipv4_pool}.
	// Experimental.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#enable_dns64 TfDefaultSubnet#enable_dns64}.
	// Experimental.
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#enable_resource_name_dns_aaaa_record_on_launch TfDefaultSubnet#enable_resource_name_dns_aaaa_record_on_launch}.
	// Experimental.
	EnableResourceNameDnsAaaaRecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecordOnLaunch" yaml:"enableResourceNameDnsAaaaRecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#enable_resource_name_dns_a_record_on_launch TfDefaultSubnet#enable_resource_name_dns_a_record_on_launch}.
	// Experimental.
	EnableResourceNameDnsARecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsARecordOnLaunch" yaml:"enableResourceNameDnsARecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#force_destroy TfDefaultSubnet#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#id TfDefaultSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#ipv6_cidr_block TfDefaultSubnet#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#ipv6_native TfDefaultSubnet#ipv6_native}.
	// Experimental.
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#map_customer_owned_ip_on_launch TfDefaultSubnet#map_customer_owned_ip_on_launch}.
	// Experimental.
	MapCustomerOwnedIpOnLaunch interface{} `field:"optional" json:"mapCustomerOwnedIpOnLaunch" yaml:"mapCustomerOwnedIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#map_public_ip_on_launch TfDefaultSubnet#map_public_ip_on_launch}.
	// Experimental.
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#private_dns_hostname_type_on_launch TfDefaultSubnet#private_dns_hostname_type_on_launch}.
	// Experimental.
	PrivateDnsHostnameTypeOnLaunch *string `field:"optional" json:"privateDnsHostnameTypeOnLaunch" yaml:"privateDnsHostnameTypeOnLaunch"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#region TfDefaultSubnet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#tags TfDefaultSubnet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#tags_all TfDefaultSubnet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_subnet#timeouts TfDefaultSubnet#timeouts}
	// Experimental.
	Timeouts *TfDefaultSubnet_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

