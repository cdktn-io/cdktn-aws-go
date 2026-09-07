package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Button() AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList
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

// The jsii proxy struct for AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference
type jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) Button() AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList {
	var returns AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList
	_jsii_.Get(
		j,
		"button",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buttonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) SubtitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference_Override(a AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetImageUrl(val *string) {
	if err := j.validateSetImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetSubtitle(val *string) {
	if err := j.validateSetSubtitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtitle",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) PutButton(value interface{}) {
	if err := a.validatePutButtonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putButton",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ResetButton() {
	_jsii_.InvokeVoid(
		a,
		"resetButton",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ResetSubtitle() {
	_jsii_.InvokeVoid(
		a,
		"resetSubtitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_ClosingSettingConditionalConditionalBranchResponseMessageGroupMessageImageResponseCardPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

