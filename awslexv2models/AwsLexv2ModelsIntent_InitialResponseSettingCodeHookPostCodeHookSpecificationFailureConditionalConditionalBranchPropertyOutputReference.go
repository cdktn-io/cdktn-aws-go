package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference interface {
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
	// Experimental.
	Condition() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionPropertyList
	// Experimental.
	ConditionInput() interface{}
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
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepPropertyList
	// Experimental.
	NextStepInput() interface{}
	// Experimental.
	Response() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponsePropertyList
	// Experimental.
	ResponseInput() interface{}
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
	PutCondition(value interface{})
	// Experimental.
	PutNextStep(value interface{})
	// Experimental.
	PutResponse(value interface{})
	// Experimental.
	ResetCondition()
	// Experimental.
	ResetNextStep()
	// Experimental.
	ResetResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference
type jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) Condition() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionPropertyList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) NextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepPropertyList
	_jsii_.Get(
		j,
		"nextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) NextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) Response() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponsePropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponsePropertyList
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference_Override(a AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) PutCondition(value interface{}) {
	if err := a.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) PutNextStep(value interface{}) {
	if err := a.validatePutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) PutResponse(value interface{}) {
	if err := a.validatePutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ResetNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

