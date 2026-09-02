package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTlsInspectionConfiguration_ScopePropertyOutputReference interface {
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
	Destination() TfTlsInspectionConfiguration_DestinationPropertyList
	// Experimental.
	DestinationInput() interface{}
	// Experimental.
	DestinationPorts() TfTlsInspectionConfiguration_DestinationPortsPropertyList
	// Experimental.
	DestinationPortsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Protocols() *[]*float64
	// Experimental.
	SetProtocols(val *[]*float64)
	// Experimental.
	ProtocolsInput() *[]*float64
	// Experimental.
	Source() TfTlsInspectionConfiguration_SourcePropertyList
	// Experimental.
	SourceInput() interface{}
	// Experimental.
	SourcePorts() TfTlsInspectionConfiguration_SourcePortsPropertyList
	// Experimental.
	SourcePortsInput() interface{}
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
	PutDestinationPorts(value interface{})
	// Experimental.
	PutSource(value interface{})
	// Experimental.
	PutSourcePorts(value interface{})
	// Experimental.
	ResetDestination()
	// Experimental.
	ResetDestinationPorts()
	// Experimental.
	ResetSource()
	// Experimental.
	ResetSourcePorts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTlsInspectionConfiguration_ScopePropertyOutputReference
type jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) Destination() TfTlsInspectionConfiguration_DestinationPropertyList {
	var returns TfTlsInspectionConfiguration_DestinationPropertyList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) DestinationPorts() TfTlsInspectionConfiguration_DestinationPortsPropertyList {
	var returns TfTlsInspectionConfiguration_DestinationPortsPropertyList
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) DestinationPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) Protocols() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ProtocolsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) Source() TfTlsInspectionConfiguration_SourcePropertyList {
	var returns TfTlsInspectionConfiguration_SourcePropertyList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) SourcePorts() TfTlsInspectionConfiguration_SourcePortsPropertyList {
	var returns TfTlsInspectionConfiguration_SourcePortsPropertyList
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) SourcePortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTlsInspectionConfiguration_ScopePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTlsInspectionConfiguration_ScopePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTlsInspectionConfiguration_ScopePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfTlsInspectionConfiguration.ScopePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTlsInspectionConfiguration_ScopePropertyOutputReference_Override(t TfTlsInspectionConfiguration_ScopePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfTlsInspectionConfiguration.ScopePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetProtocols(val *[]*float64) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) PutDestination(value interface{}) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) PutDestinationPorts(value interface{}) {
	if err := t.validatePutDestinationPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationPorts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) PutSource(value interface{}) {
	if err := t.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) PutSourcePorts(value interface{}) {
	if err := t.validatePutSourcePortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourcePorts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		t,
		"resetSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTlsInspectionConfiguration_ScopePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

