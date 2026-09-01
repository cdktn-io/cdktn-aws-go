package awselastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsElastictranscoderPreset_VideoPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AspectRatio() *string
	// Experimental.
	SetAspectRatio(val *string)
	// Experimental.
	AspectRatioInput() *string
	// Experimental.
	BitRate() *string
	// Experimental.
	SetBitRate(val *string)
	// Experimental.
	BitRateInput() *string
	// Experimental.
	Codec() *string
	// Experimental.
	SetCodec(val *string)
	// Experimental.
	CodecInput() *string
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
	DisplayAspectRatio() *string
	// Experimental.
	SetDisplayAspectRatio(val *string)
	// Experimental.
	DisplayAspectRatioInput() *string
	// Experimental.
	FixedGop() *string
	// Experimental.
	SetFixedGop(val *string)
	// Experimental.
	FixedGopInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameRate() *string
	// Experimental.
	SetFrameRate(val *string)
	// Experimental.
	FrameRateInput() *string
	// Experimental.
	InternalValue() *AwsElastictranscoderPreset_VideoProperty
	// Experimental.
	SetInternalValue(val *AwsElastictranscoderPreset_VideoProperty)
	// Experimental.
	KeyframesMaxDist() *string
	// Experimental.
	SetKeyframesMaxDist(val *string)
	// Experimental.
	KeyframesMaxDistInput() *string
	// Experimental.
	MaxFrameRate() *string
	// Experimental.
	SetMaxFrameRate(val *string)
	// Experimental.
	MaxFrameRateInput() *string
	// Experimental.
	MaxHeight() *string
	// Experimental.
	SetMaxHeight(val *string)
	// Experimental.
	MaxHeightInput() *string
	// Experimental.
	MaxWidth() *string
	// Experimental.
	SetMaxWidth(val *string)
	// Experimental.
	MaxWidthInput() *string
	// Experimental.
	PaddingPolicy() *string
	// Experimental.
	SetPaddingPolicy(val *string)
	// Experimental.
	PaddingPolicyInput() *string
	// Experimental.
	Resolution() *string
	// Experimental.
	SetResolution(val *string)
	// Experimental.
	ResolutionInput() *string
	// Experimental.
	SizingPolicy() *string
	// Experimental.
	SetSizingPolicy(val *string)
	// Experimental.
	SizingPolicyInput() *string
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
	ResetAspectRatio()
	// Experimental.
	ResetBitRate()
	// Experimental.
	ResetCodec()
	// Experimental.
	ResetDisplayAspectRatio()
	// Experimental.
	ResetFixedGop()
	// Experimental.
	ResetFrameRate()
	// Experimental.
	ResetKeyframesMaxDist()
	// Experimental.
	ResetMaxFrameRate()
	// Experimental.
	ResetMaxHeight()
	// Experimental.
	ResetMaxWidth()
	// Experimental.
	ResetPaddingPolicy()
	// Experimental.
	ResetResolution()
	// Experimental.
	ResetSizingPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsElastictranscoderPreset_VideoPropertyOutputReference
type jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) AspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) AspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) BitRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) BitRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) DisplayAspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) DisplayAspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) FixedGop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) FixedGopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) FrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) FrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) InternalValue() *AwsElastictranscoderPreset_VideoProperty {
	var returns *AwsElastictranscoderPreset_VideoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) KeyframesMaxDist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) KeyframesMaxDistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxFrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxFrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) PaddingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) PaddingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) Resolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsElastictranscoderPreset_VideoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsElastictranscoderPreset_VideoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsElastictranscoderPreset_VideoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsElastictranscoderPreset.VideoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsElastictranscoderPreset_VideoPropertyOutputReference_Override(a AwsElastictranscoderPreset_VideoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsElastictranscoderPreset.VideoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetAspectRatio(val *string) {
	if err := j.validateSetAspectRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aspectRatio",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetBitRate(val *string) {
	if err := j.validateSetBitRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitRate",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetCodec(val *string) {
	if err := j.validateSetCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetDisplayAspectRatio(val *string) {
	if err := j.validateSetDisplayAspectRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayAspectRatio",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetFixedGop(val *string) {
	if err := j.validateSetFixedGopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedGop",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetFrameRate(val *string) {
	if err := j.validateSetFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameRate",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetInternalValue(val *AwsElastictranscoderPreset_VideoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetKeyframesMaxDist(val *string) {
	if err := j.validateSetKeyframesMaxDistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyframesMaxDist",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetMaxFrameRate(val *string) {
	if err := j.validateSetMaxFrameRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFrameRate",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetMaxHeight(val *string) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetMaxWidth(val *string) {
	if err := j.validateSetMaxWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetPaddingPolicy(val *string) {
	if err := j.validateSetPaddingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paddingPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetResolution(val *string) {
	if err := j.validateSetResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolution",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetSizingPolicy(val *string) {
	if err := j.validateSetSizingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		a,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetBitRate() {
	_jsii_.InvokeVoid(
		a,
		"resetBitRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		a,
		"resetCodec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetDisplayAspectRatio() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayAspectRatio",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetFixedGop() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedGop",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetFrameRate() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetKeyframesMaxDist() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyframesMaxDist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetMaxFrameRate() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxFrameRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetPaddingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPaddingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsElastictranscoderPreset_VideoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

