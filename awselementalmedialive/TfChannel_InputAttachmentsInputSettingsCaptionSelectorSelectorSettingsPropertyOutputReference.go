package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AncillarySourceSettings() TfChannel_AncillarySourceSettingsPropertyOutputReference
	// Experimental.
	AncillarySourceSettingsInput() *TfChannel_AncillarySourceSettingsProperty
	// Experimental.
	AribSourceSettings() TfChannel_AribSourceSettingsPropertyOutputReference
	// Experimental.
	AribSourceSettingsInput() *TfChannel_AribSourceSettingsProperty
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
	DvbSubSourceSettings() TfChannel_DvbSubSourceSettingsPropertyOutputReference
	// Experimental.
	DvbSubSourceSettingsInput() *TfChannel_DvbSubSourceSettingsProperty
	// Experimental.
	EmbeddedSourceSettings() TfChannel_EmbeddedSourceSettingsPropertyOutputReference
	// Experimental.
	EmbeddedSourceSettingsInput() *TfChannel_EmbeddedSourceSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty)
	// Experimental.
	Scte20SourceSettings() TfChannel_Scte20SourceSettingsPropertyOutputReference
	// Experimental.
	Scte20SourceSettingsInput() *TfChannel_Scte20SourceSettingsProperty
	// Experimental.
	Scte27SourceSettings() TfChannel_Scte27SourceSettingsPropertyOutputReference
	// Experimental.
	Scte27SourceSettingsInput() *TfChannel_Scte27SourceSettingsProperty
	// Experimental.
	TeletextSourceSettings() TfChannel_TeletextSourceSettingsPropertyOutputReference
	// Experimental.
	TeletextSourceSettingsInput() *TfChannel_TeletextSourceSettingsProperty
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
	PutAncillarySourceSettings(value *TfChannel_AncillarySourceSettingsProperty)
	// Experimental.
	PutAribSourceSettings(value *TfChannel_AribSourceSettingsProperty)
	// Experimental.
	PutDvbSubSourceSettings(value *TfChannel_DvbSubSourceSettingsProperty)
	// Experimental.
	PutEmbeddedSourceSettings(value *TfChannel_EmbeddedSourceSettingsProperty)
	// Experimental.
	PutScte20SourceSettings(value *TfChannel_Scte20SourceSettingsProperty)
	// Experimental.
	PutScte27SourceSettings(value *TfChannel_Scte27SourceSettingsProperty)
	// Experimental.
	PutTeletextSourceSettings(value *TfChannel_TeletextSourceSettingsProperty)
	// Experimental.
	ResetAncillarySourceSettings()
	// Experimental.
	ResetAribSourceSettings()
	// Experimental.
	ResetDvbSubSourceSettings()
	// Experimental.
	ResetEmbeddedSourceSettings()
	// Experimental.
	ResetScte20SourceSettings()
	// Experimental.
	ResetScte27SourceSettings()
	// Experimental.
	ResetTeletextSourceSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference
type jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AncillarySourceSettings() TfChannel_AncillarySourceSettingsPropertyOutputReference {
	var returns TfChannel_AncillarySourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ancillarySourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AncillarySourceSettingsInput() *TfChannel_AncillarySourceSettingsProperty {
	var returns *TfChannel_AncillarySourceSettingsProperty
	_jsii_.Get(
		j,
		"ancillarySourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AribSourceSettings() TfChannel_AribSourceSettingsPropertyOutputReference {
	var returns TfChannel_AribSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aribSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AribSourceSettingsInput() *TfChannel_AribSourceSettingsProperty {
	var returns *TfChannel_AribSourceSettingsProperty
	_jsii_.Get(
		j,
		"aribSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) DvbSubSourceSettings() TfChannel_DvbSubSourceSettingsPropertyOutputReference {
	var returns TfChannel_DvbSubSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSubSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) DvbSubSourceSettingsInput() *TfChannel_DvbSubSourceSettingsProperty {
	var returns *TfChannel_DvbSubSourceSettingsProperty
	_jsii_.Get(
		j,
		"dvbSubSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) EmbeddedSourceSettings() TfChannel_EmbeddedSourceSettingsPropertyOutputReference {
	var returns TfChannel_EmbeddedSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) EmbeddedSourceSettingsInput() *TfChannel_EmbeddedSourceSettingsProperty {
	var returns *TfChannel_EmbeddedSourceSettingsProperty
	_jsii_.Get(
		j,
		"embeddedSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InternalValue() *TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty {
	var returns *TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte20SourceSettings() TfChannel_Scte20SourceSettingsPropertyOutputReference {
	var returns TfChannel_Scte20SourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte20SourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte20SourceSettingsInput() *TfChannel_Scte20SourceSettingsProperty {
	var returns *TfChannel_Scte20SourceSettingsProperty
	_jsii_.Get(
		j,
		"scte20SourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte27SourceSettings() TfChannel_Scte27SourceSettingsPropertyOutputReference {
	var returns TfChannel_Scte27SourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte27SourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte27SourceSettingsInput() *TfChannel_Scte27SourceSettingsProperty {
	var returns *TfChannel_Scte27SourceSettingsProperty
	_jsii_.Get(
		j,
		"scte27SourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TeletextSourceSettings() TfChannel_TeletextSourceSettingsPropertyOutputReference {
	var returns TfChannel_TeletextSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"teletextSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TeletextSourceSettingsInput() *TfChannel_TeletextSourceSettingsProperty {
	var returns *TfChannel_TeletextSourceSettingsProperty
	_jsii_.Get(
		j,
		"teletextSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference_Override(t TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutAncillarySourceSettings(value *TfChannel_AncillarySourceSettingsProperty) {
	if err := t.validatePutAncillarySourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAncillarySourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutAribSourceSettings(value *TfChannel_AribSourceSettingsProperty) {
	if err := t.validatePutAribSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAribSourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutDvbSubSourceSettings(value *TfChannel_DvbSubSourceSettingsProperty) {
	if err := t.validatePutDvbSubSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDvbSubSourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutEmbeddedSourceSettings(value *TfChannel_EmbeddedSourceSettingsProperty) {
	if err := t.validatePutEmbeddedSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmbeddedSourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutScte20SourceSettings(value *TfChannel_Scte20SourceSettingsProperty) {
	if err := t.validatePutScte20SourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScte20SourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutScte27SourceSettings(value *TfChannel_Scte27SourceSettingsProperty) {
	if err := t.validatePutScte27SourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScte27SourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutTeletextSourceSettings(value *TfChannel_TeletextSourceSettingsProperty) {
	if err := t.validatePutTeletextSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTeletextSourceSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetAncillarySourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAncillarySourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetAribSourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAribSourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetDvbSubSourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbSubSourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetEmbeddedSourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmbeddedSourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetScte20SourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetScte20SourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetScte27SourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetScte27SourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetTeletextSourceSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTeletextSourceSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

