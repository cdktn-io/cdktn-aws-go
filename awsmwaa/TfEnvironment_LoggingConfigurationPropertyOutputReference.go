package awsmwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmwaa/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmwaa/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEnvironment_LoggingConfigurationPropertyOutputReference interface {
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
	DagProcessingLogs() TfEnvironment_DagProcessingLogsPropertyOutputReference
	// Experimental.
	DagProcessingLogsInput() *TfEnvironment_DagProcessingLogsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEnvironment_LoggingConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfEnvironment_LoggingConfigurationProperty)
	// Experimental.
	SchedulerLogs() TfEnvironment_SchedulerLogsPropertyOutputReference
	// Experimental.
	SchedulerLogsInput() *TfEnvironment_SchedulerLogsProperty
	// Experimental.
	TaskLogs() TfEnvironment_TaskLogsPropertyOutputReference
	// Experimental.
	TaskLogsInput() *TfEnvironment_TaskLogsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WebserverLogs() TfEnvironment_WebserverLogsPropertyOutputReference
	// Experimental.
	WebserverLogsInput() *TfEnvironment_WebserverLogsProperty
	// Experimental.
	WorkerLogs() TfEnvironment_WorkerLogsPropertyOutputReference
	// Experimental.
	WorkerLogsInput() *TfEnvironment_WorkerLogsProperty
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
	PutDagProcessingLogs(value *TfEnvironment_DagProcessingLogsProperty)
	// Experimental.
	PutSchedulerLogs(value *TfEnvironment_SchedulerLogsProperty)
	// Experimental.
	PutTaskLogs(value *TfEnvironment_TaskLogsProperty)
	// Experimental.
	PutWebserverLogs(value *TfEnvironment_WebserverLogsProperty)
	// Experimental.
	PutWorkerLogs(value *TfEnvironment_WorkerLogsProperty)
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

// The jsii proxy struct for TfEnvironment_LoggingConfigurationPropertyOutputReference
type jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogs() TfEnvironment_DagProcessingLogsPropertyOutputReference {
	var returns TfEnvironment_DagProcessingLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"dagProcessingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) DagProcessingLogsInput() *TfEnvironment_DagProcessingLogsProperty {
	var returns *TfEnvironment_DagProcessingLogsProperty
	_jsii_.Get(
		j,
		"dagProcessingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) InternalValue() *TfEnvironment_LoggingConfigurationProperty {
	var returns *TfEnvironment_LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogs() TfEnvironment_SchedulerLogsPropertyOutputReference {
	var returns TfEnvironment_SchedulerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"schedulerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) SchedulerLogsInput() *TfEnvironment_SchedulerLogsProperty {
	var returns *TfEnvironment_SchedulerLogsProperty
	_jsii_.Get(
		j,
		"schedulerLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogs() TfEnvironment_TaskLogsPropertyOutputReference {
	var returns TfEnvironment_TaskLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"taskLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) TaskLogsInput() *TfEnvironment_TaskLogsProperty {
	var returns *TfEnvironment_TaskLogsProperty
	_jsii_.Get(
		j,
		"taskLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogs() TfEnvironment_WebserverLogsPropertyOutputReference {
	var returns TfEnvironment_WebserverLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"webserverLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) WebserverLogsInput() *TfEnvironment_WebserverLogsProperty {
	var returns *TfEnvironment_WebserverLogsProperty
	_jsii_.Get(
		j,
		"webserverLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogs() TfEnvironment_WorkerLogsPropertyOutputReference {
	var returns TfEnvironment_WorkerLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"workerLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) WorkerLogsInput() *TfEnvironment_WorkerLogsProperty {
	var returns *TfEnvironment_WorkerLogsProperty
	_jsii_.Get(
		j,
		"workerLogsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEnvironment_LoggingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEnvironment_LoggingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEnvironment_LoggingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-mwaa.TfEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEnvironment_LoggingConfigurationPropertyOutputReference_Override(t TfEnvironment_LoggingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mwaa.TfEnvironment.LoggingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference)SetInternalValue(val *TfEnvironment_LoggingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) PutDagProcessingLogs(value *TfEnvironment_DagProcessingLogsProperty) {
	if err := t.validatePutDagProcessingLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDagProcessingLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) PutSchedulerLogs(value *TfEnvironment_SchedulerLogsProperty) {
	if err := t.validatePutSchedulerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchedulerLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) PutTaskLogs(value *TfEnvironment_TaskLogsProperty) {
	if err := t.validatePutTaskLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTaskLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) PutWebserverLogs(value *TfEnvironment_WebserverLogsProperty) {
	if err := t.validatePutWebserverLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWebserverLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) PutWorkerLogs(value *TfEnvironment_WorkerLogsProperty) {
	if err := t.validatePutWorkerLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkerLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ResetDagProcessingLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetDagProcessingLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ResetSchedulerLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedulerLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ResetTaskLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ResetWebserverLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetWebserverLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ResetWorkerLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkerLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEnvironment_LoggingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

