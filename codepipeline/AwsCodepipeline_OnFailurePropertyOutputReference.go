package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodepipeline_OnFailurePropertyOutputReference interface {
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
	Condition() AwsCodepipeline_StageOnFailureConditionPropertyOutputReference
	// Experimental.
	ConditionInput() *AwsCodepipeline_StageOnFailureConditionProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCodepipeline_OnFailureProperty
	// Experimental.
	SetInternalValue(val *AwsCodepipeline_OnFailureProperty)
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
	// Experimental.
	ResultInput() *string
	// Experimental.
	RetryConfiguration() AwsCodepipeline_RetryConfigurationPropertyOutputReference
	// Experimental.
	RetryConfigurationInput() *AwsCodepipeline_RetryConfigurationProperty
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
	PutCondition(value *AwsCodepipeline_StageOnFailureConditionProperty)
	// Experimental.
	PutRetryConfiguration(value *AwsCodepipeline_RetryConfigurationProperty)
	// Experimental.
	ResetCondition()
	// Experimental.
	ResetResult()
	// Experimental.
	ResetRetryConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCodepipeline_OnFailurePropertyOutputReference
type jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) Condition() AwsCodepipeline_StageOnFailureConditionPropertyOutputReference {
	var returns AwsCodepipeline_StageOnFailureConditionPropertyOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ConditionInput() *AwsCodepipeline_StageOnFailureConditionProperty {
	var returns *AwsCodepipeline_StageOnFailureConditionProperty
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) InternalValue() *AwsCodepipeline_OnFailureProperty {
	var returns *AwsCodepipeline_OnFailureProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ResultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) RetryConfiguration() AwsCodepipeline_RetryConfigurationPropertyOutputReference {
	var returns AwsCodepipeline_RetryConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"retryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) RetryConfigurationInput() *AwsCodepipeline_RetryConfigurationProperty {
	var returns *AwsCodepipeline_RetryConfigurationProperty
	_jsii_.Get(
		j,
		"retryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodepipeline_OnFailurePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCodepipeline_OnFailurePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodepipeline_OnFailurePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.OnFailurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodepipeline_OnFailurePropertyOutputReference_Override(a AwsCodepipeline_OnFailurePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.OnFailurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetInternalValue(val *AwsCodepipeline_OnFailureProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetResult(val *string) {
	if err := j.validateSetResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) PutCondition(value *AwsCodepipeline_StageOnFailureConditionProperty) {
	if err := a.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) PutRetryConfiguration(value *AwsCodepipeline_RetryConfigurationProperty) {
	if err := a.validatePutRetryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ResetResult() {
	_jsii_.InvokeVoid(
		a,
		"resetResult",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ResetRetryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_OnFailurePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

