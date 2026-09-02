package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_DvbSubDestinationSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Alignment() *string
	// Experimental.
	SetAlignment(val *string)
	// Experimental.
	AlignmentInput() *string
	// Experimental.
	BackgroundColor() *string
	// Experimental.
	SetBackgroundColor(val *string)
	// Experimental.
	BackgroundColorInput() *string
	// Experimental.
	BackgroundOpacity() *float64
	// Experimental.
	SetBackgroundOpacity(val *float64)
	// Experimental.
	BackgroundOpacityInput() *float64
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
	Font() TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontPropertyOutputReference
	// Experimental.
	FontColor() *string
	// Experimental.
	SetFontColor(val *string)
	// Experimental.
	FontColorInput() *string
	// Experimental.
	FontInput() *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty
	// Experimental.
	FontOpacity() *float64
	// Experimental.
	SetFontOpacity(val *float64)
	// Experimental.
	FontOpacityInput() *float64
	// Experimental.
	FontResolution() *float64
	// Experimental.
	SetFontResolution(val *float64)
	// Experimental.
	FontResolutionInput() *float64
	// Experimental.
	FontSize() *string
	// Experimental.
	SetFontSize(val *string)
	// Experimental.
	FontSizeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_DvbSubDestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_DvbSubDestinationSettingsProperty)
	// Experimental.
	OutlineColor() *string
	// Experimental.
	SetOutlineColor(val *string)
	// Experimental.
	OutlineColorInput() *string
	// Experimental.
	OutlineSize() *float64
	// Experimental.
	SetOutlineSize(val *float64)
	// Experimental.
	OutlineSizeInput() *float64
	// Experimental.
	ShadowColor() *string
	// Experimental.
	SetShadowColor(val *string)
	// Experimental.
	ShadowColorInput() *string
	// Experimental.
	ShadowOpacity() *float64
	// Experimental.
	SetShadowOpacity(val *float64)
	// Experimental.
	ShadowOpacityInput() *float64
	// Experimental.
	ShadowXOffset() *float64
	// Experimental.
	SetShadowXOffset(val *float64)
	// Experimental.
	ShadowXOffsetInput() *float64
	// Experimental.
	ShadowYOffset() *float64
	// Experimental.
	SetShadowYOffset(val *float64)
	// Experimental.
	ShadowYOffsetInput() *float64
	// Experimental.
	TeletextGridControl() *string
	// Experimental.
	SetTeletextGridControl(val *string)
	// Experimental.
	TeletextGridControlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	XPosition() *float64
	// Experimental.
	SetXPosition(val *float64)
	// Experimental.
	XPositionInput() *float64
	// Experimental.
	YPosition() *float64
	// Experimental.
	SetYPosition(val *float64)
	// Experimental.
	YPositionInput() *float64
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
	PutFont(value *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty)
	// Experimental.
	ResetAlignment()
	// Experimental.
	ResetBackgroundColor()
	// Experimental.
	ResetBackgroundOpacity()
	// Experimental.
	ResetFont()
	// Experimental.
	ResetFontColor()
	// Experimental.
	ResetFontOpacity()
	// Experimental.
	ResetFontResolution()
	// Experimental.
	ResetFontSize()
	// Experimental.
	ResetOutlineColor()
	// Experimental.
	ResetOutlineSize()
	// Experimental.
	ResetShadowColor()
	// Experimental.
	ResetShadowOpacity()
	// Experimental.
	ResetShadowXOffset()
	// Experimental.
	ResetShadowYOffset()
	// Experimental.
	ResetTeletextGridControl()
	// Experimental.
	ResetXPosition()
	// Experimental.
	ResetYPosition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_DvbSubDestinationSettingsPropertyOutputReference
type jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) Alignment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) AlignmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) BackgroundOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) BackgroundOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) Font() TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontPropertyOutputReference {
	var returns TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontPropertyOutputReference
	_jsii_.Get(
		j,
		"font",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontInput() *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty {
	var returns *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty
	_jsii_.Get(
		j,
		"fontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontResolution() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontResolutionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) InternalValue() *TfChannel_DvbSubDestinationSettingsProperty {
	var returns *TfChannel_DvbSubDestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) OutlineColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) OutlineColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) OutlineSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) OutlineSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowXOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowXOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowYOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ShadowYOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) TeletextGridControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) TeletextGridControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) XPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) XPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) YPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) YPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPositionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_DvbSubDestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_DvbSubDestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_DvbSubDestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.DvbSubDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_DvbSubDestinationSettingsPropertyOutputReference_Override(t TfChannel_DvbSubDestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.DvbSubDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetAlignment(val *string) {
	if err := j.validateSetAlignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alignment",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetBackgroundOpacity(val *float64) {
	if err := j.validateSetBackgroundOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundOpacity",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetFontColor(val *string) {
	if err := j.validateSetFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontColor",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetFontOpacity(val *float64) {
	if err := j.validateSetFontOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontOpacity",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetFontResolution(val *float64) {
	if err := j.validateSetFontResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontResolution",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_DvbSubDestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetOutlineColor(val *string) {
	if err := j.validateSetOutlineColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineColor",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetOutlineSize(val *float64) {
	if err := j.validateSetOutlineSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineSize",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetShadowColor(val *string) {
	if err := j.validateSetShadowColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowColor",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetShadowOpacity(val *float64) {
	if err := j.validateSetShadowOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowOpacity",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetShadowXOffset(val *float64) {
	if err := j.validateSetShadowXOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowXOffset",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetShadowYOffset(val *float64) {
	if err := j.validateSetShadowYOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowYOffset",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetTeletextGridControl(val *string) {
	if err := j.validateSetTeletextGridControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teletextGridControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetXPosition(val *float64) {
	if err := j.validateSetXPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xPosition",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference)SetYPosition(val *float64) {
	if err := j.validateSetYPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yPosition",
		val,
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) PutFont(value *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty) {
	if err := t.validatePutFontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFont",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetAlignment() {
	_jsii_.InvokeVoid(
		t,
		"resetAlignment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		t,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetBackgroundOpacity() {
	_jsii_.InvokeVoid(
		t,
		"resetBackgroundOpacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetFont() {
	_jsii_.InvokeVoid(
		t,
		"resetFont",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		t,
		"resetFontColor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetFontOpacity() {
	_jsii_.InvokeVoid(
		t,
		"resetFontOpacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetFontResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetFontResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		t,
		"resetFontSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetOutlineColor() {
	_jsii_.InvokeVoid(
		t,
		"resetOutlineColor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetOutlineSize() {
	_jsii_.InvokeVoid(
		t,
		"resetOutlineSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetShadowColor() {
	_jsii_.InvokeVoid(
		t,
		"resetShadowColor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetShadowOpacity() {
	_jsii_.InvokeVoid(
		t,
		"resetShadowOpacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetShadowXOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetShadowXOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetShadowYOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetShadowYOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetTeletextGridControl() {
	_jsii_.InvokeVoid(
		t,
		"resetTeletextGridControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetXPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetXPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ResetYPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetYPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_DvbSubDestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

