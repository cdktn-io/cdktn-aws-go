package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference interface {
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
	DestinationAddresses() *[]*string
	// Experimental.
	SetDestinationAddresses(val *[]*string)
	// Experimental.
	DestinationAddressesInput() *[]*string
	// Experimental.
	DestinationPorts() *[]*string
	// Experimental.
	SetDestinationPorts(val *[]*string)
	// Experimental.
	DestinationPortsInput() *[]*string
	// Experimental.
	DestinationPrefixLists() *[]*string
	// Experimental.
	SetDestinationPrefixLists(val *[]*string)
	// Experimental.
	DestinationPrefixListsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Protocols() *[]*string
	// Experimental.
	SetProtocols(val *[]*string)
	// Experimental.
	ProtocolsInput() *[]*string
	// Experimental.
	SourceAddresses() *[]*string
	// Experimental.
	SetSourceAddresses(val *[]*string)
	// Experimental.
	SourceAddressesInput() *[]*string
	// Experimental.
	SourcePorts() *[]*string
	// Experimental.
	SetSourcePorts(val *[]*string)
	// Experimental.
	SourcePortsInput() *[]*string
	// Experimental.
	SourcePrefixLists() *[]*string
	// Experimental.
	SetSourcePrefixLists(val *[]*string)
	// Experimental.
	SourcePrefixListsInput() *[]*string
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
	ResetDestinationAddresses()
	// Experimental.
	ResetDestinationPorts()
	// Experimental.
	ResetDestinationPrefixLists()
	// Experimental.
	ResetProtocols()
	// Experimental.
	ResetSourceAddresses()
	// Experimental.
	ResetSourcePorts()
	// Experimental.
	ResetSourcePrefixLists()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference
type jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationPrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) DestinationPrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourcePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourcePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourcePrefixLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) SourcePrefixListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePrefixListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAccessScope.ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference_Override(t TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsAccessScope.ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetDestinationAddresses(val *[]*string) {
	if err := j.validateSetDestinationAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddresses",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetDestinationPrefixLists(val *[]*string) {
	if err := j.validateSetDestinationPrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixLists",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetSourcePorts(val *[]*string) {
	if err := j.validateSetSourcePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePorts",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetSourcePrefixLists(val *[]*string) {
	if err := j.validateSetSourcePrefixListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePrefixLists",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetDestinationAddresses() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationAddresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetDestinationPrefixLists() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPrefixLists",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetSourcePorts() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ResetSourcePrefixLists() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePrefixLists",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

