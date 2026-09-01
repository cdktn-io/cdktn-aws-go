package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acl() AwsEc2NetworkInsightsAnalysis_AclPropertyList
	// Experimental.
	AclRule() AwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList
	// Experimental.
	Address() *string
	// Experimental.
	Addresses() *[]*string
	// Experimental.
	AttachedTo() AwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList
	// Experimental.
	AvailabilityZones() *[]*string
	// Experimental.
	Cidrs() *[]*string
	// Experimental.
	ClassicLoadBalancerListener() AwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList
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
	Component() AwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomerGateway() AwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList
	// Experimental.
	Destination() AwsEc2NetworkInsightsAnalysis_DestinationPropertyList
	// Experimental.
	DestinationVpc() AwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList
	// Experimental.
	Direction() *string
	// Experimental.
	ElasticLoadBalancerListener() AwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList
	// Experimental.
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IngressRouteTable() AwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList
	// Experimental.
	InternalValue() *AwsEc2NetworkInsightsAnalysis_ExplanationsProperty
	// Experimental.
	SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ExplanationsProperty)
	// Experimental.
	InternetGateway() AwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList
	// Experimental.
	LoadBalancerArn() *string
	// Experimental.
	LoadBalancerListenerPort() *float64
	// Experimental.
	LoadBalancerTargetGroup() AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList
	// Experimental.
	LoadBalancerTargetGroups() AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList
	// Experimental.
	LoadBalancerTargetPort() *float64
	// Experimental.
	MissingComponent() *string
	// Experimental.
	NatGateway() AwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList
	// Experimental.
	NetworkInterface() AwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList
	// Experimental.
	PacketField() *string
	// Experimental.
	Port() *float64
	// Experimental.
	PortRanges() AwsEc2NetworkInsightsAnalysis_PortRangesPropertyList
	// Experimental.
	PrefixList() AwsEc2NetworkInsightsAnalysis_PrefixListPropertyList
	// Experimental.
	Protocols() *[]*string
	// Experimental.
	RouteTable() AwsEc2NetworkInsightsAnalysis_RouteTablePropertyList
	// Experimental.
	RouteTableRoute() AwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList
	// Experimental.
	SecurityGroup() AwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList
	// Experimental.
	SecurityGroupRule() AwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList
	// Experimental.
	SecurityGroups() AwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList
	// Experimental.
	SourceVpc() AwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList
	// Experimental.
	State() *string
	// Experimental.
	Subnet() AwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList
	// Experimental.
	SubnetRouteTable() AwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGateway() AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList
	// Experimental.
	TransitGatewayAttachment() AwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList
	// Experimental.
	TransitGatewayRouteTable() AwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList
	// Experimental.
	TransitGatewayRouteTableRoute() AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList
	// Experimental.
	Vpc() AwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList
	// Experimental.
	VpcEndpoint() AwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList
	// Experimental.
	VpcPeeringConnection() AwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList
	// Experimental.
	VpnConnection() AwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList
	// Experimental.
	VpnGateway() AwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList
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

// The jsii proxy struct for AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference
type jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Acl() AwsEc2NetworkInsightsAnalysis_AclPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_AclPropertyList
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AclRule() AwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsAclRulePropertyList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AttachedTo() AwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsAttachedToPropertyList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ClassicLoadBalancerListener() AwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ClassicLoadBalancerListenerPropertyList
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Component() AwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsComponentPropertyList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) CustomerGateway() AwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_CustomerGatewayPropertyList
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Destination() AwsEc2NetworkInsightsAnalysis_DestinationPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_DestinationPropertyList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) DestinationVpc() AwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsDestinationVpcPropertyList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ElasticLoadBalancerListener() AwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ElasticLoadBalancerListenerPropertyList
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) IngressRouteTable() AwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_IngressRouteTablePropertyList
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InternalValue() *AwsEc2NetworkInsightsAnalysis_ExplanationsProperty {
	var returns *AwsEc2NetworkInsightsAnalysis_ExplanationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InternetGateway() AwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_InternetGatewayPropertyList
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetGroup() AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupPropertyList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetGroups() AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_LoadBalancerTargetGroupsPropertyList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) NatGateway() AwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_NatGatewayPropertyList
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) NetworkInterface() AwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_NetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PortRanges() AwsEc2NetworkInsightsAnalysis_PortRangesPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_PortRangesPropertyList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) PrefixList() AwsEc2NetworkInsightsAnalysis_PrefixListPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_PrefixListPropertyList
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) RouteTable() AwsEc2NetworkInsightsAnalysis_RouteTablePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_RouteTablePropertyList
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) RouteTableRoute() AwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroup() AwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_SecurityGroupPropertyList
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroupRule() AwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsSecurityGroupRulePropertyList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SecurityGroups() AwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_SecurityGroupsPropertyList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SourceVpc() AwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsSourceVpcPropertyList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Subnet() AwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsSubnetPropertyList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) SubnetRouteTable() AwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_SubnetRouteTablePropertyList
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGateway() AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayPropertyList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayAttachment() AwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_TransitGatewayAttachmentPropertyList
	_jsii_.Get(
		j,
		"transitGatewayAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayRouteTable() AwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_TransitGatewayRouteTablePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) TransitGatewayRouteTableRoute() AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsTransitGatewayRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Vpc() AwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ExplanationsVpcPropertyList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpcEndpoint() AwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_VpcEndpointPropertyList
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpcPeeringConnection() AwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_VpcPeeringConnectionPropertyList
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpnConnection() AwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_VpnConnectionPropertyList
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) VpnGateway() AwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_VpnGatewayPropertyList
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ExplanationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference_Override(a AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ExplanationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ExplanationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ExplanationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

