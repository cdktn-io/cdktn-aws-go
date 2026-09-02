package awslexmodelbuilding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_SlotPropertyOutputReference interface {
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	ResponseCard() *string
	// Experimental.
	SetResponseCard(val *string)
	// Experimental.
	ResponseCardInput() *string
	// Experimental.
	SampleUtterances() *[]*string
	// Experimental.
	SetSampleUtterances(val *[]*string)
	// Experimental.
	SampleUtterancesInput() *[]*string
	// Experimental.
	SlotConstraint() *string
	// Experimental.
	SetSlotConstraint(val *string)
	// Experimental.
	SlotConstraintInput() *string
	// Experimental.
	SlotType() *string
	// Experimental.
	SetSlotType(val *string)
	// Experimental.
	SlotTypeInput() *string
	// Experimental.
	SlotTypeVersion() *string
	// Experimental.
	SetSlotTypeVersion(val *string)
	// Experimental.
	SlotTypeVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ValueElicitationPrompt() TfIntent_ValueElicitationPromptPropertyOutputReference
	// Experimental.
	ValueElicitationPromptInput() *TfIntent_ValueElicitationPromptProperty
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
	PutValueElicitationPrompt(value *TfIntent_ValueElicitationPromptProperty)
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetPriority()
	// Experimental.
	ResetResponseCard()
	// Experimental.
	ResetSampleUtterances()
	// Experimental.
	ResetSlotTypeVersion()
	// Experimental.
	ResetValueElicitationPrompt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_SlotPropertyOutputReference
type jsiiProxy_TfIntent_SlotPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SampleUtterances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SampleUtterancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) SlotTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ValueElicitationPrompt() TfIntent_ValueElicitationPromptPropertyOutputReference {
	var returns TfIntent_ValueElicitationPromptPropertyOutputReference
	_jsii_.Get(
		j,
		"valueElicitationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference) ValueElicitationPromptInput() *TfIntent_ValueElicitationPromptProperty {
	var returns *TfIntent_ValueElicitationPromptProperty
	_jsii_.Get(
		j,
		"valueElicitationPromptInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_SlotPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_SlotPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_SlotPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_SlotPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfIntent.SlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_SlotPropertyOutputReference_Override(t TfIntent_SlotPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfIntent.SlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetResponseCard(val *string) {
	if err := j.validateSetResponseCardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetSampleUtterances(val *[]*string) {
	if err := j.validateSetSampleUtterancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleUtterances",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetSlotConstraint(val *string) {
	if err := j.validateSetSlotConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotConstraint",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetSlotType(val *string) {
	if err := j.validateSetSlotTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotType",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetSlotTypeVersion(val *string) {
	if err := j.validateSetSlotTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotTypeVersion",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_SlotPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) PutValueElicitationPrompt(value *TfIntent_ValueElicitationPromptProperty) {
	if err := t.validatePutValueElicitationPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putValueElicitationPrompt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		t,
		"resetPriority",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		t,
		"resetResponseCard",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetSampleUtterances() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleUtterances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetSlotTypeVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetSlotTypeVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ResetValueElicitationPrompt() {
	_jsii_.InvokeVoid(
		t,
		"resetValueElicitationPrompt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_SlotPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

