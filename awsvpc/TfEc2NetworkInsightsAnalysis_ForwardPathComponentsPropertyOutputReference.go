package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AclRule() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAclRulePropertyList
	// Experimental.
	AdditionalDetails() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAdditionalDetailsPropertyList
	// Experimental.
	AttachedTo() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAttachedToPropertyList
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
	Component() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsComponentPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DestinationVpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsDestinationVpcPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InboundHeader() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderPropertyList
	// Experimental.
	InternalValue() *TfEc2NetworkInsightsAnalysis_ForwardPathComponentsProperty
	// Experimental.
	SetInternalValue(val *TfEc2NetworkInsightsAnalysis_ForwardPathComponentsProperty)
	// Experimental.
	OutboundHeader() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsOutboundHeaderPropertyList
	// Experimental.
	RouteTableRoute() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsRouteTableRoutePropertyList
	// Experimental.
	SecurityGroupRule() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSecurityGroupRulePropertyList
	// Experimental.
	SequenceNumber() *float64
	// Experimental.
	SourceVpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSourceVpcPropertyList
	// Experimental.
	Subnet() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSubnetPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGateway() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayPropertyList
	// Experimental.
	TransitGatewayRouteTableRoute() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayRouteTableRoutePropertyList
	// Experimental.
	Vpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsVpcPropertyList
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

// The jsii proxy struct for TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference
type jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) AclRule() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAclRulePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAclRulePropertyList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) AdditionalDetails() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAdditionalDetailsPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAdditionalDetailsPropertyList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) AttachedTo() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAttachedToPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsAttachedToPropertyList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) Component() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsComponentPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsComponentPropertyList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) DestinationVpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsDestinationVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsDestinationVpcPropertyList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) InboundHeader() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderPropertyList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) InternalValue() *TfEc2NetworkInsightsAnalysis_ForwardPathComponentsProperty {
	var returns *TfEc2NetworkInsightsAnalysis_ForwardPathComponentsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) OutboundHeader() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsOutboundHeaderPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsOutboundHeaderPropertyList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) RouteTableRoute() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsRouteTableRoutePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) SecurityGroupRule() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSecurityGroupRulePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSecurityGroupRulePropertyList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) SourceVpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSourceVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSourceVpcPropertyList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) Subnet() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSubnetPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsSubnetPropertyList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) TransitGateway() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayPropertyList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) TransitGatewayRouteTableRoute() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayRouteTableRoutePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsTransitGatewayRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) Vpc() TfEc2NetworkInsightsAnalysis_ForwardPathComponentsVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ForwardPathComponentsVpcPropertyList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAnalysis.ForwardPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference_Override(t TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAnalysis.ForwardPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference)SetInternalValue(val *TfEc2NetworkInsightsAnalysis_ForwardPathComponentsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

