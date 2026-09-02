package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Button() TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardButtonPropertyList
	// Experimental.
	ButtonInput() interface{}
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
	ImageUrl() *string
	// Experimental.
	SetImageUrl(val *string)
	// Experimental.
	ImageUrlInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Subtitle() *string
	// Experimental.
	SetSubtitle(val *string)
	// Experimental.
	SubtitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() *string
	// Experimental.
	SetTitle(val *string)
	// Experimental.
	TitleInput() *string
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
	PutButton(value interface{})
	// Experimental.
	ResetButton()
	// Experimental.
	ResetImageUrl()
	// Experimental.
	ResetSubtitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference
type jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) Button() TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardButtonPropertyList {
	var returns TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardButtonPropertyList
	_jsii_.Get(
		j,
		"button",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buttonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) SubtitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference_Override(t TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetImageUrl(val *string) {
	if err := j.validateSetImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetSubtitle(val *string) {
	if err := j.validateSetSubtitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtitle",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) PutButton(value interface{}) {
	if err := t.validatePutButtonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putButton",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ResetButton() {
	_jsii_.InvokeVoid(
		t,
		"resetButton",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ResetSubtitle() {
	_jsii_.InvokeVoid(
		t,
		"resetSubtitle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingDeclinationConditionalConditionalBranchResponseMessageGroupVariationImageResponseCardPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

