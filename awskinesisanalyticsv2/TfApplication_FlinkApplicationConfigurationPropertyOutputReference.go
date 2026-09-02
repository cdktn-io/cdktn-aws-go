package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_FlinkApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CheckpointConfiguration() TfApplication_CheckpointConfigurationPropertyOutputReference
	// Experimental.
	CheckpointConfigurationInput() *TfApplication_CheckpointConfigurationProperty
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
	InternalValue() *TfApplication_FlinkApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfApplication_FlinkApplicationConfigurationProperty)
	// Experimental.
	MonitoringConfiguration() TfApplication_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *TfApplication_MonitoringConfigurationProperty
	// Experimental.
	ParallelismConfiguration() TfApplication_ParallelismConfigurationPropertyOutputReference
	// Experimental.
	ParallelismConfigurationInput() *TfApplication_ParallelismConfigurationProperty
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
	PutCheckpointConfiguration(value *TfApplication_CheckpointConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *TfApplication_MonitoringConfigurationProperty)
	// Experimental.
	PutParallelismConfiguration(value *TfApplication_ParallelismConfigurationProperty)
	// Experimental.
	ResetCheckpointConfiguration()
	// Experimental.
	ResetMonitoringConfiguration()
	// Experimental.
	ResetParallelismConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_FlinkApplicationConfigurationPropertyOutputReference
type jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfiguration() TfApplication_CheckpointConfigurationPropertyOutputReference {
	var returns TfApplication_CheckpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"checkpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfigurationInput() *TfApplication_CheckpointConfigurationProperty {
	var returns *TfApplication_CheckpointConfigurationProperty
	_jsii_.Get(
		j,
		"checkpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) InternalValue() *TfApplication_FlinkApplicationConfigurationProperty {
	var returns *TfApplication_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfiguration() TfApplication_MonitoringConfigurationPropertyOutputReference {
	var returns TfApplication_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfigurationInput() *TfApplication_MonitoringConfigurationProperty {
	var returns *TfApplication_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfiguration() TfApplication_ParallelismConfigurationPropertyOutputReference {
	var returns TfApplication_ParallelismConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelismConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfigurationInput() *TfApplication_ParallelismConfigurationProperty {
	var returns *TfApplication_ParallelismConfigurationProperty
	_jsii_.Get(
		j,
		"parallelismConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_FlinkApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_FlinkApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_FlinkApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_FlinkApplicationConfigurationPropertyOutputReference_Override(t TfApplication_FlinkApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference)SetInternalValue(val *TfApplication_FlinkApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) PutCheckpointConfiguration(value *TfApplication_CheckpointConfigurationProperty) {
	if err := t.validatePutCheckpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCheckpointConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *TfApplication_MonitoringConfigurationProperty) {
	if err := t.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) PutParallelismConfiguration(value *TfApplication_ParallelismConfigurationProperty) {
	if err := t.validatePutParallelismConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParallelismConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetCheckpointConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetParallelismConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelismConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_FlinkApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

