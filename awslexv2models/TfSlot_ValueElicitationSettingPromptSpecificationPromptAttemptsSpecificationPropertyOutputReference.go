package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedInputTypes() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypesPropertyList
	// Experimental.
	AllowedInputTypesInput() interface{}
	// Experimental.
	AllowInterrupt() interface{}
	// Experimental.
	SetAllowInterrupt(val interface{})
	// Experimental.
	AllowInterruptInput() interface{}
	// Experimental.
	AudioAndDtmfInputSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationPropertyList
	// Experimental.
	AudioAndDtmfInputSpecificationInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MapBlockKey() *string
	// Experimental.
	SetMapBlockKey(val *string)
	// Experimental.
	MapBlockKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextInputSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecificationPropertyList
	// Experimental.
	TextInputSpecificationInput() interface{}
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
	PutAllowedInputTypes(value interface{})
	// Experimental.
	PutAudioAndDtmfInputSpecification(value interface{})
	// Experimental.
	PutTextInputSpecification(value interface{})
	// Experimental.
	ResetAllowedInputTypes()
	// Experimental.
	ResetAllowInterrupt()
	// Experimental.
	ResetAudioAndDtmfInputSpecification()
	// Experimental.
	ResetTextInputSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference
type jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AllowedInputTypes() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypesPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypesPropertyList
	_jsii_.Get(
		j,
		"allowedInputTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AllowedInputTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInputTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AllowInterrupt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterrupt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AllowInterruptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterruptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AudioAndDtmfInputSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationPropertyList
	_jsii_.Get(
		j,
		"audioAndDtmfInputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) AudioAndDtmfInputSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioAndDtmfInputSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) MapBlockKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) MapBlockKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) TextInputSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecificationPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecificationPropertyList
	_jsii_.Get(
		j,
		"textInputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) TextInputSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textInputSpecificationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference_Override(t TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetAllowInterrupt(val interface{}) {
	if err := j.validateSetAllowInterruptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInterrupt",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetMapBlockKey(val *string) {
	if err := j.validateSetMapBlockKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBlockKey",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) PutAllowedInputTypes(value interface{}) {
	if err := t.validatePutAllowedInputTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllowedInputTypes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) PutAudioAndDtmfInputSpecification(value interface{}) {
	if err := t.validatePutAudioAndDtmfInputSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioAndDtmfInputSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) PutTextInputSpecification(value interface{}) {
	if err := t.validatePutTextInputSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTextInputSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ResetAllowedInputTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedInputTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ResetAllowInterrupt() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowInterrupt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ResetAudioAndDtmfInputSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioAndDtmfInputSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ResetTextInputSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetTextInputSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

