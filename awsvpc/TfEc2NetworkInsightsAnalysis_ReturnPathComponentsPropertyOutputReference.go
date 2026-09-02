package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AclRule() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList
	// Experimental.
	AdditionalDetails() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList
	// Experimental.
	AttachedTo() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList
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
	Component() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DestinationVpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InboundHeader() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList
	// Experimental.
	InternalValue() *TfEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty
	// Experimental.
	SetInternalValue(val *TfEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty)
	// Experimental.
	OutboundHeader() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList
	// Experimental.
	RouteTableRoute() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList
	// Experimental.
	SecurityGroupRule() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList
	// Experimental.
	SequenceNumber() *float64
	// Experimental.
	SourceVpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList
	// Experimental.
	Subnet() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitGateway() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList
	// Experimental.
	TransitGatewayRouteTableRoute() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList
	// Experimental.
	Vpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList
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

// The jsii proxy struct for TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference
type jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AclRule() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAclRulePropertyList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AdditionalDetails() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAdditionalDetailsPropertyList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) AttachedTo() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Component() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsComponentPropertyList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) DestinationVpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsDestinationVpcPropertyList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InboundHeader() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsInboundHeaderPropertyList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InternalValue() *TfEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty {
	var returns *TfEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) OutboundHeader() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsOutboundHeaderPropertyList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) RouteTableRoute() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SecurityGroupRule() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSecurityGroupRulePropertyList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) SourceVpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSourceVpcPropertyList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Subnet() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsSubnetPropertyList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TransitGateway() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayPropertyList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) TransitGatewayRouteTableRoute() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsTransitGatewayRouteTableRoutePropertyList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Vpc() TfEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList {
	var returns TfEc2NetworkInsightsAnalysis_ReturnPathComponentsVpcPropertyList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAnalysis.ReturnPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference_Override(t TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAnalysis.ReturnPathComponentsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetInternalValue(val *TfEc2NetworkInsightsAnalysis_ReturnPathComponentsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ReturnPathComponentsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

