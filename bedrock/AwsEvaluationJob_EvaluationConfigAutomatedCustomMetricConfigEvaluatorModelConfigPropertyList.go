package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList interface {
	cdktn.ComplexList
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList
type jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList {
	_init_.Initialize()

	if err := validateNewAwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsEvaluationJob.EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList_Override(a AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsEvaluationJob.EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) Get(index *float64) AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEvaluationJob_EvaluationConfigAutomatedCustomMetricConfigEvaluatorModelConfigPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

