package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRouteTable_RoutePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CarrierGatewayId() *string
	// Experimental.
	SetCarrierGatewayId(val *string)
	// Experimental.
	CarrierGatewayIdInput() *string
	// Experimental.
	CidrBlock() *string
	// Experimental.
	SetCidrBlock(val *string)
	// Experimental.
	CidrBlockInput() *string
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
	CoreNetworkArn() *string
	// Experimental.
	SetCoreNetworkArn(val *string)
	// Experimental.
	CoreNetworkArnInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DestinationPrefixListId() *string
	// Experimental.
	SetDestinationPrefixListId(val *string)
	// Experimental.
	DestinationPrefixListIdInput() *string
	// Experimental.
	EgressOnlyGatewayId() *string
	// Experimental.
	SetEgressOnlyGatewayId(val *string)
	// Experimental.
	EgressOnlyGatewayIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GatewayId() *string
	// Experimental.
	SetGatewayId(val *string)
	// Experimental.
	GatewayIdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ipv6CidrBlock() *string
	// Experimental.
	SetIpv6CidrBlock(val *string)
	// Experimental.
	Ipv6CidrBlockInput() *string
	// Experimental.
	LocalGatewayId() *string
	// Experimental.
	SetLocalGatewayId(val *string)
	// Experimental.
	LocalGatewayIdInput() *string
	// Experimental.
	NatGatewayId() *string
	// Experimental.
	SetNatGatewayId(val *string)
	// Experimental.
	NatGatewayIdInput() *string
	// Experimental.
	NetworkInterfaceId() *string
	// Experimental.
	SetNetworkInterfaceId(val *string)
	// Experimental.
	NetworkInterfaceIdInput() *string
	// Experimental.
	OdbNetworkArn() *string
	// Experimental.
	SetOdbNetworkArn(val *string)
	// Experimental.
	OdbNetworkArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGatewayId() *string
	// Experimental.
	SetTransitGatewayId(val *string)
	// Experimental.
	TransitGatewayIdInput() *string
	// Experimental.
	VpcEndpointId() *string
	// Experimental.
	SetVpcEndpointId(val *string)
	// Experimental.
	VpcEndpointIdInput() *string
	// Experimental.
	VpcPeeringConnectionId() *string
	// Experimental.
	SetVpcPeeringConnectionId(val *string)
	// Experimental.
	VpcPeeringConnectionIdInput() *string
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
	// Experimental.
	ResetCarrierGatewayId()
	// Experimental.
	ResetCidrBlock()
	// Experimental.
	ResetCoreNetworkArn()
	// Experimental.
	ResetDestinationPrefixListId()
	// Experimental.
	ResetEgressOnlyGatewayId()
	// Experimental.
	ResetGatewayId()
	// Experimental.
	ResetIpv6CidrBlock()
	// Experimental.
	ResetLocalGatewayId()
	// Experimental.
	ResetNatGatewayId()
	// Experimental.
	ResetNetworkInterfaceId()
	// Experimental.
	ResetOdbNetworkArn()
	// Experimental.
	ResetTransitGatewayId()
	// Experimental.
	ResetVpcEndpointId()
	// Experimental.
	ResetVpcPeeringConnectionId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRouteTable_RoutePropertyOutputReference
type jsiiProxy_AwsRouteTable_RoutePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CarrierGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CarrierGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CoreNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) DestinationPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) EgressOnlyGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) EgressOnlyGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) LocalGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) LocalGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) NatGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) NatGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) OdbNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) VpcPeeringConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionIdInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRouteTable_RoutePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsRouteTable_RoutePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRouteTable_RoutePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRouteTable_RoutePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsRouteTable.RoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRouteTable_RoutePropertyOutputReference_Override(a AwsRouteTable_RoutePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsRouteTable.RoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetCarrierGatewayId(val *string) {
	if err := j.validateSetCarrierGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"carrierGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetCoreNetworkArn(val *string) {
	if err := j.validateSetCoreNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkArn",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetDestinationPrefixListId(val *string) {
	if err := j.validateSetDestinationPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixListId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetEgressOnlyGatewayId(val *string) {
	if err := j.validateSetEgressOnlyGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressOnlyGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetLocalGatewayId(val *string) {
	if err := j.validateSetLocalGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetNatGatewayId(val *string) {
	if err := j.validateSetNatGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetOdbNetworkArn(val *string) {
	if err := j.validateSetOdbNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkArn",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetVpcEndpointId(val *string) {
	if err := j.validateSetVpcEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference)SetVpcPeeringConnectionId(val *string) {
	if err := j.validateSetVpcPeeringConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringConnectionId",
		val,
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetCarrierGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetCarrierGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetCoreNetworkArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreNetworkArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetDestinationPrefixListId() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPrefixListId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetEgressOnlyGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetEgressOnlyGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetLocalGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetNatGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetNatGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetOdbNetworkArn() {
	_jsii_.InvokeVoid(
		a,
		"resetOdbNetworkArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ResetVpcPeeringConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcPeeringConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRouteTable_RoutePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

