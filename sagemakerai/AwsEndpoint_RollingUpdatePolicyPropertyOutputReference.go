package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_RollingUpdatePolicyPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_RollingUpdatePolicyProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_RollingUpdatePolicyProperty)
	// Experimental.
	MaximumBatchSize() AwsEndpoint_MaximumBatchSizePropertyOutputReference
	// Experimental.
	MaximumBatchSizeInput() *AwsEndpoint_MaximumBatchSizeProperty
	// Experimental.
	MaximumExecutionTimeoutInSeconds() *float64
	// Experimental.
	SetMaximumExecutionTimeoutInSeconds(val *float64)
	// Experimental.
	MaximumExecutionTimeoutInSecondsInput() *float64
	// Experimental.
	RollbackMaximumBatchSize() AwsEndpoint_RollbackMaximumBatchSizePropertyOutputReference
	// Experimental.
	RollbackMaximumBatchSizeInput() *AwsEndpoint_RollbackMaximumBatchSizeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WaitIntervalInSeconds() *float64
	// Experimental.
	SetWaitIntervalInSeconds(val *float64)
	// Experimental.
	WaitIntervalInSecondsInput() *float64
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
	PutMaximumBatchSize(value *AwsEndpoint_MaximumBatchSizeProperty)
	// Experimental.
	PutRollbackMaximumBatchSize(value *AwsEndpoint_RollbackMaximumBatchSizeProperty)
	// Experimental.
	ResetMaximumExecutionTimeoutInSeconds()
	// Experimental.
	ResetRollbackMaximumBatchSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEndpoint_RollingUpdatePolicyPropertyOutputReference
type jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) InternalValue() *AwsEndpoint_RollingUpdatePolicyProperty {
	var returns *AwsEndpoint_RollingUpdatePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumBatchSize() AwsEndpoint_MaximumBatchSizePropertyOutputReference {
	var returns AwsEndpoint_MaximumBatchSizePropertyOutputReference
	_jsii_.Get(
		j,
		"maximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumBatchSizeInput() *AwsEndpoint_MaximumBatchSizeProperty {
	var returns *AwsEndpoint_MaximumBatchSizeProperty
	_jsii_.Get(
		j,
		"maximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) RollbackMaximumBatchSize() AwsEndpoint_RollbackMaximumBatchSizePropertyOutputReference {
	var returns AwsEndpoint_RollbackMaximumBatchSizePropertyOutputReference
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) RollbackMaximumBatchSizeInput() *AwsEndpoint_RollbackMaximumBatchSizeProperty {
	var returns *AwsEndpoint_RollbackMaximumBatchSizeProperty
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) WaitIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) WaitIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_RollingUpdatePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_RollingUpdatePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_RollingUpdatePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.RollingUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_RollingUpdatePolicyPropertyOutputReference_Override(a AwsEndpoint_RollingUpdatePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.RollingUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetInternalValue(val *AwsEndpoint_RollingUpdatePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetMaximumExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetMaximumExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference)SetWaitIntervalInSeconds(val *float64) {
	if err := j.validateSetWaitIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIntervalInSeconds",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) PutMaximumBatchSize(value *AwsEndpoint_MaximumBatchSizeProperty) {
	if err := a.validatePutMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaximumBatchSize",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) PutRollbackMaximumBatchSize(value *AwsEndpoint_RollbackMaximumBatchSizeProperty) {
	if err := a.validatePutRollbackMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRollbackMaximumBatchSize",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ResetMaximumExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ResetRollbackMaximumBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetRollbackMaximumBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_RollingUpdatePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

