package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNatGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#allocation_id TfNatGateway#allocation_id}.
	// Experimental.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#availability_mode TfNatGateway#availability_mode}.
	// Experimental.
	AvailabilityMode *string `field:"optional" json:"availabilityMode" yaml:"availabilityMode"`
	// availability_zone_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#availability_zone_address TfNatGateway#availability_zone_address}
	// Experimental.
	AvailabilityZoneAddress interface{} `field:"optional" json:"availabilityZoneAddress" yaml:"availabilityZoneAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#connectivity_type TfNatGateway#connectivity_type}.
	// Experimental.
	ConnectivityType *string `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#id TfNatGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#private_ip TfNatGateway#private_ip}.
	// Experimental.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#region TfNatGateway#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#secondary_allocation_ids TfNatGateway#secondary_allocation_ids}.
	// Experimental.
	SecondaryAllocationIds *[]*string `field:"optional" json:"secondaryAllocationIds" yaml:"secondaryAllocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#secondary_private_ip_address_count TfNatGateway#secondary_private_ip_address_count}.
	// Experimental.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#secondary_private_ip_addresses TfNatGateway#secondary_private_ip_addresses}.
	// Experimental.
	SecondaryPrivateIpAddresses *[]*string `field:"optional" json:"secondaryPrivateIpAddresses" yaml:"secondaryPrivateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#subnet_id TfNatGateway#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#tags TfNatGateway#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#tags_all TfNatGateway#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#timeouts TfNatGateway#timeouts}
	// Experimental.
	Timeouts *TfNatGateway_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#vpc_id TfNatGateway#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

