package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference interface {
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
	Shape() *string
	// Experimental.
	SetShape(val *string)
	// Experimental.
	ShapeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Value() TfIntent_ConfirmationSettingFailureNextStepIntentSlotValuePropertyList
	// Experimental.
	ValueInput() interface{}
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
	PutValue(value interface{})
	// Experimental.
	ResetShape()
	// Experimental.
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference
type jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) MapBlockKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) MapBlockKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) Value() TfIntent_ConfirmationSettingFailureNextStepIntentSlotValuePropertyList {
	var returns TfIntent_ConfirmationSettingFailureNextStepIntentSlotValuePropertyList
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference_Override(t TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetMapBlockKey(val *string) {
	if err := j.validateSetMapBlockKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBlockKey",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) PutValue(value interface{}) {
	if err := t.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ResetShape() {
	_jsii_.InvokeVoid(
		t,
		"resetShape",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		t,
		"resetValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingFailureNextStepIntentSlotPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

