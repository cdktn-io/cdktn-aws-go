package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioHlsRenditionSelection() AwsChannel_AudioHlsRenditionSelectionPropertyOutputReference
	// Experimental.
	AudioHlsRenditionSelectionInput() *AwsChannel_AudioHlsRenditionSelectionProperty
	// Experimental.
	AudioLanguageSelection() AwsChannel_AudioLanguageSelectionPropertyOutputReference
	// Experimental.
	AudioLanguageSelectionInput() *AwsChannel_AudioLanguageSelectionProperty
	// Experimental.
	AudioPidSelection() AwsChannel_AudioPidSelectionPropertyOutputReference
	// Experimental.
	AudioPidSelectionInput() *AwsChannel_AudioPidSelectionProperty
	// Experimental.
	AudioTrackSelection() AwsChannel_AudioTrackSelectionPropertyOutputReference
	// Experimental.
	AudioTrackSelectionInput() *AwsChannel_AudioTrackSelectionProperty
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
	InternalValue() *AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty)
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
	PutAudioHlsRenditionSelection(value *AwsChannel_AudioHlsRenditionSelectionProperty)
	// Experimental.
	PutAudioLanguageSelection(value *AwsChannel_AudioLanguageSelectionProperty)
	// Experimental.
	PutAudioPidSelection(value *AwsChannel_AudioPidSelectionProperty)
	// Experimental.
	PutAudioTrackSelection(value *AwsChannel_AudioTrackSelectionProperty)
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

// The jsii proxy struct for AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelection() AwsChannel_AudioHlsRenditionSelectionPropertyOutputReference {
	var returns AwsChannel_AudioHlsRenditionSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioHlsRenditionSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioHlsRenditionSelectionInput() *AwsChannel_AudioHlsRenditionSelectionProperty {
	var returns *AwsChannel_AudioHlsRenditionSelectionProperty
	_jsii_.Get(
		j,
		"audioHlsRenditionSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelection() AwsChannel_AudioLanguageSelectionPropertyOutputReference {
	var returns AwsChannel_AudioLanguageSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioLanguageSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioLanguageSelectionInput() *AwsChannel_AudioLanguageSelectionProperty {
	var returns *AwsChannel_AudioLanguageSelectionProperty
	_jsii_.Get(
		j,
		"audioLanguageSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelection() AwsChannel_AudioPidSelectionPropertyOutputReference {
	var returns AwsChannel_AudioPidSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioPidSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioPidSelectionInput() *AwsChannel_AudioPidSelectionProperty {
	var returns *AwsChannel_AudioPidSelectionProperty
	_jsii_.Get(
		j,
		"audioPidSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelection() AwsChannel_AudioTrackSelectionPropertyOutputReference {
	var returns AwsChannel_AudioTrackSelectionPropertyOutputReference
	_jsii_.Get(
		j,
		"audioTrackSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) AudioTrackSelectionInput() *AwsChannel_AudioTrackSelectionProperty {
	var returns *AwsChannel_AudioTrackSelectionProperty
	_jsii_.Get(
		j,
		"audioTrackSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InternalValue() *AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty {
	var returns *AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference_Override(a AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioHlsRenditionSelection(value *AwsChannel_AudioHlsRenditionSelectionProperty) {
	if err := a.validatePutAudioHlsRenditionSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioHlsRenditionSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioLanguageSelection(value *AwsChannel_AudioLanguageSelectionProperty) {
	if err := a.validatePutAudioLanguageSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioLanguageSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioPidSelection(value *AwsChannel_AudioPidSelectionProperty) {
	if err := a.validatePutAudioPidSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioPidSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) PutAudioTrackSelection(value *AwsChannel_AudioTrackSelectionProperty) {
	if err := a.validatePutAudioTrackSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioTrackSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioHlsRenditionSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioHlsRenditionSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioLanguageSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioLanguageSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioPidSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioPidSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ResetAudioTrackSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioTrackSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

