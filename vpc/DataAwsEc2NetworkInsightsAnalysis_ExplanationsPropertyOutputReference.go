package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acl() DataAwsEc2NetworkInsightsAnalysis_AclPropertyList
	// Experimental.
	AclRule() DataAwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList
	// Experimental.
	Address() *string
	// Experimental.
	Addresses() *[]*string
	// Experimental.
	AttachedTo() DataAwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList
	// Experimental.
	AvailabilityZones() *[]*string
	// Experimental.
	Cidrs() *[]*string
	// Experimental.
	ClassicLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	Component() DataAwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomerGateway() DataAwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList
	// Experimental.
	Destination() DataAwsEc2NetworkInsightsAnalysis_DestinationPropertyList
	// Experimental.
	DestinationVpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList
	// Experimental.
	Direction() *string
	// Experimental.
	ElasticLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList
	// Experimental.
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IngressRouteTable() DataAwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList
	// Experimental.
	InternalValue() *DataAwsEc2NetworkInsightsAnalysis_ExplanationsProperty
	// Experimental.
	SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysis_ExplanationsProperty)
	// Experimental.
	InternetGateway() DataAwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList
	// Experimental.
	LoadBalancerArn() *string
	// Experimental.
	LoadBalancerListenerPort() *float64
	// Experimental.
	LoadBalancerTargetGroup() DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList
	// Experimental.
	LoadBalancerTargetGroups() DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList
	// Experimental.
	LoadBalancerTargetPort() *float64
	// Experimental.
	MissingComponent() *string
	// Experimental.
	NatGateway() DataAwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList
	// Experimental.
	NetworkInterface() DataAwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList
	// Experimental.
	PacketField() *string
	// Experimental.
	Port() *float64
	// Experimental.
	PortRanges() DataAwsEc2NetworkInsightsAnalysis_PortRangesPropertyList
	// Experimental.
	PrefixList() DataAwsEc2NetworkInsightsAnalysis_PrefixListPropertyList
	// Experimental.
	Protocols() *[]*string
	// Experimental.
	RouteTable() DataAwsEc2NetworkInsightsAnalysis_RouteTablePropertyList
	// Experimental.
	RouteTableRoute() DataAwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList
	// Experimental.
	SecurityGroup() DataAwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList
	// Experimental.
	SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList
	// Experimental.
	SecurityGroups() DataAwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList
	// Experimental.
	SourceVpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList
	// Experimental.
	State() *string
	// Experimental.
	Subnet() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList
	// Experimental.
	SubnetRouteTable() DataAwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGateway() DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList
	// Experimental.
	TransitGatewayAttachment() DataAwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList
	// Experimental.
	TransitGatewayRouteTable() DataAwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList
	// Experimental.
	TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList
	// Experimental.
	Vpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList
	// Experimental.
	VpcEndpoint() DataAwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList
	// Experimental.
	VpcPeeringConnection() DataAwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList
	// Experimental.
	VpnConnection() DataAwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList
	// Experimental.
	VpnGateway() DataAwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference
type jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Acl() DataAwsEc2NetworkInsightsAnalysis_AclPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_AclPropertyList
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AclRule() DataAwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AttachedTo() DataAwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ClassicLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Component() DataAwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) CustomerGateway() DataAwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Destination() DataAwsEc2NetworkInsightsAnalysis_DestinationPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_DestinationPropertyList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) DestinationVpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ElasticLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) IngressRouteTable() DataAwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InternalValue() *DataAwsEc2NetworkInsightsAnalysis_ExplanationsProperty {
	var returns *DataAwsEc2NetworkInsightsAnalysis_ExplanationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InternetGateway() DataAwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetGroup() DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetGroups() DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) NatGateway() DataAwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) NetworkInterface() DataAwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PortRanges() DataAwsEc2NetworkInsightsAnalysis_PortRangesPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_PortRangesPropertyList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PrefixList() DataAwsEc2NetworkInsightsAnalysis_PrefixListPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_PrefixListPropertyList
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) RouteTable() DataAwsEc2NetworkInsightsAnalysis_RouteTablePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_RouteTablePropertyList
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) RouteTableRoute() DataAwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroup() DataAwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroups() DataAwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SourceVpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Subnet() DataAwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SubnetRouteTable() DataAwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGateway() DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayAttachment() DataAwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList
	_jsii_.Get(
		j,
		"transitGatewayAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayRouteTable() DataAwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Vpc() DataAwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpcEndpoint() DataAwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpcPeeringConnection() DataAwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpnConnection() DataAwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpnGateway() DataAwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList {
	var returns DataAwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.DataAwsEc2NetworkInsightsAnalysis.ExplanationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference_Override(d DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.DataAwsEc2NetworkInsightsAnalysis.ExplanationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysis_ExplanationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

