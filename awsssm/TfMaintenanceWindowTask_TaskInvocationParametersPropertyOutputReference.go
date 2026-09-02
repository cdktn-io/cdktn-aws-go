package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutomationParameters() TfMaintenanceWindowTask_AutomationParametersPropertyOutputReference
	// Experimental.
	AutomationParametersInput() *TfMaintenanceWindowTask_AutomationParametersProperty
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
	InternalValue() *TfMaintenanceWindowTask_TaskInvocationParametersProperty
	// Experimental.
	SetInternalValue(val *TfMaintenanceWindowTask_TaskInvocationParametersProperty)
	// Experimental.
	LambdaParameters() TfMaintenanceWindowTask_LambdaParametersPropertyOutputReference
	// Experimental.
	LambdaParametersInput() *TfMaintenanceWindowTask_LambdaParametersProperty
	// Experimental.
	RunCommandParameters() TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference
	// Experimental.
	RunCommandParametersInput() *TfMaintenanceWindowTask_RunCommandParametersProperty
	// Experimental.
	StepFunctionsParameters() TfMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference
	// Experimental.
	StepFunctionsParametersInput() *TfMaintenanceWindowTask_StepFunctionsParametersProperty
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
	PutAutomationParameters(value *TfMaintenanceWindowTask_AutomationParametersProperty)
	// Experimental.
	PutLambdaParameters(value *TfMaintenanceWindowTask_LambdaParametersProperty)
	// Experimental.
	PutRunCommandParameters(value *TfMaintenanceWindowTask_RunCommandParametersProperty)
	// Experimental.
	PutStepFunctionsParameters(value *TfMaintenanceWindowTask_StepFunctionsParametersProperty)
	// Experimental.
	ResetAutomationParameters()
	// Experimental.
	ResetLambdaParameters()
	// Experimental.
	ResetRunCommandParameters()
	// Experimental.
	ResetStepFunctionsParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference
type jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) AutomationParameters() TfMaintenanceWindowTask_AutomationParametersPropertyOutputReference {
	var returns TfMaintenanceWindowTask_AutomationParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"automationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) AutomationParametersInput() *TfMaintenanceWindowTask_AutomationParametersProperty {
	var returns *TfMaintenanceWindowTask_AutomationParametersProperty
	_jsii_.Get(
		j,
		"automationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InternalValue() *TfMaintenanceWindowTask_TaskInvocationParametersProperty {
	var returns *TfMaintenanceWindowTask_TaskInvocationParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) LambdaParameters() TfMaintenanceWindowTask_LambdaParametersPropertyOutputReference {
	var returns TfMaintenanceWindowTask_LambdaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) LambdaParametersInput() *TfMaintenanceWindowTask_LambdaParametersProperty {
	var returns *TfMaintenanceWindowTask_LambdaParametersProperty
	_jsii_.Get(
		j,
		"lambdaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) RunCommandParameters() TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference {
	var returns TfMaintenanceWindowTask_RunCommandParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"runCommandParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) RunCommandParametersInput() *TfMaintenanceWindowTask_RunCommandParametersProperty {
	var returns *TfMaintenanceWindowTask_RunCommandParametersProperty
	_jsii_.Get(
		j,
		"runCommandParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) StepFunctionsParameters() TfMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference {
	var returns TfMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) StepFunctionsParametersInput() *TfMaintenanceWindowTask_StepFunctionsParametersProperty {
	var returns *TfMaintenanceWindowTask_StepFunctionsParametersProperty
	_jsii_.Get(
		j,
		"stepFunctionsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask.TaskInvocationParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference_Override(t TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.TfMaintenanceWindowTask.TaskInvocationParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetInternalValue(val *TfMaintenanceWindowTask_TaskInvocationParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutAutomationParameters(value *TfMaintenanceWindowTask_AutomationParametersProperty) {
	if err := t.validatePutAutomationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutomationParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutLambdaParameters(value *TfMaintenanceWindowTask_LambdaParametersProperty) {
	if err := t.validatePutLambdaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutRunCommandParameters(value *TfMaintenanceWindowTask_RunCommandParametersProperty) {
	if err := t.validatePutRunCommandParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRunCommandParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutStepFunctionsParameters(value *TfMaintenanceWindowTask_StepFunctionsParametersProperty) {
	if err := t.validatePutStepFunctionsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStepFunctionsParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetAutomationParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetAutomationParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetLambdaParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetRunCommandParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetRunCommandParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetStepFunctionsParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetStepFunctionsParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

