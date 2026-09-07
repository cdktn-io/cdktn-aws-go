package mwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/mwaa/jsii"

	"github.com/cdktn-io/cdktn-aws-go/mwaa/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEnvironment_LoggingConfigurationPropertyOutputReference interface {
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
	DagProcessingLogs() AwsEnvironment_DagProcessingLogsPropertyOutputReference
	// Experimental.
	DagProcessingLogsInput() *AwsEnvironment_DagProcessingLogsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEnvironment_LoggingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsEnvironment_LoggingConfigurationProperty)
	// Experimental.
	SchedulerLogs() AwsEnvironment_SchedulerLogsPropertyOutputReference
	// Experimental.
	SchedulerLogsInput() *AwsEnvironment_SchedulerLogsProperty
	// Experimental.
	TaskLogs() AwsEnvironment_TaskLogsPropertyOutputReference
	// Experimental.
	TaskLogsInput() *AwsEnvironment_TaskLogsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WebserverLogs() AwsEnvironment_WebserverLogsPropertyOutputReference
	// Experimental.
	WebserverLogsInput() *AwsEnvironment_WebserverLogsProperty
	// Experimental.
	WorkerLogs() AwsEnvironment_WorkerLogsPropertyOutputReference
	// Experimental.
	WorkerLogsInput() *AwsEnvironment_WorkerLogsProperty
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
	PutDagProcessingLogs(value *AwsEnvironment_DagProcessingLogsProperty)
	// Experimental.
	PutSchedulerLogs(value *AwsEnvironment_SchedulerLogsProperty)
	// Experimental.
	PutTaskLogs(value *AwsEnvironment_TaskLogsProperty)
	// Experimental.
	PutWebserverLogs(value *AwsEnvironment_WebserverLogsProperty)
	// Experimental.
	PutWorkerLogs(value *AwsEnvironment_WorkerLogsProperty)
	// Experimental.
	ResetDagProcessingLogs()
	// Experimental.
	ResetSchedulerLogs()
	// Experimental.
	ResetTaskLogs()
	// Experimental.
	ResetWebserverLogs()
	// Experimental.
	ResetWorkerLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEnvironment_LoggingConfigurationPropertyOutputReference
type jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogs() AwsEnvironment_DagProcessingLogsPropertyOutputReference {
	var returns AwsEnvironment_DagProcessingLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"dagProcessingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogsInput() *AwsEnvironment_DagProcessingLogsProperty {
	var returns *AwsEnvironment_DagProcessingLogsProperty
	_jsii_.Get(
		j,
		"dagProcessingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) InternalValue() *AwsEnvironment_LoggingConfigurationProperty {
	var returns *AwsEnvironment_LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogs() AwsEnvironment_SchedulerLogsPropertyOutputReference {
	var returns AwsEnvironment_SchedulerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"schedulerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogsInput() *AwsEnvironment_SchedulerLogsProperty {
	var returns *AwsEnvironment_SchedulerLogsProperty
	_jsii_.Get(
		j,
		"schedulerLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogs() AwsEnvironment_TaskLogsPropertyOutputReference {
	var returns AwsEnvironment_TaskLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"taskLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogsInput() *AwsEnvironment_TaskLogsProperty {
	var returns *AwsEnvironment_TaskLogsProperty
	_jsii_.Get(
		j,
		"taskLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogs() AwsEnvironment_WebserverLogsPropertyOutputReference {
	var returns AwsEnvironment_WebserverLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"webserverLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogsInput() *AwsEnvironment_WebserverLogsProperty {
	var returns *AwsEnvironment_WebserverLogsProperty
	_jsii_.Get(
		j,
		"webserverLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogs() AwsEnvironment_WorkerLogsPropertyOutputReference {
	var returns AwsEnvironment_WorkerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"workerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogsInput() *AwsEnvironment_WorkerLogsProperty {
	var returns *AwsEnvironment_WorkerLogsProperty
	_jsii_.Get(
		j,
		"workerLogsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEnvironment_LoggingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEnvironment_LoggingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEnvironment_LoggingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-mwaa.AwsEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEnvironment_LoggingConfigurationPropertyOutputReference_Override(a AwsEnvironment_LoggingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mwaa.AwsEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference)SetInternalValue(val *AwsEnvironment_LoggingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) PutDagProcessingLogs(value *AwsEnvironment_DagProcessingLogsProperty) {
	if err := a.validatePutDagProcessingLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDagProcessingLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) PutSchedulerLogs(value *AwsEnvironment_SchedulerLogsProperty) {
	if err := a.validatePutSchedulerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedulerLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) PutTaskLogs(value *AwsEnvironment_TaskLogsProperty) {
	if err := a.validatePutTaskLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTaskLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) PutWebserverLogs(value *AwsEnvironment_WebserverLogsProperty) {
	if err := a.validatePutWebserverLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebserverLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) PutWorkerLogs(value *AwsEnvironment_WorkerLogsProperty) {
	if err := a.validatePutWorkerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkerLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ResetDagProcessingLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetDagProcessingLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ResetSchedulerLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulerLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ResetTaskLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ResetWebserverLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetWebserverLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ResetWorkerLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkerLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEnvironment_LoggingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

