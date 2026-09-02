package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_RollingUpdatePolicyPropertyOutputReference interface {
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
	InternalValue() *TfEndpoint_RollingUpdatePolicyProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_RollingUpdatePolicyProperty)
	// Experimental.
	MaximumBatchSize() TfEndpoint_MaximumBatchSizePropertyOutputReference
	// Experimental.
	MaximumBatchSizeInput() *TfEndpoint_MaximumBatchSizeProperty
	// Experimental.
	MaximumExecutionTimeoutInSeconds() *float64
	// Experimental.
	SetMaximumExecutionTimeoutInSeconds(val *float64)
	// Experimental.
	MaximumExecutionTimeoutInSecondsInput() *float64
	// Experimental.
	RollbackMaximumBatchSize() TfEndpoint_RollbackMaximumBatchSizePropertyOutputReference
	// Experimental.
	RollbackMaximumBatchSizeInput() *TfEndpoint_RollbackMaximumBatchSizeProperty
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
	PutMaximumBatchSize(value *TfEndpoint_MaximumBatchSizeProperty)
	// Experimental.
	PutRollbackMaximumBatchSize(value *TfEndpoint_RollbackMaximumBatchSizeProperty)
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

// The jsii proxy struct for TfEndpoint_RollingUpdatePolicyPropertyOutputReference
type jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) InternalValue() *TfEndpoint_RollingUpdatePolicyProperty {
	var returns *TfEndpoint_RollingUpdatePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumBatchSize() TfEndpoint_MaximumBatchSizePropertyOutputReference {
	var returns TfEndpoint_MaximumBatchSizePropertyOutputReference
	_jsii_.Get(
		j,
		"maximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumBatchSizeInput() *TfEndpoint_MaximumBatchSizeProperty {
	var returns *TfEndpoint_MaximumBatchSizeProperty
	_jsii_.Get(
		j,
		"maximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) RollbackMaximumBatchSize() TfEndpoint_RollbackMaximumBatchSizePropertyOutputReference {
	var returns TfEndpoint_RollbackMaximumBatchSizePropertyOutputReference
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) RollbackMaximumBatchSizeInput() *TfEndpoint_RollbackMaximumBatchSizeProperty {
	var returns *TfEndpoint_RollbackMaximumBatchSizeProperty
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) WaitIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) WaitIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_RollingUpdatePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_RollingUpdatePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_RollingUpdatePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.RollingUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_RollingUpdatePolicyPropertyOutputReference_Override(t TfEndpoint_RollingUpdatePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.RollingUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetInternalValue(val *TfEndpoint_RollingUpdatePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetMaximumExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetMaximumExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference)SetWaitIntervalInSeconds(val *float64) {
	if err := j.validateSetWaitIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIntervalInSeconds",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) PutMaximumBatchSize(value *TfEndpoint_MaximumBatchSizeProperty) {
	if err := t.validatePutMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaximumBatchSize",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) PutRollbackMaximumBatchSize(value *TfEndpoint_RollbackMaximumBatchSizeProperty) {
	if err := t.validatePutRollbackMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRollbackMaximumBatchSize",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ResetMaximumExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ResetRollbackMaximumBatchSize() {
	_jsii_.InvokeVoid(
		t,
		"resetRollbackMaximumBatchSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_RollingUpdatePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

