package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference interface {
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
	Font() AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontPropertyOutputReference
	// Experimental.
	FontColor() *string
	// Experimental.
	SetFontColor(val *string)
	// Experimental.
	FontColorInput() *string
	// Experimental.
	FontInput() *AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontProperty
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
	InternalValue() *AwsMedialiveChannel_BurnInDestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_BurnInDestinationSettingsProperty)
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
	PutFont(value *AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) Alignment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) AlignmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) BackgroundOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) BackgroundOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) Font() AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontPropertyOutputReference {
	var returns AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontPropertyOutputReference
	_jsii_.Get(
		j,
		"font",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontInput() *AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontProperty {
	var returns *AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontProperty
	_jsii_.Get(
		j,
		"fontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontResolution() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontResolutionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_BurnInDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_BurnInDestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) OutlineColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) OutlineColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) OutlineSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) OutlineSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowXOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowXOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowYOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ShadowYOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) TeletextGridControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) TeletextGridControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) XPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) XPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) YPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) YPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPositionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.BurnInDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.BurnInDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetAlignment(val *string) {
	if err := j.validateSetAlignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alignment",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetBackgroundOpacity(val *float64) {
	if err := j.validateSetBackgroundOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundOpacity",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetFontColor(val *string) {
	if err := j.validateSetFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontColor",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetFontOpacity(val *float64) {
	if err := j.validateSetFontOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontOpacity",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetFontResolution(val *float64) {
	if err := j.validateSetFontResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontResolution",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_BurnInDestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetOutlineColor(val *string) {
	if err := j.validateSetOutlineColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineColor",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetOutlineSize(val *float64) {
	if err := j.validateSetOutlineSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineSize",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetShadowColor(val *string) {
	if err := j.validateSetShadowColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowColor",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetShadowOpacity(val *float64) {
	if err := j.validateSetShadowOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowOpacity",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetShadowXOffset(val *float64) {
	if err := j.validateSetShadowXOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowXOffset",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetShadowYOffset(val *float64) {
	if err := j.validateSetShadowYOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowYOffset",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetTeletextGridControl(val *string) {
	if err := j.validateSetTeletextGridControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teletextGridControl",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetXPosition(val *float64) {
	if err := j.validateSetXPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xPosition",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference)SetYPosition(val *float64) {
	if err := j.validateSetYPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yPosition",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) PutFont(value *AwsMedialiveChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontProperty) {
	if err := a.validatePutFontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFont",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetAlignment() {
	_jsii_.InvokeVoid(
		a,
		"resetAlignment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		a,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetBackgroundOpacity() {
	_jsii_.InvokeVoid(
		a,
		"resetBackgroundOpacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetFont() {
	_jsii_.InvokeVoid(
		a,
		"resetFont",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		a,
		"resetFontColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetFontOpacity() {
	_jsii_.InvokeVoid(
		a,
		"resetFontOpacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetFontResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetFontResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		a,
		"resetFontSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetOutlineSize() {
	_jsii_.InvokeVoid(
		a,
		"resetOutlineSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetShadowColor() {
	_jsii_.InvokeVoid(
		a,
		"resetShadowColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetShadowOpacity() {
	_jsii_.InvokeVoid(
		a,
		"resetShadowOpacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetShadowXOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetShadowXOffset",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetShadowYOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetShadowYOffset",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetXPosition() {
	_jsii_.InvokeVoid(
		a,
		"resetXPosition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ResetYPosition() {
	_jsii_.InvokeVoid(
		a,
		"resetYPosition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

