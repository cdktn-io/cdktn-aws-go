package awselastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPreset_VideoWatermarksPropertyOutputReference interface {
	cdktn.ComplexObject
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
	HorizontalAlign() *string
	// Experimental.
	SetHorizontalAlign(val *string)
	// Experimental.
	HorizontalAlignInput() *string
	// Experimental.
	HorizontalOffset() *string
	// Experimental.
	SetHorizontalOffset(val *string)
	// Experimental.
	HorizontalOffsetInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	Opacity() *string
	// Experimental.
	SetOpacity(val *string)
	// Experimental.
	OpacityInput() *string
	// Experimental.
	SizingPolicy() *string
	// Experimental.
	SetSizingPolicy(val *string)
	// Experimental.
	SizingPolicyInput() *string
	// Experimental.
	Target() *string
	// Experimental.
	SetTarget(val *string)
	// Experimental.
	TargetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VerticalAlign() *string
	// Experimental.
	SetVerticalAlign(val *string)
	// Experimental.
	VerticalAlignInput() *string
	// Experimental.
	VerticalOffset() *string
	// Experimental.
	SetVerticalOffset(val *string)
	// Experimental.
	VerticalOffsetInput() *string
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
	ResetHorizontalAlign()
	// Experimental.
	ResetHorizontalOffset()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMaxHeight()
	// Experimental.
	ResetMaxWidth()
	// Experimental.
	ResetOpacity()
	// Experimental.
	ResetSizingPolicy()
	// Experimental.
	ResetTarget()
	// Experimental.
	ResetVerticalAlign()
	// Experimental.
	ResetVerticalOffset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPreset_VideoWatermarksPropertyOutputReference
type jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) HorizontalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) HorizontalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) HorizontalOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) HorizontalOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) Opacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) OpacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) VerticalOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) VerticalOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalOffsetInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPreset_VideoWatermarksPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPreset_VideoWatermarksPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPreset_VideoWatermarksPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPreset.VideoWatermarksPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPreset_VideoWatermarksPropertyOutputReference_Override(t TfPreset_VideoWatermarksPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.TfPreset.VideoWatermarksPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetHorizontalAlign(val *string) {
	if err := j.validateSetHorizontalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalAlign",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetHorizontalOffset(val *string) {
	if err := j.validateSetHorizontalOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalOffset",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetMaxHeight(val *string) {
	if err := j.validateSetMaxHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetMaxWidth(val *string) {
	if err := j.validateSetMaxWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetOpacity(val *string) {
	if err := j.validateSetOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opacity",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetSizingPolicy(val *string) {
	if err := j.validateSetSizingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (j *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference)SetVerticalOffset(val *string) {
	if err := j.validateSetVerticalOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalOffset",
		val,
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetHorizontalAlign() {
	_jsii_.InvokeVoid(
		t,
		"resetHorizontalAlign",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetHorizontalOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetHorizontalOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetOpacity() {
	_jsii_.InvokeVoid(
		t,
		"resetOpacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		t,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ResetVerticalOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetVerticalOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPreset_VideoWatermarksPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

