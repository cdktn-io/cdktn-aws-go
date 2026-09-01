package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CheckpointConfiguration() AwsKinesisanalyticsv2Application_CheckpointConfigurationPropertyOutputReference
	// Experimental.
	CheckpointConfigurationInput() *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty
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
	InternalValue() *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty)
	// Experimental.
	MonitoringConfiguration() AwsKinesisanalyticsv2Application_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty
	// Experimental.
	ParallelismConfiguration() AwsKinesisanalyticsv2Application_ParallelismConfigurationPropertyOutputReference
	// Experimental.
	ParallelismConfigurationInput() *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty
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
	PutCheckpointConfiguration(value *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty)
	// Experimental.
	PutParallelismConfiguration(value *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty)
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

// The jsii proxy struct for AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference
type jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfiguration() AwsKinesisanalyticsv2Application_CheckpointConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_CheckpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"checkpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) CheckpointConfigurationInput() *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty
	_jsii_.Get(
		j,
		"checkpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) InternalValue() *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfiguration() AwsKinesisanalyticsv2Application_MonitoringConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) MonitoringConfigurationInput() *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfiguration() AwsKinesisanalyticsv2Application_ParallelismConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_ParallelismConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelismConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ParallelismConfigurationInput() *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty
	_jsii_.Get(
		j,
		"parallelismConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference_Override(a AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.FlinkApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference)SetInternalValue(val *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) PutCheckpointConfiguration(value *AwsKinesisanalyticsv2Application_CheckpointConfigurationProperty) {
	if err := a.validatePutCheckpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCheckpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *AwsKinesisanalyticsv2Application_MonitoringConfigurationProperty) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) PutParallelismConfiguration(value *AwsKinesisanalyticsv2Application_ParallelismConfigurationProperty) {
	if err := a.validatePutParallelismConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelismConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ResetCheckpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ResetParallelismConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelismConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

