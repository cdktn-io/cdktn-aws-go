package networkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRuleGroup_MatchAttributesPropertyOutputReference interface {
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
	Destination() AwsRuleGroup_DestinationPropertyList
	// Experimental.
	DestinationInput() interface{}
	// Experimental.
	DestinationPort() AwsRuleGroup_DestinationPortPropertyList
	// Experimental.
	DestinationPortInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsRuleGroup_MatchAttributesProperty
	// Experimental.
	SetInternalValue(val *AwsRuleGroup_MatchAttributesProperty)
	// Experimental.
	Protocols() *[]*float64
	// Experimental.
	SetProtocols(val *[]*float64)
	// Experimental.
	ProtocolsInput() *[]*float64
	// Experimental.
	Source() AwsRuleGroup_SourcePropertyList
	// Experimental.
	SourceInput() interface{}
	// Experimental.
	SourcePort() AwsRuleGroup_SourcePortPropertyList
	// Experimental.
	SourcePortInput() interface{}
	// Experimental.
	TcpFlag() AwsRuleGroup_TcpFlagPropertyList
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

// The jsii proxy struct for AwsRuleGroup_MatchAttributesPropertyOutputReference
type jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) Destination() AwsRuleGroup_DestinationPropertyList {
	var returns AwsRuleGroup_DestinationPropertyList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) DestinationPort() AwsRuleGroup_DestinationPortPropertyList {
	var returns AwsRuleGroup_DestinationPortPropertyList
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) DestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) InternalValue() *AwsRuleGroup_MatchAttributesProperty {
	var returns *AwsRuleGroup_MatchAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) Protocols() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ProtocolsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) Source() AwsRuleGroup_SourcePropertyList {
	var returns AwsRuleGroup_SourcePropertyList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) SourcePort() AwsRuleGroup_SourcePortPropertyList {
	var returns AwsRuleGroup_SourcePortPropertyList
	_jsii_.Get(
		j,
		"sourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) SourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) TcpFlag() AwsRuleGroup_TcpFlagPropertyList {
	var returns AwsRuleGroup_TcpFlagPropertyList
	_jsii_.Get(
		j,
		"tcpFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) TcpFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRuleGroup_MatchAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRuleGroup_MatchAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRuleGroup_MatchAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsRuleGroup.MatchAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRuleGroup_MatchAttributesPropertyOutputReference_Override(a AwsRuleGroup_MatchAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsRuleGroup.MatchAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetInternalValue(val *AwsRuleGroup_MatchAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetProtocols(val *[]*float64) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) PutDestination(value interface{}) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) PutDestinationPort(value interface{}) {
	if err := a.validatePutDestinationPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationPort",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) PutSource(value interface{}) {
	if err := a.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) PutSourcePort(value interface{}) {
	if err := a.validatePutSourcePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourcePort",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) PutTcpFlag(value interface{}) {
	if err := a.validatePutTcpFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTcpFlag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetSourcePort() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcePort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ResetTcpFlag() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpFlag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_MatchAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

