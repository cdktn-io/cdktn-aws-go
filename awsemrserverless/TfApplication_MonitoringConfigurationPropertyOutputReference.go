package awsemrserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrserverless/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemrserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLoggingConfiguration() TfApplication_CloudwatchLoggingConfigurationPropertyOutputReference
	// Experimental.
	CloudwatchLoggingConfigurationInput() *TfApplication_CloudwatchLoggingConfigurationProperty
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
	InternalValue() *TfApplication_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfApplication_MonitoringConfigurationProperty)
	// Experimental.
	ManagedPersistenceMonitoringConfiguration() TfApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference
	// Experimental.
	ManagedPersistenceMonitoringConfigurationInput() *TfApplication_ManagedPersistenceMonitoringConfigurationProperty
	// Experimental.
	PrometheusMonitoringConfiguration() TfApplication_PrometheusMonitoringConfigurationPropertyOutputReference
	// Experimental.
	PrometheusMonitoringConfigurationInput() *TfApplication_PrometheusMonitoringConfigurationProperty
	// Experimental.
	S3MonitoringConfiguration() TfApplication_S3MonitoringConfigurationPropertyOutputReference
	// Experimental.
	S3MonitoringConfigurationInput() *TfApplication_S3MonitoringConfigurationProperty
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
	PutCloudwatchLoggingConfiguration(value *TfApplication_CloudwatchLoggingConfigurationProperty)
	// Experimental.
	PutManagedPersistenceMonitoringConfiguration(value *TfApplication_ManagedPersistenceMonitoringConfigurationProperty)
	// Experimental.
	PutPrometheusMonitoringConfiguration(value *TfApplication_PrometheusMonitoringConfigurationProperty)
	// Experimental.
	PutS3MonitoringConfiguration(value *TfApplication_S3MonitoringConfigurationProperty)
	// Experimental.
	ResetCloudwatchLoggingConfiguration()
	// Experimental.
	ResetManagedPersistenceMonitoringConfiguration()
	// Experimental.
	ResetPrometheusMonitoringConfiguration()
	// Experimental.
	ResetS3MonitoringConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) CloudwatchLoggingConfiguration() TfApplication_CloudwatchLoggingConfigurationPropertyOutputReference {
	var returns TfApplication_CloudwatchLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) CloudwatchLoggingConfigurationInput() *TfApplication_CloudwatchLoggingConfigurationProperty {
	var returns *TfApplication_CloudwatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) InternalValue() *TfApplication_MonitoringConfigurationProperty {
	var returns *TfApplication_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ManagedPersistenceMonitoringConfiguration() TfApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference {
	var returns TfApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ManagedPersistenceMonitoringConfigurationInput() *TfApplication_ManagedPersistenceMonitoringConfigurationProperty {
	var returns *TfApplication_ManagedPersistenceMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PrometheusMonitoringConfiguration() TfApplication_PrometheusMonitoringConfigurationPropertyOutputReference {
	var returns TfApplication_PrometheusMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"prometheusMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PrometheusMonitoringConfigurationInput() *TfApplication_PrometheusMonitoringConfigurationProperty {
	var returns *TfApplication_PrometheusMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"prometheusMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfiguration() TfApplication_S3MonitoringConfigurationPropertyOutputReference {
	var returns TfApplication_S3MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3MonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfigurationInput() *TfApplication_S3MonitoringConfigurationProperty {
	var returns *TfApplication_S3MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"s3MonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.TfApplication.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_MonitoringConfigurationPropertyOutputReference_Override(t TfApplication_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.TfApplication.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *TfApplication_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PutCloudwatchLoggingConfiguration(value *TfApplication_CloudwatchLoggingConfigurationProperty) {
	if err := t.validatePutCloudwatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PutManagedPersistenceMonitoringConfiguration(value *TfApplication_ManagedPersistenceMonitoringConfigurationProperty) {
	if err := t.validatePutManagedPersistenceMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedPersistenceMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PutPrometheusMonitoringConfiguration(value *TfApplication_PrometheusMonitoringConfigurationProperty) {
	if err := t.validatePutPrometheusMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrometheusMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) PutS3MonitoringConfiguration(value *TfApplication_S3MonitoringConfigurationProperty) {
	if err := t.validatePutS3MonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3MonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ResetCloudwatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ResetManagedPersistenceMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedPersistenceMonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ResetPrometheusMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetPrometheusMonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ResetS3MonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3MonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

