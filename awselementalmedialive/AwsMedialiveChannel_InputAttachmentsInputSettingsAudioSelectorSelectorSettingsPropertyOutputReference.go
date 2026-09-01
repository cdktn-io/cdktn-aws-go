package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioHlsRenditionSelection() AwsMedialiveChannel_AudioHlsRenditionSelectionPropertyOutputReference
	// Experimental.
	AudioHlsRenditionSelectionInput() *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty
	// Experimental.
	AudioLanguageSelection() AwsMedialiveChannel_AudioLanguageSelectionPropertyOutputReference
	// Experimental.
	AudioLanguageSelectionInput() *AwsMedialiveChannel_AudioLanguageSelectionProperty
	// Experimental.
	AudioPidSelection() AwsMedialiveChannel_AudioPidSelectionPropertyOutputReference
	// Experimental.
	AudioPidSelectionInput() *AwsMedialiveChannel_AudioPidSelectionProperty
	// Experimental.
	AudioTrackSelection() AwsMedialiveChannel_AudioTrackSelectionPropertyOutputReference
	// Experimental.
	AudioTrackSelectionInput() *AwsMedialiveChannel_AudioTrackSelectionProperty
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
	InternalValue() *AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty)
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
	PutAudioHlsRenditionSelection(value *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty)
	// Experimental.
	PutAudioLanguageSelection(value *AwsMedialiveChannel_AudioLanguageSelectionProperty)
	// Experimental.
	PutAudioPidSelection(value *AwsMedialiveChannel_AudioPidSelectionProperty)
	// Experimental.
	PutAudioTrackSelection(value *AwsMedialiveChannel_AudioTrackSelectionProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelection() AwsMedialiveChannel_AudioHlsRenditionSelectionPropertyOutputReference {
	var returns AwsMedialiveChannel_AudioHlsRenditionSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioHlsRenditionSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelectionInput() *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty {
	var returns *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty
	_jsii_.Get(
		j,
		"audioHlsRenditionSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelection() AwsMedialiveChannel_AudioLanguageSelectionPropertyOutputReference {
	var returns AwsMedialiveChannel_AudioLanguageSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioLanguageSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelectionInput() *AwsMedialiveChannel_AudioLanguageSelectionProperty {
	var returns *AwsMedialiveChannel_AudioLanguageSelectionProperty
	_jsii_.Get(
		j,
		"audioLanguageSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelection() AwsMedialiveChannel_AudioPidSelectionPropertyOutputReference {
	var returns AwsMedialiveChannel_AudioPidSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioPidSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelectionInput() *AwsMedialiveChannel_AudioPidSelectionProperty {
	var returns *AwsMedialiveChannel_AudioPidSelectionProperty
	_jsii_.Get(
		j,
		"audioPidSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelection() AwsMedialiveChannel_AudioTrackSelectionPropertyOutputReference {
	var returns AwsMedialiveChannel_AudioTrackSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioTrackSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelectionInput() *AwsMedialiveChannel_AudioTrackSelectionProperty {
	var returns *AwsMedialiveChannel_AudioTrackSelectionProperty
	_jsii_.Get(
		j,
		"audioTrackSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty {
	var returns *AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioHlsRenditionSelection(value *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty) {
	if err := a.validatePutAudioHlsRenditionSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioHlsRenditionSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioLanguageSelection(value *AwsMedialiveChannel_AudioLanguageSelectionProperty) {
	if err := a.validatePutAudioLanguageSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioLanguageSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioPidSelection(value *AwsMedialiveChannel_AudioPidSelectionProperty) {
	if err := a.validatePutAudioPidSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioPidSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioTrackSelection(value *AwsMedialiveChannel_AudioTrackSelectionProperty) {
	if err := a.validatePutAudioTrackSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioTrackSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioHlsRenditionSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioHlsRenditionSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioLanguageSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioLanguageSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioPidSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioPidSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioTrackSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioTrackSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

