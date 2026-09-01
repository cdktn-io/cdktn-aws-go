package awslexmodelbuilding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexIntent_SlotPropertyOutputReference interface {
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
	ValueElicitationPrompt() AwsLexIntent_ValueElicitationPromptPropertyOutputReference
	// Experimental.
	ValueElicitationPromptInput() *AwsLexIntent_ValueElicitationPromptProperty
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
	PutValueElicitationPrompt(value *AwsLexIntent_ValueElicitationPromptProperty)
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

// The jsii proxy struct for AwsLexIntent_SlotPropertyOutputReference
type jsiiProxy_AwsLexIntent_SlotPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResponseCard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResponseCardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SampleUtterances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SampleUtterancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) SlotTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ValueElicitationPrompt() AwsLexIntent_ValueElicitationPromptPropertyOutputReference {
	var returns AwsLexIntent_ValueElicitationPromptPropertyOutputReference
	_jsii_.Get(
		j,
		"valueElicitationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ValueElicitationPromptInput() *AwsLexIntent_ValueElicitationPromptProperty {
	var returns *AwsLexIntent_ValueElicitationPromptProperty
	_jsii_.Get(
		j,
		"valueElicitationPromptInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexIntent_SlotPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLexIntent_SlotPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLexIntent_SlotPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexIntent_SlotPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.AwsLexIntent.SlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexIntent_SlotPropertyOutputReference_Override(a AwsLexIntent_SlotPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.AwsLexIntent.SlotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetResponseCard(val *string) {
	if err := j.validateSetResponseCardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCard",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetSampleUtterances(val *[]*string) {
	if err := j.validateSetSampleUtterancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleUtterances",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetSlotConstraint(val *string) {
	if err := j.validateSetSlotConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotConstraint",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetSlotType(val *string) {
	if err := j.validateSetSlotTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotType",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetSlotTypeVersion(val *string) {
	if err := j.validateSetSlotTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotTypeVersion",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) PutValueElicitationPrompt(value *AwsLexIntent_ValueElicitationPromptProperty) {
	if err := a.validatePutValueElicitationPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putValueElicitationPrompt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetResponseCard() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseCard",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetSampleUtterances() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleUtterances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetSlotTypeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetSlotTypeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ResetValueElicitationPrompt() {
	_jsii_.InvokeVoid(
		a,
		"resetValueElicitationPrompt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexIntent_SlotPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

