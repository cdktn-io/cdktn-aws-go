package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_FailoverConditionSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioSilenceSettings() TfChannel_AudioSilenceSettingsPropertyOutputReference
	// Experimental.
	AudioSilenceSettingsInput() *TfChannel_AudioSilenceSettingsProperty
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
	InputLossSettings() TfChannel_InputLossSettingsPropertyOutputReference
	// Experimental.
	InputLossSettingsInput() *TfChannel_InputLossSettingsProperty
	// Experimental.
	InternalValue() *TfChannel_FailoverConditionSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_FailoverConditionSettingsProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VideoBlackSettings() TfChannel_VideoBlackSettingsPropertyOutputReference
	// Experimental.
	VideoBlackSettingsInput() *TfChannel_VideoBlackSettingsProperty
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
	PutAudioSilenceSettings(value *TfChannel_AudioSilenceSettingsProperty)
	// Experimental.
	PutInputLossSettings(value *TfChannel_InputLossSettingsProperty)
	// Experimental.
	PutVideoBlackSettings(value *TfChannel_VideoBlackSettingsProperty)
	// Experimental.
	ResetAudioSilenceSettings()
	// Experimental.
	ResetInputLossSettings()
	// Experimental.
	ResetVideoBlackSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_FailoverConditionSettingsPropertyOutputReference
type jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) AudioSilenceSettings() TfChannel_AudioSilenceSettingsPropertyOutputReference {
	var returns TfChannel_AudioSilenceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioSilenceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) AudioSilenceSettingsInput() *TfChannel_AudioSilenceSettingsProperty {
	var returns *TfChannel_AudioSilenceSettingsProperty
	_jsii_.Get(
		j,
		"audioSilenceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) InputLossSettings() TfChannel_InputLossSettingsPropertyOutputReference {
	var returns TfChannel_InputLossSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"inputLossSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) InputLossSettingsInput() *TfChannel_InputLossSettingsProperty {
	var returns *TfChannel_InputLossSettingsProperty
	_jsii_.Get(
		j,
		"inputLossSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) InternalValue() *TfChannel_FailoverConditionSettingsProperty {
	var returns *TfChannel_FailoverConditionSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) VideoBlackSettings() TfChannel_VideoBlackSettingsPropertyOutputReference {
	var returns TfChannel_VideoBlackSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"videoBlackSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) VideoBlackSettingsInput() *TfChannel_VideoBlackSettingsProperty {
	var returns *TfChannel_VideoBlackSettingsProperty
	_jsii_.Get(
		j,
		"videoBlackSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_FailoverConditionSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_FailoverConditionSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_FailoverConditionSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.FailoverConditionSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_FailoverConditionSettingsPropertyOutputReference_Override(t TfChannel_FailoverConditionSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.FailoverConditionSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_FailoverConditionSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) PutAudioSilenceSettings(value *TfChannel_AudioSilenceSettingsProperty) {
	if err := t.validatePutAudioSilenceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioSilenceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) PutInputLossSettings(value *TfChannel_InputLossSettingsProperty) {
	if err := t.validatePutInputLossSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputLossSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) PutVideoBlackSettings(value *TfChannel_VideoBlackSettingsProperty) {
	if err := t.validatePutVideoBlackSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVideoBlackSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ResetAudioSilenceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioSilenceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ResetInputLossSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetInputLossSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ResetVideoBlackSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoBlackSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_FailoverConditionSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

