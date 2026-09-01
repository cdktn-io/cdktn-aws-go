package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsathena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsathena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchLoggingConfiguration() AwsAthenaWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference
	// Experimental.
	CloudWatchLoggingConfigurationInput() *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty
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
	InternalValue() *AwsAthenaWorkgroup_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsAthenaWorkgroup_MonitoringConfigurationProperty)
	// Experimental.
	ManagedLoggingConfiguration() AwsAthenaWorkgroup_ManagedLoggingConfigurationPropertyOutputReference
	// Experimental.
	ManagedLoggingConfigurationInput() *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty
	// Experimental.
	S3LoggingConfiguration() AwsAthenaWorkgroup_S3LoggingConfigurationPropertyOutputReference
	// Experimental.
	S3LoggingConfigurationInput() *AwsAthenaWorkgroup_S3LoggingConfigurationProperty
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
	PutCloudWatchLoggingConfiguration(value *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty)
	// Experimental.
	PutManagedLoggingConfiguration(value *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty)
	// Experimental.
	PutS3LoggingConfiguration(value *AwsAthenaWorkgroup_S3LoggingConfigurationProperty)
	// Experimental.
	ResetCloudWatchLoggingConfiguration()
	// Experimental.
	ResetManagedLoggingConfiguration()
	// Experimental.
	ResetS3LoggingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) CloudWatchLoggingConfiguration() AwsAthenaWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) CloudWatchLoggingConfigurationInput() *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty {
	var returns *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) InternalValue() *AwsAthenaWorkgroup_MonitoringConfigurationProperty {
	var returns *AwsAthenaWorkgroup_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ManagedLoggingConfiguration() AwsAthenaWorkgroup_ManagedLoggingConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_ManagedLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ManagedLoggingConfigurationInput() *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty {
	var returns *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"managedLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) S3LoggingConfiguration() AwsAthenaWorkgroup_S3LoggingConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_S3LoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3LoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) S3LoggingConfigurationInput() *AwsAthenaWorkgroup_S3LoggingConfigurationProperty {
	var returns *AwsAthenaWorkgroup_S3LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"s3LoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.AwsAthenaWorkgroup.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference_Override(a AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.AwsAthenaWorkgroup.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *AwsAthenaWorkgroup_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) PutCloudWatchLoggingConfiguration(value *AwsAthenaWorkgroup_CloudWatchLoggingConfigurationProperty) {
	if err := a.validatePutCloudWatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) PutManagedLoggingConfiguration(value *AwsAthenaWorkgroup_ManagedLoggingConfigurationProperty) {
	if err := a.validatePutManagedLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) PutS3LoggingConfiguration(value *AwsAthenaWorkgroup_S3LoggingConfigurationProperty) {
	if err := a.validatePutS3LoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3LoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetCloudWatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetManagedLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetS3LoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3LoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

