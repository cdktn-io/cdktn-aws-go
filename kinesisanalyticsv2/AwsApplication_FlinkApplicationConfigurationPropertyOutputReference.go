package kinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_FlinkApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CheckpointConfiguration() AwsApplication_CheckpointConfigurationPropertyOutputReference
	// Experimental.
	CheckpointConfigurationInput() *AwsApplication_CheckpointConfigurationProperty
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
	InternalValue() *AwsApplication_FlinkApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_FlinkApplicationConfigurationProperty)
	// Experimental.
	MonitoringConfiguration() AwsApplication_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *AwsApplication_MonitoringConfigurationProperty
	// Experimental.
	ParallelismConfiguration() AwsApplication_ParallelismConfigurationPropertyOutputReference
	// Experimental.
	ParallelismConfigurationInput() *AwsApplication_ParallelismConfigurationProperty
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
	PutCheckpointConfiguration(value *AwsApplication_CheckpointConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *AwsApplication_MonitoringConfigurationProperty)
	// Experimental.
	PutParallelismConfiguration(value *AwsApplication_ParallelismConfigurationProperty)
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

// The jsii proxy struct for AwsApplication_FlinkApplicationConfigurationPropertyOutputReference
type jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfiguration() AwsApplication_CheckpointConfigurationPropertyOutputReference {
	var returns AwsApplication_CheckpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"checkpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfigurationInput() *AwsApplication_CheckpointConfigurationProperty {
	var returns *AwsApplication_CheckpointConfigurationProperty
	_jsii_.Get(
		j,
		"checkpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) InternalValue() *AwsApplication_FlinkApplicationConfigurationProperty {
	var returns *AwsApplication_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfiguration() AwsApplication_MonitoringConfigurationPropertyOutputReference {
	var returns AwsApplication_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfigurationInput() *AwsApplication_MonitoringConfigurationProperty {
	var returns *AwsApplication_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfiguration() AwsApplication_ParallelismConfigurationPropertyOutputReference {
	var returns AwsApplication_ParallelismConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelismConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfigurationInput() *AwsApplication_ParallelismConfigurationProperty {
	var returns *AwsApplication_ParallelismConfigurationProperty
	_jsii_.Get(
		j,
		"parallelismConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_FlinkApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_FlinkApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_FlinkApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_FlinkApplicationConfigurationPropertyOutputReference_Override(a AwsApplication_FlinkApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference)SetInternalValue(val *AwsApplication_FlinkApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) PutCheckpointConfiguration(value *AwsApplication_CheckpointConfigurationProperty) {
	if err := a.validatePutCheckpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCheckpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *AwsApplication_MonitoringConfigurationProperty) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) PutParallelismConfiguration(value *AwsApplication_ParallelismConfigurationProperty) {
	if err := a.validatePutParallelismConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelismConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetCheckpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ResetParallelismConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelismConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_FlinkApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

