package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrainingJob_ServerlessJobConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcceptEula() interface{}
	// Experimental.
	SetAcceptEula(val interface{})
	// Experimental.
	AcceptEulaInput() interface{}
	// Experimental.
	BaseModelArn() *string
	// Experimental.
	SetBaseModelArn(val *string)
	// Experimental.
	BaseModelArnInput() *string
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
	CustomizationTechnique() *string
	// Experimental.
	SetCustomizationTechnique(val *string)
	// Experimental.
	CustomizationTechniqueInput() *string
	// Experimental.
	EvaluationType() *string
	// Experimental.
	SetEvaluationType(val *string)
	// Experimental.
	EvaluationTypeInput() *string
	// Experimental.
	EvaluatorArn() *string
	// Experimental.
	SetEvaluatorArn(val *string)
	// Experimental.
	EvaluatorArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JobType() *string
	// Experimental.
	SetJobType(val *string)
	// Experimental.
	JobTypeInput() *string
	// Experimental.
	Peft() *string
	// Experimental.
	SetPeft(val *string)
	// Experimental.
	PeftInput() *string
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
	ResetAcceptEula()
	// Experimental.
	ResetCustomizationTechnique()
	// Experimental.
	ResetEvaluationType()
	// Experimental.
	ResetEvaluatorArn()
	// Experimental.
	ResetPeft()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTrainingJob_ServerlessJobConfigPropertyOutputReference
type jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) AcceptEula() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) AcceptEulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) BaseModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) BaseModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) CustomizationTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) CustomizationTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) EvaluationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) EvaluationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) EvaluatorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) EvaluatorArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) JobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) Peft() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) PeftInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTrainingJob_ServerlessJobConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTrainingJob_ServerlessJobConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTrainingJob_ServerlessJobConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob.ServerlessJobConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTrainingJob_ServerlessJobConfigPropertyOutputReference_Override(t TfTrainingJob_ServerlessJobConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob.ServerlessJobConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetAcceptEula(val interface{}) {
	if err := j.validateSetAcceptEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptEula",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetBaseModelArn(val *string) {
	if err := j.validateSetBaseModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseModelArn",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetCustomizationTechnique(val *string) {
	if err := j.validateSetCustomizationTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customizationTechnique",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetEvaluationType(val *string) {
	if err := j.validateSetEvaluationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationType",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetEvaluatorArn(val *string) {
	if err := j.validateSetEvaluatorArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluatorArn",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetJobType(val *string) {
	if err := j.validateSetJobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetPeft(val *string) {
	if err := j.validateSetPeftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peft",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ResetAcceptEula() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceptEula",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ResetCustomizationTechnique() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizationTechnique",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ResetEvaluationType() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluationType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ResetEvaluatorArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluatorArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ResetPeft() {
	_jsii_.InvokeVoid(
		t,
		"resetPeft",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTrainingJob_ServerlessJobConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

