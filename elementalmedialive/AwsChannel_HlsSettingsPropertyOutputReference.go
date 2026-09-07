package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_HlsSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioOnlyHlsSettings() AwsChannel_AudioOnlyHlsSettingsPropertyOutputReference
	// Experimental.
	AudioOnlyHlsSettingsInput() *AwsChannel_AudioOnlyHlsSettingsProperty
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
	Fmp4HlsSettings() AwsChannel_Fmp4HlsSettingsPropertyOutputReference
	// Experimental.
	Fmp4HlsSettingsInput() *AwsChannel_Fmp4HlsSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameCaptureHlsSettings() AwsChannel_FrameCaptureHlsSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureHlsSettingsInput() *AwsChannel_FrameCaptureHlsSettingsProperty
	// Experimental.
	InternalValue() *AwsChannel_HlsSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_HlsSettingsProperty)
	// Experimental.
	StandardHlsSettings() AwsChannel_StandardHlsSettingsPropertyOutputReference
	// Experimental.
	StandardHlsSettingsInput() *AwsChannel_StandardHlsSettingsProperty
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
	PutAudioOnlyHlsSettings(value *AwsChannel_AudioOnlyHlsSettingsProperty)
	// Experimental.
	PutFmp4HlsSettings(value *AwsChannel_Fmp4HlsSettingsProperty)
	// Experimental.
	PutFrameCaptureHlsSettings(value *AwsChannel_FrameCaptureHlsSettingsProperty)
	// Experimental.
	PutStandardHlsSettings(value *AwsChannel_StandardHlsSettingsProperty)
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

// The jsii proxy struct for AwsChannel_HlsSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) AudioOnlyHlsSettings() AwsChannel_AudioOnlyHlsSettingsPropertyOutputReference {
	var returns AwsChannel_AudioOnlyHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioOnlyHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) AudioOnlyHlsSettingsInput() *AwsChannel_AudioOnlyHlsSettingsProperty {
	var returns *AwsChannel_AudioOnlyHlsSettingsProperty
	_jsii_.Get(
		j,
		"audioOnlyHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) Fmp4HlsSettings() AwsChannel_Fmp4HlsSettingsPropertyOutputReference {
	var returns AwsChannel_Fmp4HlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"fmp4HlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) Fmp4HlsSettingsInput() *AwsChannel_Fmp4HlsSettingsProperty {
	var returns *AwsChannel_Fmp4HlsSettingsProperty
	_jsii_.Get(
		j,
		"fmp4HlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) FrameCaptureHlsSettings() AwsChannel_FrameCaptureHlsSettingsPropertyOutputReference {
	var returns AwsChannel_FrameCaptureHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) FrameCaptureHlsSettingsInput() *AwsChannel_FrameCaptureHlsSettingsProperty {
	var returns *AwsChannel_FrameCaptureHlsSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) InternalValue() *AwsChannel_HlsSettingsProperty {
	var returns *AwsChannel_HlsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) StandardHlsSettings() AwsChannel_StandardHlsSettingsPropertyOutputReference {
	var returns AwsChannel_StandardHlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"standardHlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) StandardHlsSettingsInput() *AwsChannel_StandardHlsSettingsProperty {
	var returns *AwsChannel_StandardHlsSettingsProperty
	_jsii_.Get(
		j,
		"standardHlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_HlsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_HlsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_HlsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_HlsSettingsPropertyOutputReference_Override(a AwsChannel_HlsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_HlsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) PutAudioOnlyHlsSettings(value *AwsChannel_AudioOnlyHlsSettingsProperty) {
	if err := a.validatePutAudioOnlyHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioOnlyHlsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) PutFmp4HlsSettings(value *AwsChannel_Fmp4HlsSettingsProperty) {
	if err := a.validatePutFmp4HlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFmp4HlsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) PutFrameCaptureHlsSettings(value *AwsChannel_FrameCaptureHlsSettingsProperty) {
	if err := a.validatePutFrameCaptureHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureHlsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) PutStandardHlsSettings(value *AwsChannel_StandardHlsSettingsProperty) {
	if err := a.validatePutStandardHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStandardHlsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ResetAudioOnlyHlsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioOnlyHlsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ResetFmp4HlsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFmp4HlsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ResetFrameCaptureHlsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureHlsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ResetStandardHlsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStandardHlsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

