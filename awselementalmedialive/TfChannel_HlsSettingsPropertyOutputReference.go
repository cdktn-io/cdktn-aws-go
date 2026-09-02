package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_HlsSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioOnlyHlsSettings() TfChannel_AudioOnlyHlsSettingsPropertyOutputReference
	// Experimental.
	AudioOnlyHlsSettingsInput() *TfChannel_AudioOnlyHlsSettingsProperty
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
	Fmp4HlsSettings() TfChannel_Fmp4HlsSettingsPropertyOutputReference
	// Experimental.
	Fmp4HlsSettingsInput() *TfChannel_Fmp4HlsSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameCaptureHlsSettings() TfChannel_FrameCaptureHlsSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureHlsSettingsInput() *TfChannel_FrameCaptureHlsSettingsProperty
	// Experimental.
	InternalValue() *TfChannel_HlsSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_HlsSettingsProperty)
	// Experimental.
	StandardHlsSettings() TfChannel_StandardHlsSettingsPropertyOutputReference
	// Experimental.
	StandardHlsSettingsInput() *TfChannel_StandardHlsSettingsProperty
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
	PutAudioOnlyHlsSettings(value *TfChannel_AudioOnlyHlsSettingsProperty)
	// Experimental.
	PutFmp4HlsSettings(value *TfChannel_Fmp4HlsSettingsProperty)
	// Experimental.
	PutFrameCaptureHlsSettings(value *TfChannel_FrameCaptureHlsSettingsProperty)
	// Experimental.
	PutStandardHlsSettings(value *TfChannel_StandardHlsSettingsProperty)
	// Experimental.
	ResetAudioOnlyHlsSettings()
	// Experimental.
	ResetFmp4HlsSettings()
	// Experimental.
	ResetFrameCaptureHlsSettings()
	// Experimental.
	ResetStandardHlsSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_HlsSettingsPropertyOutputReference
type jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) AudioOnlyHlsSettings() TfChannel_AudioOnlyHlsSettingsPropertyOutputReference {
	var returns TfChannel_AudioOnlyHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioOnlyHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) AudioOnlyHlsSettingsInput() *TfChannel_AudioOnlyHlsSettingsProperty {
	var returns *TfChannel_AudioOnlyHlsSettingsProperty
	_jsii_.Get(
		j,
		"audioOnlyHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) Fmp4HlsSettings() TfChannel_Fmp4HlsSettingsPropertyOutputReference {
	var returns TfChannel_Fmp4HlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"fmp4HlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) Fmp4HlsSettingsInput() *TfChannel_Fmp4HlsSettingsProperty {
	var returns *TfChannel_Fmp4HlsSettingsProperty
	_jsii_.Get(
		j,
		"fmp4HlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) FrameCaptureHlsSettings() TfChannel_FrameCaptureHlsSettingsPropertyOutputReference {
	var returns TfChannel_FrameCaptureHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) FrameCaptureHlsSettingsInput() *TfChannel_FrameCaptureHlsSettingsProperty {
	var returns *TfChannel_FrameCaptureHlsSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) InternalValue() *TfChannel_HlsSettingsProperty {
	var returns *TfChannel_HlsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) StandardHlsSettings() TfChannel_StandardHlsSettingsPropertyOutputReference {
	var returns TfChannel_StandardHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"standardHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) StandardHlsSettingsInput() *TfChannel_StandardHlsSettingsProperty {
	var returns *TfChannel_StandardHlsSettingsProperty
	_jsii_.Get(
		j,
		"standardHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_HlsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_HlsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_HlsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_HlsSettingsPropertyOutputReference_Override(t TfChannel_HlsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_HlsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) PutAudioOnlyHlsSettings(value *TfChannel_AudioOnlyHlsSettingsProperty) {
	if err := t.validatePutAudioOnlyHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioOnlyHlsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) PutFmp4HlsSettings(value *TfChannel_Fmp4HlsSettingsProperty) {
	if err := t.validatePutFmp4HlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFmp4HlsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) PutFrameCaptureHlsSettings(value *TfChannel_FrameCaptureHlsSettingsProperty) {
	if err := t.validatePutFrameCaptureHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFrameCaptureHlsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) PutStandardHlsSettings(value *TfChannel_StandardHlsSettingsProperty) {
	if err := t.validatePutStandardHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStandardHlsSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ResetAudioOnlyHlsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioOnlyHlsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ResetFmp4HlsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFmp4HlsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ResetFrameCaptureHlsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameCaptureHlsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ResetStandardHlsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetStandardHlsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_HlsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

