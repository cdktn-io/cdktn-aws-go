package emrserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrserverless/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLoggingConfiguration() AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference
	// Experimental.
	CloudwatchLoggingConfigurationInput() *AwsApplication_CloudwatchLoggingConfigurationProperty
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
	InternalValue() *AwsApplication_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_MonitoringConfigurationProperty)
	// Experimental.
	ManagedPersistenceMonitoringConfiguration() AwsApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference
	// Experimental.
	ManagedPersistenceMonitoringConfigurationInput() *AwsApplication_ManagedPersistenceMonitoringConfigurationProperty
	// Experimental.
	PrometheusMonitoringConfiguration() AwsApplication_PrometheusMonitoringConfigurationPropertyOutputReference
	// Experimental.
	PrometheusMonitoringConfigurationInput() *AwsApplication_PrometheusMonitoringConfigurationProperty
	// Experimental.
	S3MonitoringConfiguration() AwsApplication_S3MonitoringConfigurationPropertyOutputReference
	// Experimental.
	S3MonitoringConfigurationInput() *AwsApplication_S3MonitoringConfigurationProperty
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
	PutCloudwatchLoggingConfiguration(value *AwsApplication_CloudwatchLoggingConfigurationProperty)
	// Experimental.
	PutManagedPersistenceMonitoringConfiguration(value *AwsApplication_ManagedPersistenceMonitoringConfigurationProperty)
	// Experimental.
	PutPrometheusMonitoringConfiguration(value *AwsApplication_PrometheusMonitoringConfigurationProperty)
	// Experimental.
	PutS3MonitoringConfiguration(value *AwsApplication_S3MonitoringConfigurationProperty)
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

// The jsii proxy struct for AwsApplication_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) CloudwatchLoggingConfiguration() AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference {
	var returns AwsApplication_CloudwatchLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) CloudwatchLoggingConfigurationInput() *AwsApplication_CloudwatchLoggingConfigurationProperty {
	var returns *AwsApplication_CloudwatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) InternalValue() *AwsApplication_MonitoringConfigurationProperty {
	var returns *AwsApplication_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ManagedPersistenceMonitoringConfiguration() AwsApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference {
	var returns AwsApplication_ManagedPersistenceMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ManagedPersistenceMonitoringConfigurationInput() *AwsApplication_ManagedPersistenceMonitoringConfigurationProperty {
	var returns *AwsApplication_ManagedPersistenceMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"managedPersistenceMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PrometheusMonitoringConfiguration() AwsApplication_PrometheusMonitoringConfigurationPropertyOutputReference {
	var returns AwsApplication_PrometheusMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"prometheusMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PrometheusMonitoringConfigurationInput() *AwsApplication_PrometheusMonitoringConfigurationProperty {
	var returns *AwsApplication_PrometheusMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"prometheusMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfiguration() AwsApplication_S3MonitoringConfigurationPropertyOutputReference {
	var returns AwsApplication_S3MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3MonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfigurationInput() *AwsApplication_S3MonitoringConfigurationProperty {
	var returns *AwsApplication_S3MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"s3MonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsApplication.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_MonitoringConfigurationPropertyOutputReference_Override(a AwsApplication_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsApplication.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *AwsApplication_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PutCloudwatchLoggingConfiguration(value *AwsApplication_CloudwatchLoggingConfigurationProperty) {
	if err := a.validatePutCloudwatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PutManagedPersistenceMonitoringConfiguration(value *AwsApplication_ManagedPersistenceMonitoringConfigurationProperty) {
	if err := a.validatePutManagedPersistenceMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedPersistenceMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PutPrometheusMonitoringConfiguration(value *AwsApplication_PrometheusMonitoringConfigurationProperty) {
	if err := a.validatePutPrometheusMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrometheusMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) PutS3MonitoringConfiguration(value *AwsApplication_S3MonitoringConfigurationProperty) {
	if err := a.validatePutS3MonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3MonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ResetCloudwatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ResetManagedPersistenceMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedPersistenceMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ResetPrometheusMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPrometheusMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ResetS3MonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3MonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

