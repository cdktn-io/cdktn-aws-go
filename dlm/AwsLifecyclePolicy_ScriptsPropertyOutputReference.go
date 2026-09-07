package dlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLifecyclePolicy_ScriptsPropertyOutputReference interface {
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
	InternalValue() *AwsLifecyclePolicy_ScriptsProperty
	// Experimental.
	SetInternalValue(val *AwsLifecyclePolicy_ScriptsProperty)
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

// The jsii proxy struct for AwsLifecyclePolicy_ScriptsPropertyOutputReference
type jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecuteOperationOnScriptFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecuteOperationOnScriptFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executeOperationOnScriptFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionHandlerServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionHandlerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ExecutionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) InternalValue() *AwsLifecyclePolicy_ScriptsProperty {
	var returns *AwsLifecyclePolicy_ScriptsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) MaximumRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) MaximumRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) StagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLifecyclePolicy_ScriptsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLifecyclePolicy_ScriptsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLifecyclePolicy_ScriptsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.ScriptsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLifecyclePolicy_ScriptsPropertyOutputReference_Override(a AwsLifecyclePolicy_ScriptsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.ScriptsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetExecuteOperationOnScriptFailure(val interface{}) {
	if err := j.validateSetExecuteOperationOnScriptFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeOperationOnScriptFailure",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionHandler(val *string) {
	if err := j.validateSetExecutionHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandler",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionHandlerService(val *string) {
	if err := j.validateSetExecutionHandlerServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionHandlerService",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetExecutionTimeout(val *float64) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetInternalValue(val *AwsLifecyclePolicy_ScriptsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetMaximumRetryCount(val *float64) {
	if err := j.validateSetMaximumRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryCount",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetStages(val *[]*string) {
	if err := j.validateSetStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stages",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecuteOperationOnScriptFailure() {
	_jsii_.InvokeVoid(
		a,
		"resetExecuteOperationOnScriptFailure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecutionHandlerService() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionHandlerService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ResetMaximumRetryCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRetryCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ResetStages() {
	_jsii_.InvokeVoid(
		a,
		"resetStages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ScriptsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

