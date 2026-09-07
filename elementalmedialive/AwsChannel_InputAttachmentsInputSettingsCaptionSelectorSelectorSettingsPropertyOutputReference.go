package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AncillarySourceSettings() AwsChannel_AncillarySourceSettingsPropertyOutputReference
	// Experimental.
	AncillarySourceSettingsInput() *AwsChannel_AncillarySourceSettingsProperty
	// Experimental.
	AribSourceSettings() AwsChannel_AribSourceSettingsPropertyOutputReference
	// Experimental.
	AribSourceSettingsInput() *AwsChannel_AribSourceSettingsProperty
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
	DvbSubSourceSettings() AwsChannel_DvbSubSourceSettingsPropertyOutputReference
	// Experimental.
	DvbSubSourceSettingsInput() *AwsChannel_DvbSubSourceSettingsProperty
	// Experimental.
	EmbeddedSourceSettings() AwsChannel_EmbeddedSourceSettingsPropertyOutputReference
	// Experimental.
	EmbeddedSourceSettingsInput() *AwsChannel_EmbeddedSourceSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty)
	// Experimental.
	Scte20SourceSettings() AwsChannel_Scte20SourceSettingsPropertyOutputReference
	// Experimental.
	Scte20SourceSettingsInput() *AwsChannel_Scte20SourceSettingsProperty
	// Experimental.
	Scte27SourceSettings() AwsChannel_Scte27SourceSettingsPropertyOutputReference
	// Experimental.
	Scte27SourceSettingsInput() *AwsChannel_Scte27SourceSettingsProperty
	// Experimental.
	TeletextSourceSettings() AwsChannel_TeletextSourceSettingsPropertyOutputReference
	// Experimental.
	TeletextSourceSettingsInput() *AwsChannel_TeletextSourceSettingsProperty
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
	PutAncillarySourceSettings(value *AwsChannel_AncillarySourceSettingsProperty)
	// Experimental.
	PutAribSourceSettings(value *AwsChannel_AribSourceSettingsProperty)
	// Experimental.
	PutDvbSubSourceSettings(value *AwsChannel_DvbSubSourceSettingsProperty)
	// Experimental.
	PutEmbeddedSourceSettings(value *AwsChannel_EmbeddedSourceSettingsProperty)
	// Experimental.
	PutScte20SourceSettings(value *AwsChannel_Scte20SourceSettingsProperty)
	// Experimental.
	PutScte27SourceSettings(value *AwsChannel_Scte27SourceSettingsProperty)
	// Experimental.
	PutTeletextSourceSettings(value *AwsChannel_TeletextSourceSettingsProperty)
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

// The jsii proxy struct for AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AncillarySourceSettings() AwsChannel_AncillarySourceSettingsPropertyOutputReference {
	var returns AwsChannel_AncillarySourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ancillarySourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AncillarySourceSettingsInput() *AwsChannel_AncillarySourceSettingsProperty {
	var returns *AwsChannel_AncillarySourceSettingsProperty
	_jsii_.Get(
		j,
		"ancillarySourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AribSourceSettings() AwsChannel_AribSourceSettingsPropertyOutputReference {
	var returns AwsChannel_AribSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aribSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) AribSourceSettingsInput() *AwsChannel_AribSourceSettingsProperty {
	var returns *AwsChannel_AribSourceSettingsProperty
	_jsii_.Get(
		j,
		"aribSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) DvbSubSourceSettings() AwsChannel_DvbSubSourceSettingsPropertyOutputReference {
	var returns AwsChannel_DvbSubSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSubSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) DvbSubSourceSettingsInput() *AwsChannel_DvbSubSourceSettingsProperty {
	var returns *AwsChannel_DvbSubSourceSettingsProperty
	_jsii_.Get(
		j,
		"dvbSubSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) EmbeddedSourceSettings() AwsChannel_EmbeddedSourceSettingsPropertyOutputReference {
	var returns AwsChannel_EmbeddedSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) EmbeddedSourceSettingsInput() *AwsChannel_EmbeddedSourceSettingsProperty {
	var returns *AwsChannel_EmbeddedSourceSettingsProperty
	_jsii_.Get(
		j,
		"embeddedSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InternalValue() *AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty {
	var returns *AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte20SourceSettings() AwsChannel_Scte20SourceSettingsPropertyOutputReference {
	var returns AwsChannel_Scte20SourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte20SourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte20SourceSettingsInput() *AwsChannel_Scte20SourceSettingsProperty {
	var returns *AwsChannel_Scte20SourceSettingsProperty
	_jsii_.Get(
		j,
		"scte20SourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte27SourceSettings() AwsChannel_Scte27SourceSettingsPropertyOutputReference {
	var returns AwsChannel_Scte27SourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte27SourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Scte27SourceSettingsInput() *AwsChannel_Scte27SourceSettingsProperty {
	var returns *AwsChannel_Scte27SourceSettingsProperty
	_jsii_.Get(
		j,
		"scte27SourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TeletextSourceSettings() AwsChannel_TeletextSourceSettingsPropertyOutputReference {
	var returns AwsChannel_TeletextSourceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"teletextSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TeletextSourceSettingsInput() *AwsChannel_TeletextSourceSettingsProperty {
	var returns *AwsChannel_TeletextSourceSettingsProperty
	_jsii_.Get(
		j,
		"teletextSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference_Override(a AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutAncillarySourceSettings(value *AwsChannel_AncillarySourceSettingsProperty) {
	if err := a.validatePutAncillarySourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAncillarySourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutAribSourceSettings(value *AwsChannel_AribSourceSettingsProperty) {
	if err := a.validatePutAribSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAribSourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutDvbSubSourceSettings(value *AwsChannel_DvbSubSourceSettingsProperty) {
	if err := a.validatePutDvbSubSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDvbSubSourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutEmbeddedSourceSettings(value *AwsChannel_EmbeddedSourceSettingsProperty) {
	if err := a.validatePutEmbeddedSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddedSourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutScte20SourceSettings(value *AwsChannel_Scte20SourceSettingsProperty) {
	if err := a.validatePutScte20SourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte20SourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutScte27SourceSettings(value *AwsChannel_Scte27SourceSettingsProperty) {
	if err := a.validatePutScte27SourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte27SourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) PutTeletextSourceSettings(value *AwsChannel_TeletextSourceSettingsProperty) {
	if err := a.validatePutTeletextSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTeletextSourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetAncillarySourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAncillarySourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetAribSourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAribSourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetDvbSubSourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDvbSubSourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetEmbeddedSourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddedSourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetScte20SourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte20SourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetScte27SourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte27SourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ResetTeletextSourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTeletextSourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

