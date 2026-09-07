package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_InputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioSelector() AwsChannel_AudioSelectorPropertyList
	// Experimental.
	AudioSelectorInput() interface{}
	// Experimental.
	CaptionSelector() AwsChannel_CaptionSelectorPropertyList
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
	InternalValue() *AwsChannel_InputSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_InputSettingsProperty)
	// Experimental.
	NetworkInputSettings() AwsChannel_NetworkInputSettingsPropertyOutputReference
	// Experimental.
	NetworkInputSettingsInput() *AwsChannel_NetworkInputSettingsProperty
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
	VideoSelector() AwsChannel_VideoSelectorPropertyOutputReference
	// Experimental.
	VideoSelectorInput() *AwsChannel_VideoSelectorProperty
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
	PutNetworkInputSettings(value *AwsChannel_NetworkInputSettingsProperty)
	// Experimental.
	PutVideoSelector(value *AwsChannel_VideoSelectorProperty)
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

// The jsii proxy struct for AwsChannel_InputSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) AudioSelector() AwsChannel_AudioSelectorPropertyList {
	var returns AwsChannel_AudioSelectorPropertyList
	_jsii_.Get(
		j,
		"audioSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) AudioSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) CaptionSelector() AwsChannel_CaptionSelectorPropertyList {
	var returns AwsChannel_CaptionSelectorPropertyList
	_jsii_.Get(
		j,
		"captionSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) CaptionSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) DeblockFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) DeblockFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deblockFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) DenoiseFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) DenoiseFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denoiseFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) FilterStrength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) FilterStrengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filterStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) InputFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) InputFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) InternalValue() *AwsChannel_InputSettingsProperty {
	var returns *AwsChannel_InputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) NetworkInputSettings() AwsChannel_NetworkInputSettingsPropertyOutputReference {
	var returns AwsChannel_NetworkInputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkInputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) NetworkInputSettingsInput() *AwsChannel_NetworkInputSettingsProperty {
	var returns *AwsChannel_NetworkInputSettingsProperty
	_jsii_.Get(
		j,
		"networkInputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Scte35Pid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Scte35PidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Smpte2038DataPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Smpte2038DataPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smpte2038DataPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) SourceEndBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) SourceEndBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) VideoSelector() AwsChannel_VideoSelectorPropertyOutputReference {
	var returns AwsChannel_VideoSelectorPropertyOutputReference
	_jsii_.Get(
		j,
		"videoSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) VideoSelectorInput() *AwsChannel_VideoSelectorProperty {
	var returns *AwsChannel_VideoSelectorProperty
	_jsii_.Get(
		j,
		"videoSelectorInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_InputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_InputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_InputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_InputSettingsPropertyOutputReference_Override(a AwsChannel_InputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetDeblockFilter(val *string) {
	if err := j.validateSetDeblockFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deblockFilter",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetDenoiseFilter(val *string) {
	if err := j.validateSetDenoiseFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denoiseFilter",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetFilterStrength(val *float64) {
	if err := j.validateSetFilterStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterStrength",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetInputFilter(val *string) {
	if err := j.validateSetInputFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFilter",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_InputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetScte35Pid(val *float64) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetSmpte2038DataPreference(val *string) {
	if err := j.validateSetSmpte2038DataPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smpte2038DataPreference",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetSourceEndBehavior(val *string) {
	if err := j.validateSetSourceEndBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) PutAudioSelector(value interface{}) {
	if err := a.validatePutAudioSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioSelector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) PutCaptionSelector(value interface{}) {
	if err := a.validatePutCaptionSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptionSelector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) PutNetworkInputSettings(value *AwsChannel_NetworkInputSettingsProperty) {
	if err := a.validatePutNetworkInputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkInputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) PutVideoSelector(value *AwsChannel_VideoSelectorProperty) {
	if err := a.validatePutVideoSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVideoSelector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetAudioSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetCaptionSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetDeblockFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetDeblockFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetDenoiseFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetDenoiseFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetFilterStrength() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterStrength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetInputFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetInputFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetNetworkInputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		a,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetSmpte2038DataPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetSmpte2038DataPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetSourceEndBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceEndBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ResetVideoSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_InputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

