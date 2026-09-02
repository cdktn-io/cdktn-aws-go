package awselastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPreset_ThumbnailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AspectRatio() *string
	// Experimental.
	SetAspectRatio(val *string)
	// Experimental.
	AspectRatioInput() *string
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
	Format() *string
	// Experimental.
	SetFormat(val *string)
	// Experimental.
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPreset_ThumbnailsProperty
	// Experimental.
	SetInternalValue(val *TfPreset_ThumbnailsProperty)
	// Experimental.
	Interval() *string
	// Experimental.
	SetInterval(val *string)
	// Experimental.
	IntervalInput() *string
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
	ResetFormat()
	// Experimental.
	ResetInterval()
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

// The jsii proxy struct for TfPreset_ThumbnailsPropertyOutputReference
type jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) AspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) AspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) InternalValue() *TfPreset_ThumbnailsProperty {
	var returns *TfPreset_ThumbnailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) PaddingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) PaddingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) Resolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPreset_ThumbnailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPreset_ThumbnailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPreset_ThumbnailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPreset.ThumbnailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPreset_ThumbnailsPropertyOutputReference_Override(t TfPreset_ThumbnailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPreset.ThumbnailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetAspectRatio(val *string) {
	if err := j.validateSetAspectRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aspectRatio",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetInternalValue(val *TfPreset_ThumbnailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetMaxHeight(val *string) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetMaxWidth(val *string) {
	if err := j.validateSetMaxWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetPaddingPolicy(val *string) {
	if err := j.validateSetPaddingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paddingPolicy",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetResolution(val *string) {
	if err := j.validateSetResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolution",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetSizingPolicy(val *string) {
	if err := j.validateSetSizingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		t,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetPaddingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPaddingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPreset_ThumbnailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

