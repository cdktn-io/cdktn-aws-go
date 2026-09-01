package awsmwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmwaa/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmwaa/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference interface {
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
	DagProcessingLogs() AwsMwaaEnvironment_DagProcessingLogsPropertyOutputReference
	// Experimental.
	DagProcessingLogsInput() *AwsMwaaEnvironment_DagProcessingLogsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMwaaEnvironment_LoggingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsMwaaEnvironment_LoggingConfigurationProperty)
	// Experimental.
	SchedulerLogs() AwsMwaaEnvironment_SchedulerLogsPropertyOutputReference
	// Experimental.
	SchedulerLogsInput() *AwsMwaaEnvironment_SchedulerLogsProperty
	// Experimental.
	TaskLogs() AwsMwaaEnvironment_TaskLogsPropertyOutputReference
	// Experimental.
	TaskLogsInput() *AwsMwaaEnvironment_TaskLogsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WebserverLogs() AwsMwaaEnvironment_WebserverLogsPropertyOutputReference
	// Experimental.
	WebserverLogsInput() *AwsMwaaEnvironment_WebserverLogsProperty
	// Experimental.
	WorkerLogs() AwsMwaaEnvironment_WorkerLogsPropertyOutputReference
	// Experimental.
	WorkerLogsInput() *AwsMwaaEnvironment_WorkerLogsProperty
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
	PutDagProcessingLogs(value *AwsMwaaEnvironment_DagProcessingLogsProperty)
	// Experimental.
	PutSchedulerLogs(value *AwsMwaaEnvironment_SchedulerLogsProperty)
	// Experimental.
	PutTaskLogs(value *AwsMwaaEnvironment_TaskLogsProperty)
	// Experimental.
	PutWebserverLogs(value *AwsMwaaEnvironment_WebserverLogsProperty)
	// Experimental.
	PutWorkerLogs(value *AwsMwaaEnvironment_WorkerLogsProperty)
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

// The jsii proxy struct for AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference
type jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogs() AwsMwaaEnvironment_DagProcessingLogsPropertyOutputReference {
	var returns AwsMwaaEnvironment_DagProcessingLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"dagProcessingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogsInput() *AwsMwaaEnvironment_DagProcessingLogsProperty {
	var returns *AwsMwaaEnvironment_DagProcessingLogsProperty
	_jsii_.Get(
		j,
		"dagProcessingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) InternalValue() *AwsMwaaEnvironment_LoggingConfigurationProperty {
	var returns *AwsMwaaEnvironment_LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogs() AwsMwaaEnvironment_SchedulerLogsPropertyOutputReference {
	var returns AwsMwaaEnvironment_SchedulerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"schedulerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogsInput() *AwsMwaaEnvironment_SchedulerLogsProperty {
	var returns *AwsMwaaEnvironment_SchedulerLogsProperty
	_jsii_.Get(
		j,
		"schedulerLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogs() AwsMwaaEnvironment_TaskLogsPropertyOutputReference {
	var returns AwsMwaaEnvironment_TaskLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"taskLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogsInput() *AwsMwaaEnvironment_TaskLogsProperty {
	var returns *AwsMwaaEnvironment_TaskLogsProperty
	_jsii_.Get(
		j,
		"taskLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogs() AwsMwaaEnvironment_WebserverLogsPropertyOutputReference {
	var returns AwsMwaaEnvironment_WebserverLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"webserverLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogsInput() *AwsMwaaEnvironment_WebserverLogsProperty {
	var returns *AwsMwaaEnvironment_WebserverLogsProperty
	_jsii_.Get(
		j,
		"webserverLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogs() AwsMwaaEnvironment_WorkerLogsPropertyOutputReference {
	var returns AwsMwaaEnvironment_WorkerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"workerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogsInput() *AwsMwaaEnvironment_WorkerLogsProperty {
	var returns *AwsMwaaEnvironment_WorkerLogsProperty
	_jsii_.Get(
		j,
		"workerLogsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMwaaEnvironment_LoggingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-mwaa.AwsMwaaEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference_Override(a AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mwaa.AwsMwaaEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference)SetInternalValue(val *AwsMwaaEnvironment_LoggingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) PutDagProcessingLogs(value *AwsMwaaEnvironment_DagProcessingLogsProperty) {
	if err := a.validatePutDagProcessingLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDagProcessingLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) PutSchedulerLogs(value *AwsMwaaEnvironment_SchedulerLogsProperty) {
	if err := a.validatePutSchedulerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedulerLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) PutTaskLogs(value *AwsMwaaEnvironment_TaskLogsProperty) {
	if err := a.validatePutTaskLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTaskLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) PutWebserverLogs(value *AwsMwaaEnvironment_WebserverLogsProperty) {
	if err := a.validatePutWebserverLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebserverLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) PutWorkerLogs(value *AwsMwaaEnvironment_WorkerLogsProperty) {
	if err := a.validatePutWorkerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkerLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ResetDagProcessingLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetDagProcessingLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ResetSchedulerLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulerLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ResetTaskLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ResetWebserverLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetWebserverLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ResetWorkerLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkerLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMwaaEnvironment_LoggingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

