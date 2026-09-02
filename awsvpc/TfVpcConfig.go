package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVpcConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#assign_generated_ipv6_cidr_block TfVpc#assign_generated_ipv6_cidr_block}.
	// Experimental.
	AssignGeneratedIpv6CidrBlock interface{} `field:"optional" json:"assignGeneratedIpv6CidrBlock" yaml:"assignGeneratedIpv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#cidr_block TfVpc#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#enable_dns_hostnames TfVpc#enable_dns_hostnames}.
	// Experimental.
	EnableDnsHostnames interface{} `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#enable_dns_support TfVpc#enable_dns_support}.
	// Experimental.
	EnableDnsSupport interface{} `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#enable_network_address_usage_metrics TfVpc#enable_network_address_usage_metrics}.
	// Experimental.
	EnableNetworkAddressUsageMetrics interface{} `field:"optional" json:"enableNetworkAddressUsageMetrics" yaml:"enableNetworkAddressUsageMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#id TfVpc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#instance_tenancy TfVpc#instance_tenancy}.
	// Experimental.
	InstanceTenancy *string `field:"optional" json:"instanceTenancy" yaml:"instanceTenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv4_ipam_pool_id TfVpc#ipv4_ipam_pool_id}.
	// Experimental.
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv4_netmask_length TfVpc#ipv4_netmask_length}.
	// Experimental.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv6_cidr_block TfVpc#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv6_cidr_block_network_border_group TfVpc#ipv6_cidr_block_network_border_group}.
	// Experimental.
	Ipv6CidrBlockNetworkBorderGroup *string `field:"optional" json:"ipv6CidrBlockNetworkBorderGroup" yaml:"ipv6CidrBlockNetworkBorderGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv6_ipam_pool_id TfVpc#ipv6_ipam_pool_id}.
	// Experimental.
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#ipv6_netmask_length TfVpc#ipv6_netmask_length}.
	// Experimental.
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#region TfVpc#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#tags TfVpc#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc#tags_all TfVpc#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

