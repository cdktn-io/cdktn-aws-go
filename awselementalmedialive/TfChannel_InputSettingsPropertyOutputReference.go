package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_InputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioSelector() TfChannel_AudioSelectorPropertyList
	// Experimental.
	AudioSelectorInput() interface{}
	// Experimental.
	CaptionSelector() TfChannel_CaptionSelectorPropertyList
	// Experimental.
	CaptionSelectorInput() interface{}
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
	DeblockFilter() *string
	// Experimental.
	SetDeblockFilter(val *string)
	// Experimental.
	DeblockFilterInput() *string
	// Experimental.
	DenoiseFilter() *string
	// Experimental.
	SetDenoiseFilter(val *string)
	// Experimental.
	DenoiseFilterInput() *string
	// Experimental.
	FilterStrength() *float64
	// Experimental.
	SetFilterStrength(val *float64)
	// Experimental.
	FilterStrengthInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InputFilter() *string
	// Experimental.
	SetInputFilter(val *string)
	// Experimental.
	InputFilterInput() *string
	// Experimental.
	InternalValue() *TfChannel_InputSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_InputSettingsProperty)
	// Experimental.
	NetworkInputSettings() TfChannel_NetworkInputSettingsPropertyOutputReference
	// Experimental.
	NetworkInputSettingsInput() *TfChannel_NetworkInputSettingsProperty
	// Experimental.
	Scte35Pid() *float64
	// Experimental.
	SetScte35Pid(val *float64)
	// Experimental.
	Scte35PidInput() *float64
	// Experimental.
	Smpte2038DataPreference() *string
	// Experimental.
	SetSmpte2038DataPreference(val *string)
	// Experimental.
	Smpte2038DataPreferenceInput() *string
	// Experimental.
	SourceEndBehavior() *string
	// Experimental.
	SetSourceEndBehavior(val *string)
	// Experimental.
	SourceEndBehaviorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VideoSelector() TfChannel_VideoSelectorPropertyOutputReference
	// Experimental.
	VideoSelectorInput() *TfChannel_VideoSelectorProperty
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
	PutAudioSelector(value interface{})
	// Experimental.
	PutCaptionSelector(value interface{})
	// Experimental.
	PutNetworkInputSettings(value *TfChannel_NetworkInputSettingsProperty)
	// Experimental.
	PutVideoSelector(value *TfChannel_VideoSelectorProperty)
	// Experimental.
	ResetAudioSelector()
	// Experimental.
	ResetCaptionSelector()
	// Experimental.
	ResetDeblockFilter()
	// Experimental.
	ResetDenoiseFilter()
	// Experimental.
	ResetFilterStrength()
	// Experimental.
	ResetInputFilter()
	// Experimental.
	ResetNetworkInputSettings()
	// Experimental.
	ResetScte35Pid()
	// Experimental.
	ResetSmpte2038DataPreference()
	// Experimental.
	ResetSourceEndBehavior()
	// Experimental.
	ResetVideoSelector()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_InputSettingsPropertyOutputReference
type jsiiProxy_TfChannel_InputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) AudioSelector() TfChannel_AudioSelectorPropertyList {
	var returns TfChannel_AudioSelectorPropertyList
	_jsii_.Get(
		j,
		"audioSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) AudioSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) CaptionSelector() TfChannel_CaptionSelectorPropertyList {
	var returns TfChannel_CaptionSelectorPropertyList
	_jsii_.Get(
		j,
		"captionSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) CaptionSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) DeblockFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) DeblockFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) DenoiseFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) DenoiseFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) FilterStrength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) FilterStrengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) InputFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) InputFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) InternalValue() *TfChannel_InputSettingsProperty {
	var returns *TfChannel_InputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) NetworkInputSettings() TfChannel_NetworkInputSettingsPropertyOutputReference {
	var returns TfChannel_NetworkInputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkInputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) NetworkInputSettingsInput() *TfChannel_NetworkInputSettingsProperty {
	var returns *TfChannel_NetworkInputSettingsProperty
	_jsii_.Get(
		j,
		"networkInputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Scte35Pid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Scte35PidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Smpte2038DataPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Smpte2038DataPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) SourceEndBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) SourceEndBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) VideoSelector() TfChannel_VideoSelectorPropertyOutputReference {
	var returns TfChannel_VideoSelectorPropertyOutputReference
	_jsii_.Get(
		j,
		"videoSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) VideoSelectorInput() *TfChannel_VideoSelectorProperty {
	var returns *TfChannel_VideoSelectorProperty
	_jsii_.Get(
		j,
		"videoSelectorInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_InputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_InputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_InputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_InputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_InputSettingsPropertyOutputReference_Override(t TfChannel_InputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetDeblockFilter(val *string) {
	if err := j.validateSetDeblockFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deblockFilter",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetDenoiseFilter(val *string) {
	if err := j.validateSetDenoiseFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denoiseFilter",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetFilterStrength(val *float64) {
	if err := j.validateSetFilterStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterStrength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetInputFilter(val *string) {
	if err := j.validateSetInputFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFilter",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_InputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetScte35Pid(val *float64) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetSmpte2038DataPreference(val *string) {
	if err := j.validateSetSmpte2038DataPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smpte2038DataPreference",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetSourceEndBehavior(val *string) {
	if err := j.validateSetSourceEndBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) PutAudioSelector(value interface{}) {
	if err := t.validatePutAudioSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) PutCaptionSelector(value interface{}) {
	if err := t.validatePutCaptionSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptionSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) PutNetworkInputSettings(value *TfChannel_NetworkInputSettingsProperty) {
	if err := t.validatePutNetworkInputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkInputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) PutVideoSelector(value *TfChannel_VideoSelectorProperty) {
	if err := t.validatePutVideoSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVideoSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetAudioSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetCaptionSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetDeblockFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetDeblockFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetDenoiseFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetDenoiseFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetFilterStrength() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterStrength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetInputFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetInputFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetNetworkInputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		t,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetSmpte2038DataPreference() {
	_jsii_.InvokeVoid(
		t,
		"resetSmpte2038DataPreference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetSourceEndBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceEndBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ResetVideoSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_InputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

