package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_NielsenWatermarksSettingsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_NielsenWatermarksSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_NielsenWatermarksSettingsProperty)
	// Experimental.
	NielsenCbetSettings() TfChannel_NielsenCbetSettingsPropertyOutputReference
	// Experimental.
	NielsenCbetSettingsInput() *TfChannel_NielsenCbetSettingsProperty
	// Experimental.
	NielsenDistributionType() *string
	// Experimental.
	SetNielsenDistributionType(val *string)
	// Experimental.
	NielsenDistributionTypeInput() *string
	// Experimental.
	NielsenNaesIiNwSettings() TfChannel_NielsenNaesIiNwSettingsPropertyList
	// Experimental.
	NielsenNaesIiNwSettingsInput() interface{}
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
	PutNielsenCbetSettings(value *TfChannel_NielsenCbetSettingsProperty)
	// Experimental.
	PutNielsenNaesIiNwSettings(value interface{})
	// Experimental.
	ResetNielsenCbetSettings()
	// Experimental.
	ResetNielsenDistributionType()
	// Experimental.
	ResetNielsenNaesIiNwSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_NielsenWatermarksSettingsPropertyOutputReference
type jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) InternalValue() *TfChannel_NielsenWatermarksSettingsProperty {
	var returns *TfChannel_NielsenWatermarksSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenCbetSettings() TfChannel_NielsenCbetSettingsPropertyOutputReference {
	var returns TfChannel_NielsenCbetSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"nielsenCbetSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenCbetSettingsInput() *TfChannel_NielsenCbetSettingsProperty {
	var returns *TfChannel_NielsenCbetSettingsProperty
	_jsii_.Get(
		j,
		"nielsenCbetSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenNaesIiNwSettings() TfChannel_NielsenNaesIiNwSettingsPropertyList {
	var returns TfChannel_NielsenNaesIiNwSettingsPropertyList
	_jsii_.Get(
		j,
		"nielsenNaesIiNwSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenNaesIiNwSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nielsenNaesIiNwSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_NielsenWatermarksSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_NielsenWatermarksSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_NielsenWatermarksSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.NielsenWatermarksSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_NielsenWatermarksSettingsPropertyOutputReference_Override(t TfChannel_NielsenWatermarksSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.NielsenWatermarksSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_NielsenWatermarksSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetNielsenDistributionType(val *string) {
	if err := j.validateSetNielsenDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nielsenDistributionType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) PutNielsenCbetSettings(value *TfChannel_NielsenCbetSettingsProperty) {
	if err := t.validatePutNielsenCbetSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNielsenCbetSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) PutNielsenNaesIiNwSettings(value interface{}) {
	if err := t.validatePutNielsenNaesIiNwSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNielsenNaesIiNwSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenCbetSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetNielsenCbetSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenDistributionType() {
	_jsii_.InvokeVoid(
		t,
		"resetNielsenDistributionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenNaesIiNwSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetNielsenNaesIiNwSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_NielsenWatermarksSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

