package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference interface {
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
	DestinationPortRange() TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangePropertyOutputReference
	// Experimental.
	DestinationPortRangeInput() *TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEc2NetworkInsightsPath_FilterAtSourceProperty
	// Experimental.
	SetInternalValue(val *TfEc2NetworkInsightsPath_FilterAtSourceProperty)
	// Experimental.
	SourceAddress() *string
	// Experimental.
	SetSourceAddress(val *string)
	// Experimental.
	SourceAddressInput() *string
	// Experimental.
	SourcePortRange() TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangePropertyOutputReference
	// Experimental.
	SourcePortRangeInput() *TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty
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
	PutDestinationPortRange(value *TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty)
	// Experimental.
	PutSourcePortRange(value *TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty)
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

// The jsii proxy struct for TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference
type jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) DestinationPortRange() TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangePropertyOutputReference {
	var returns TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) DestinationPortRangeInput() *TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) InternalValue() *TfEc2NetworkInsightsPath_FilterAtSourceProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtSourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) SourcePortRange() TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangePropertyOutputReference {
	var returns TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) SourcePortRangeInput() *TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty {
	var returns *TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsPath.FilterAtSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference_Override(t TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2NetworkInsightsPath.FilterAtSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetInternalValue(val *TfEc2NetworkInsightsPath_FilterAtSourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) PutDestinationPortRange(value *TfEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty) {
	if err := t.validatePutDestinationPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationPortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) PutSourcePortRange(value *TfEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty) {
	if err := t.validatePutSourcePortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourcePortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEc2NetworkInsightsPath_FilterAtSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

