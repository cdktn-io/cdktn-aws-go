package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference interface {
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
	DestinationAddress() *string
	// Experimental.
	SetDestinationAddress(val *string)
	// Experimental.
	DestinationAddressInput() *string
	// Experimental.
	DestinationPortRange() TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference
	// Experimental.
	DestinationPortRangeInput() *TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEc2NetworkInsightsPath_FilterAtDestinationProperty
	// Experimental.
	SetInternalValue(val *TfEc2NetworkInsightsPath_FilterAtDestinationProperty)
	// Experimental.
	SourceAddress() *string
	// Experimental.
	SetSourceAddress(val *string)
	// Experimental.
	SourceAddressInput() *string
	// Experimental.
	SourcePortRange() TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference
	// Experimental.
	SourcePortRangeInput() *TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty
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
	PutDestinationPortRange(value *TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty)
	// Experimental.
	PutSourcePortRange(value *TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty)
	// Experimental.
	ResetDestinationAddress()
	// Experimental.
	ResetDestinationPortRange()
	// Experimental.
	ResetSourceAddress()
	// Experimental.
	ResetSourcePortRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference
type jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationPortRange() TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference {
	var returns TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) DestinationPortRangeInput() *TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InternalValue() *TfEc2NetworkInsightsPath_FilterAtDestinationProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourcePortRange() TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference {
	var returns TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) SourcePortRangeInput() *TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsPath.FilterAtDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference_Override(t TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsPath.FilterAtDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetInternalValue(val *TfEc2NetworkInsightsPath_FilterAtDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) PutDestinationPortRange(value *TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty) {
	if err := t.validatePutDestinationPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationPortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) PutSourcePortRange(value *TfEc2NetworkInsightsPath_FilterAtDestinationSourcePortRangeProperty) {
	if err := t.validatePutSourcePortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourcePortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

