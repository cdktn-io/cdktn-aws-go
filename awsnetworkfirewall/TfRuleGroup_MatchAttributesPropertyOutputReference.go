package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleGroup_MatchAttributesPropertyOutputReference interface {
	cdktn.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Destination() TfRuleGroup_DestinationPropertyList
	// Experimental.
	DestinationInput() interface{}
	// Experimental.
	DestinationPort() TfRuleGroup_DestinationPortPropertyList
	// Experimental.
	DestinationPortInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfRuleGroup_MatchAttributesProperty
	// Experimental.
	SetInternalValue(val *TfRuleGroup_MatchAttributesProperty)
	// Experimental.
	Protocols() *[]*float64
	// Experimental.
	SetProtocols(val *[]*float64)
	// Experimental.
	ProtocolsInput() *[]*float64
	// Experimental.
	Source() TfRuleGroup_SourcePropertyList
	// Experimental.
	SourceInput() interface{}
	// Experimental.
	SourcePort() TfRuleGroup_SourcePortPropertyList
	// Experimental.
	SourcePortInput() interface{}
	// Experimental.
	TcpFlag() TfRuleGroup_TcpFlagPropertyList
	// Experimental.
	TcpFlagInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutDestination(value interface{})
	// Experimental.
	PutDestinationPort(value interface{})
	// Experimental.
	PutSource(value interface{})
	// Experimental.
	PutSourcePort(value interface{})
	// Experimental.
	PutTcpFlag(value interface{})
	// Experimental.
	ResetDestination()
	// Experimental.
	ResetDestinationPort()
	// Experimental.
	ResetProtocols()
	// Experimental.
	ResetSource()
	// Experimental.
	ResetSourcePort()
	// Experimental.
	ResetTcpFlag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRuleGroup_MatchAttributesPropertyOutputReference
type jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) Destination() TfRuleGroup_DestinationPropertyList {
	var returns TfRuleGroup_DestinationPropertyList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) DestinationPort() TfRuleGroup_DestinationPortPropertyList {
	var returns TfRuleGroup_DestinationPortPropertyList
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) DestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) InternalValue() *TfRuleGroup_MatchAttributesProperty {
	var returns *TfRuleGroup_MatchAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) Protocols() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ProtocolsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) Source() TfRuleGroup_SourcePropertyList {
	var returns TfRuleGroup_SourcePropertyList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) SourcePort() TfRuleGroup_SourcePortPropertyList {
	var returns TfRuleGroup_SourcePortPropertyList
	_jsii_.Get(
		j,
		"sourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) SourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) TcpFlag() TfRuleGroup_TcpFlagPropertyList {
	var returns TfRuleGroup_TcpFlagPropertyList
	_jsii_.Get(
		j,
		"tcpFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) TcpFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleGroup_MatchAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRuleGroup_MatchAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleGroup_MatchAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.MatchAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleGroup_MatchAttributesPropertyOutputReference_Override(t TfRuleGroup_MatchAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.MatchAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetInternalValue(val *TfRuleGroup_MatchAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetProtocols(val *[]*float64) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) PutDestination(value interface{}) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) PutDestinationPort(value interface{}) {
	if err := t.validatePutDestinationPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationPort",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) PutSource(value interface{}) {
	if err := t.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) PutSourcePort(value interface{}) {
	if err := t.validatePutSourcePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourcePort",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) PutTcpFlag(value interface{}) {
	if err := t.validatePutTcpFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTcpFlag",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		t,
		"resetSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetSourcePort() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ResetTcpFlag() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpFlag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleGroup_MatchAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

