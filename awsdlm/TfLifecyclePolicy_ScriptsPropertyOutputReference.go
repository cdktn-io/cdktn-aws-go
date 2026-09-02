package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLifecyclePolicy_ScriptsPropertyOutputReference interface {
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
	ExecuteOperationOnScriptFailure() interface{}
	// Experimental.
	SetExecuteOperationOnScriptFailure(val interface{})
	// Experimental.
	ExecuteOperationOnScriptFailureInput() interface{}
	// Experimental.
	ExecutionHandler() *string
	// Experimental.
	SetExecutionHandler(val *string)
	// Experimental.
	ExecutionHandlerInput() *string
	// Experimental.
	ExecutionHandlerService() *string
	// Experimental.
	SetExecutionHandlerService(val *string)
	// Experimental.
	ExecutionHandlerServiceInput() *string
	// Experimental.
	ExecutionTimeout() *float64
	// Experimental.
	SetExecutionTimeout(val *float64)
	// Experimental.
	ExecutionTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfLifecyclePolicy_ScriptsProperty
	// Experimental.
	SetInternalValue(val *TfLifecyclePolicy_ScriptsProperty)
	// Experimental.
	MaximumRetryCount() *float64
	// Experimental.
	SetMaximumRetryCount(val *float64)
	// Experimental.
	MaximumRetryCountInput() *float64
	// Experimental.
	Stages() *[]*string
	// Experimental.
	SetStages(val *[]*string)
	// Experimental.
	StagesInput() *[]*string
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
	ResetExecuteOperationOnScriptFailure()
	// Experimental.
	ResetExecutionHandlerService()
	// Experimental.
	ResetExecutionTimeout()
	// Experimental.
	ResetMaximumRetryCount()
	// Experimental.
	ResetStages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLifecyclePolicy_ScriptsPropertyOutputReference
type jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecuteOperationOnScriptFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecuteOperationOnScriptFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) InternalValue() *TfLifecyclePolicy_ScriptsProperty {
	var returns *TfLifecyclePolicy_ScriptsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) MaximumRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) MaximumRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) StagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLifecyclePolicy_ScriptsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLifecyclePolicy_ScriptsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLifecyclePolicy_ScriptsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.ScriptsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLifecyclePolicy_ScriptsPropertyOutputReference_Override(t TfLifecyclePolicy_ScriptsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.ScriptsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetExecuteOperationOnScriptFailure(val interface{}) {
	if err := j.validateSetExecuteOperationOnScriptFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeOperationOnScriptFailure",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionHandler(val *string) {
	if err := j.validateSetExecutionHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandler",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionHandlerService(val *string) {
	if err := j.validateSetExecutionHandlerServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandlerService",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionTimeout(val *float64) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetInternalValue(val *TfLifecyclePolicy_ScriptsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetMaximumRetryCount(val *float64) {
	if err := j.validateSetMaximumRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryCount",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetStages(val *[]*string) {
	if err := j.validateSetStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stages",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecuteOperationOnScriptFailure() {
	_jsii_.InvokeVoid(
		t,
		"resetExecuteOperationOnScriptFailure",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecutionHandlerService() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionHandlerService",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ResetMaximumRetryCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRetryCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ResetStages() {
	_jsii_.InvokeVoid(
		t,
		"resetStages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_ScriptsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

