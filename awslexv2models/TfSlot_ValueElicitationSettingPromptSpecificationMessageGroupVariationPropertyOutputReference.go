package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference interface {
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
	CustomPayload() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationCustomPayloadPropertyList
	// Experimental.
	CustomPayloadInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageResponseCard() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationImageResponseCardPropertyList
	// Experimental.
	ImageResponseCardInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PlainTextMessage() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPlainTextMessagePropertyList
	// Experimental.
	PlainTextMessageInput() interface{}
	// Experimental.
	SsmlMessage() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationSsmlMessagePropertyList
	// Experimental.
	SsmlMessageInput() interface{}
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
	PutCustomPayload(value interface{})
	// Experimental.
	PutImageResponseCard(value interface{})
	// Experimental.
	PutPlainTextMessage(value interface{})
	// Experimental.
	PutSsmlMessage(value interface{})
	// Experimental.
	ResetCustomPayload()
	// Experimental.
	ResetImageResponseCard()
	// Experimental.
	ResetPlainTextMessage()
	// Experimental.
	ResetSsmlMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference
type jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) CustomPayload() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationCustomPayloadPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationCustomPayloadPropertyList
	_jsii_.Get(
		j,
		"customPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) CustomPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ImageResponseCard() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationImageResponseCardPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationImageResponseCardPropertyList
	_jsii_.Get(
		j,
		"imageResponseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ImageResponseCardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageResponseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PlainTextMessage() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPlainTextMessagePropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPlainTextMessagePropertyList
	_jsii_.Get(
		j,
		"plainTextMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PlainTextMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) SsmlMessage() TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationSsmlMessagePropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationSsmlMessagePropertyList
	_jsii_.Get(
		j,
		"ssmlMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) SsmlMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmlMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference_Override(t TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PutCustomPayload(value interface{}) {
	if err := t.validatePutCustomPayloadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomPayload",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PutImageResponseCard(value interface{}) {
	if err := t.validatePutImageResponseCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageResponseCard",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PutPlainTextMessage(value interface{}) {
	if err := t.validatePutPlainTextMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlainTextMessage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) PutSsmlMessage(value interface{}) {
	if err := t.validatePutSsmlMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSsmlMessage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ResetCustomPayload() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPayload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ResetImageResponseCard() {
	_jsii_.InvokeVoid(
		t,
		"resetImageResponseCard",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ResetPlainTextMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetPlainTextMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ResetSsmlMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetSsmlMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationMessageGroupVariationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

