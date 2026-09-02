package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioHlsRenditionSelection() TfChannel_AudioHlsRenditionSelectionPropertyOutputReference
	// Experimental.
	AudioHlsRenditionSelectionInput() *TfChannel_AudioHlsRenditionSelectionProperty
	// Experimental.
	AudioLanguageSelection() TfChannel_AudioLanguageSelectionPropertyOutputReference
	// Experimental.
	AudioLanguageSelectionInput() *TfChannel_AudioLanguageSelectionProperty
	// Experimental.
	AudioPidSelection() TfChannel_AudioPidSelectionPropertyOutputReference
	// Experimental.
	AudioPidSelectionInput() *TfChannel_AudioPidSelectionProperty
	// Experimental.
	AudioTrackSelection() TfChannel_AudioTrackSelectionPropertyOutputReference
	// Experimental.
	AudioTrackSelectionInput() *TfChannel_AudioTrackSelectionProperty
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
	InternalValue() *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty)
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
	PutAudioHlsRenditionSelection(value *TfChannel_AudioHlsRenditionSelectionProperty)
	// Experimental.
	PutAudioLanguageSelection(value *TfChannel_AudioLanguageSelectionProperty)
	// Experimental.
	PutAudioPidSelection(value *TfChannel_AudioPidSelectionProperty)
	// Experimental.
	PutAudioTrackSelection(value *TfChannel_AudioTrackSelectionProperty)
	// Experimental.
	ResetAudioHlsRenditionSelection()
	// Experimental.
	ResetAudioLanguageSelection()
	// Experimental.
	ResetAudioPidSelection()
	// Experimental.
	ResetAudioTrackSelection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference
type jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelection() TfChannel_AudioHlsRenditionSelectionPropertyOutputReference {
	var returns TfChannel_AudioHlsRenditionSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioHlsRenditionSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelectionInput() *TfChannel_AudioHlsRenditionSelectionProperty {
	var returns *TfChannel_AudioHlsRenditionSelectionProperty
	_jsii_.Get(
		j,
		"audioHlsRenditionSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelection() TfChannel_AudioLanguageSelectionPropertyOutputReference {
	var returns TfChannel_AudioLanguageSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioLanguageSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelectionInput() *TfChannel_AudioLanguageSelectionProperty {
	var returns *TfChannel_AudioLanguageSelectionProperty
	_jsii_.Get(
		j,
		"audioLanguageSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelection() TfChannel_AudioPidSelectionPropertyOutputReference {
	var returns TfChannel_AudioPidSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioPidSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelectionInput() *TfChannel_AudioPidSelectionProperty {
	var returns *TfChannel_AudioPidSelectionProperty
	_jsii_.Get(
		j,
		"audioPidSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelection() TfChannel_AudioTrackSelectionPropertyOutputReference {
	var returns TfChannel_AudioTrackSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioTrackSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelectionInput() *TfChannel_AudioTrackSelectionProperty {
	var returns *TfChannel_AudioTrackSelectionProperty
	_jsii_.Get(
		j,
		"audioTrackSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InternalValue() *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty {
	var returns *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference_Override(t TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioHlsRenditionSelection(value *TfChannel_AudioHlsRenditionSelectionProperty) {
	if err := t.validatePutAudioHlsRenditionSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioHlsRenditionSelection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioLanguageSelection(value *TfChannel_AudioLanguageSelectionProperty) {
	if err := t.validatePutAudioLanguageSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioLanguageSelection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioPidSelection(value *TfChannel_AudioPidSelectionProperty) {
	if err := t.validatePutAudioPidSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioPidSelection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioTrackSelection(value *TfChannel_AudioTrackSelectionProperty) {
	if err := t.validatePutAudioTrackSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioTrackSelection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioHlsRenditionSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioHlsRenditionSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioLanguageSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioLanguageSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioPidSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioPidSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioTrackSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioTrackSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

