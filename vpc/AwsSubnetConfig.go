package vpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSubnetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#vpc_id AwsSubnet#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#assign_ipv6_address_on_creation AwsSubnet#assign_ipv6_address_on_creation}.
	// Experimental.
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#availability_zone AwsSubnet#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#availability_zone_id AwsSubnet#availability_zone_id}.
	// Experimental.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#cidr_block AwsSubnet#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#customer_owned_ipv4_pool AwsSubnet#customer_owned_ipv4_pool}.
	// Experimental.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#enable_dns64 AwsSubnet#enable_dns64}.
	// Experimental.
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#enable_lni_at_device_index AwsSubnet#enable_lni_at_device_index}.
	// Experimental.
	EnableLniAtDeviceIndex *float64 `field:"optional" json:"enableLniAtDeviceIndex" yaml:"enableLniAtDeviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#enable_resource_name_dns_aaaa_record_on_launch AwsSubnet#enable_resource_name_dns_aaaa_record_on_launch}.
	// Experimental.
	EnableResourceNameDnsAaaaRecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecordOnLaunch" yaml:"enableResourceNameDnsAaaaRecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#enable_resource_name_dns_a_record_on_launch AwsSubnet#enable_resource_name_dns_a_record_on_launch}.
	// Experimental.
	EnableResourceNameDnsARecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsARecordOnLaunch" yaml:"enableResourceNameDnsARecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#id AwsSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv4_ipam_pool_id AwsSubnet#ipv4_ipam_pool_id}.
	// Experimental.
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv4_netmask_length AwsSubnet#ipv4_netmask_length}.
	// Experimental.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv6_cidr_block AwsSubnet#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv6_ipam_pool_id AwsSubnet#ipv6_ipam_pool_id}.
	// Experimental.
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv6_native AwsSubnet#ipv6_native}.
	// Experimental.
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#ipv6_netmask_length AwsSubnet#ipv6_netmask_length}.
	// Experimental.
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#map_customer_owned_ip_on_launch AwsSubnet#map_customer_owned_ip_on_launch}.
	// Experimental.
	MapCustomerOwnedIpOnLaunch interface{} `field:"optional" json:"mapCustomerOwnedIpOnLaunch" yaml:"mapCustomerOwnedIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#map_public_ip_on_launch AwsSubnet#map_public_ip_on_launch}.
	// Experimental.
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#outpost_arn AwsSubnet#outpost_arn}.
	// Experimental.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#private_dns_hostname_type_on_launch AwsSubnet#private_dns_hostname_type_on_launch}.
	// Experimental.
	PrivateDnsHostnameTypeOnLaunch *string `field:"optional" json:"privateDnsHostnameTypeOnLaunch" yaml:"privateDnsHostnameTypeOnLaunch"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#region AwsSubnet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#tags AwsSubnet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#tags_all AwsSubnet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#timeouts AwsSubnet#timeouts}
	// Experimental.
	Timeouts *AwsSubnet_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

