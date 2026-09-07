package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BestObjectiveNotImproving() AwsHyperParameterTuningJob_BestObjectiveNotImprovingPropertyList
	// Experimental.
	BestObjectiveNotImprovingInput() interface{}
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
	ConvergenceDetected() AwsHyperParameterTuningJob_ConvergenceDetectedPropertyList
	// Experimental.
	ConvergenceDetectedInput() interface{}
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
	TargetObjectiveMetricValue() *float64
	// Experimental.
	SetTargetObjectiveMetricValue(val *float64)
	// Experimental.
	TargetObjectiveMetricValueInput() *float64
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
	PutBestObjectiveNotImproving(value interface{})
	// Experimental.
	PutConvergenceDetected(value interface{})
	// Experimental.
	ResetBestObjectiveNotImproving()
	// Experimental.
	ResetConvergenceDetected()
	// Experimental.
	ResetTargetObjectiveMetricValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference
type jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) BestObjectiveNotImproving() AwsHyperParameterTuningJob_BestObjectiveNotImprovingPropertyList {
	var returns AwsHyperParameterTuningJob_BestObjectiveNotImprovingPropertyList
	_jsii_.Get(
		j,
		"bestObjectiveNotImproving",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) BestObjectiveNotImprovingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bestObjectiveNotImprovingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ConvergenceDetected() AwsHyperParameterTuningJob_ConvergenceDetectedPropertyList {
	var returns AwsHyperParameterTuningJob_ConvergenceDetectedPropertyList
	_jsii_.Get(
		j,
		"convergenceDetected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ConvergenceDetectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convergenceDetectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) TargetObjectiveMetricValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetObjectiveMetricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) TargetObjectiveMetricValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetObjectiveMetricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TuningJobCompletionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference_Override(a AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TuningJobCompletionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetTargetObjectiveMetricValue(val *float64) {
	if err := j.validateSetTargetObjectiveMetricValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetObjectiveMetricValue",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) PutBestObjectiveNotImproving(value interface{}) {
	if err := a.validatePutBestObjectiveNotImprovingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBestObjectiveNotImproving",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) PutConvergenceDetected(value interface{}) {
	if err := a.validatePutConvergenceDetectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConvergenceDetected",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ResetBestObjectiveNotImproving() {
	_jsii_.InvokeVoid(
		a,
		"resetBestObjectiveNotImproving",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ResetConvergenceDetected() {
	_jsii_.InvokeVoid(
		a,
		"resetConvergenceDetected",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ResetTargetObjectiveMetricValue() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetObjectiveMetricValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

