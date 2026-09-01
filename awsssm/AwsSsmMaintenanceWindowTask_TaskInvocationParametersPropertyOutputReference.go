package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutomationParameters() AwsSsmMaintenanceWindowTask_AutomationParametersPropertyOutputReference
	// Experimental.
	AutomationParametersInput() *AwsSsmMaintenanceWindowTask_AutomationParametersProperty
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
	InternalValue() *AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty
	// Experimental.
	SetInternalValue(val *AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty)
	// Experimental.
	LambdaParameters() AwsSsmMaintenanceWindowTask_LambdaParametersPropertyOutputReference
	// Experimental.
	LambdaParametersInput() *AwsSsmMaintenanceWindowTask_LambdaParametersProperty
	// Experimental.
	RunCommandParameters() AwsSsmMaintenanceWindowTask_RunCommandParametersPropertyOutputReference
	// Experimental.
	RunCommandParametersInput() *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty
	// Experimental.
	StepFunctionsParameters() AwsSsmMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference
	// Experimental.
	StepFunctionsParametersInput() *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty
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
	PutAutomationParameters(value *AwsSsmMaintenanceWindowTask_AutomationParametersProperty)
	// Experimental.
	PutLambdaParameters(value *AwsSsmMaintenanceWindowTask_LambdaParametersProperty)
	// Experimental.
	PutRunCommandParameters(value *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty)
	// Experimental.
	PutStepFunctionsParameters(value *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty)
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

// The jsii proxy struct for AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference
type jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) AutomationParameters() AwsSsmMaintenanceWindowTask_AutomationParametersPropertyOutputReference {
	var returns AwsSsmMaintenanceWindowTask_AutomationParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"automationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) AutomationParametersInput() *AwsSsmMaintenanceWindowTask_AutomationParametersProperty {
	var returns *AwsSsmMaintenanceWindowTask_AutomationParametersProperty
	_jsii_.Get(
		j,
		"automationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InternalValue() *AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty {
	var returns *AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) LambdaParameters() AwsSsmMaintenanceWindowTask_LambdaParametersPropertyOutputReference {
	var returns AwsSsmMaintenanceWindowTask_LambdaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) LambdaParametersInput() *AwsSsmMaintenanceWindowTask_LambdaParametersProperty {
	var returns *AwsSsmMaintenanceWindowTask_LambdaParametersProperty
	_jsii_.Get(
		j,
		"lambdaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) RunCommandParameters() AwsSsmMaintenanceWindowTask_RunCommandParametersPropertyOutputReference {
	var returns AwsSsmMaintenanceWindowTask_RunCommandParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"runCommandParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) RunCommandParametersInput() *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty {
	var returns *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty
	_jsii_.Get(
		j,
		"runCommandParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) StepFunctionsParameters() AwsSsmMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference {
	var returns AwsSsmMaintenanceWindowTask_StepFunctionsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) StepFunctionsParametersInput() *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty {
	var returns *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty
	_jsii_.Get(
		j,
		"stepFunctionsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmMaintenanceWindowTask.TaskInvocationParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference_Override(a AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmMaintenanceWindowTask.TaskInvocationParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetInternalValue(val *AwsSsmMaintenanceWindowTask_TaskInvocationParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutAutomationParameters(value *AwsSsmMaintenanceWindowTask_AutomationParametersProperty) {
	if err := a.validatePutAutomationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutomationParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutLambdaParameters(value *AwsSsmMaintenanceWindowTask_LambdaParametersProperty) {
	if err := a.validatePutLambdaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutRunCommandParameters(value *AwsSsmMaintenanceWindowTask_RunCommandParametersProperty) {
	if err := a.validatePutRunCommandParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunCommandParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) PutStepFunctionsParameters(value *AwsSsmMaintenanceWindowTask_StepFunctionsParametersProperty) {
	if err := a.validatePutStepFunctionsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctionsParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetAutomationParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomationParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetLambdaParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetRunCommandParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRunCommandParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ResetStepFunctionsParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctionsParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_TaskInvocationParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

