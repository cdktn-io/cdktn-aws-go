package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AclRule() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList
	// Experimental.
	AdditionalDetails() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList
	// Experimental.
	AttachedTo() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList
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
	Component() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DestinationVpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InboundHeader() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList
	// Experimental.
	InternalValue() *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty
	// Experimental.
	SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty)
	// Experimental.
	OutboundHeader() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList
	// Experimental.
	RouteTableRoute() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList
	// Experimental.
	SecurityGroupRule() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList
	// Experimental.
	SequenceNumber() *float64
	// Experimental.
	SourceVpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList
	// Experimental.
	Subnet() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGateway() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList
	// Experimental.
	TransitGatewayRouteTableRoute() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList
	// Experimental.
	Vpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList
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

// The jsii proxy struct for AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference
type jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AclRule() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AdditionalDetails() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AttachedTo() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Component() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) DestinationVpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InboundHeader() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InternalValue() *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty {
	var returns *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) OutboundHeader() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) RouteTableRoute() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SecurityGroupRule() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SourceVpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Subnet() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TransitGateway() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TransitGatewayRouteTableRoute() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Vpc() AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList {
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference_Override(a AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetInternalValue(val *AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

