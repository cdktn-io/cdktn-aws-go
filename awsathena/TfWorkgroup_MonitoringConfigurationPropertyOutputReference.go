package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsathena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsathena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkgroup_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchLoggingConfiguration() TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference
	// Experimental.
	CloudWatchLoggingConfigurationInput() *TfWorkgroup_CloudWatchLoggingConfigurationProperty
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
	InternalValue() *TfWorkgroup_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfWorkgroup_MonitoringConfigurationProperty)
	// Experimental.
	ManagedLoggingConfiguration() TfWorkgroup_ManagedLoggingConfigurationPropertyOutputReference
	// Experimental.
	ManagedLoggingConfigurationInput() *TfWorkgroup_ManagedLoggingConfigurationProperty
	// Experimental.
	S3LoggingConfiguration() TfWorkgroup_S3LoggingConfigurationPropertyOutputReference
	// Experimental.
	S3LoggingConfigurationInput() *TfWorkgroup_S3LoggingConfigurationProperty
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
	PutCloudWatchLoggingConfiguration(value *TfWorkgroup_CloudWatchLoggingConfigurationProperty)
	// Experimental.
	PutManagedLoggingConfiguration(value *TfWorkgroup_ManagedLoggingConfigurationProperty)
	// Experimental.
	PutS3LoggingConfiguration(value *TfWorkgroup_S3LoggingConfigurationProperty)
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

// The jsii proxy struct for TfWorkgroup_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) CloudWatchLoggingConfiguration() TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference {
	var returns TfWorkgroup_CloudWatchLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) CloudWatchLoggingConfigurationInput() *TfWorkgroup_CloudWatchLoggingConfigurationProperty {
	var returns *TfWorkgroup_CloudWatchLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"cloudWatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) InternalValue() *TfWorkgroup_MonitoringConfigurationProperty {
	var returns *TfWorkgroup_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ManagedLoggingConfiguration() TfWorkgroup_ManagedLoggingConfigurationPropertyOutputReference {
	var returns TfWorkgroup_ManagedLoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ManagedLoggingConfigurationInput() *TfWorkgroup_ManagedLoggingConfigurationProperty {
	var returns *TfWorkgroup_ManagedLoggingConfigurationProperty
	_jsii_.Get(
		j,
		"managedLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) S3LoggingConfiguration() TfWorkgroup_S3LoggingConfigurationPropertyOutputReference {
	var returns TfWorkgroup_S3LoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3LoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) S3LoggingConfigurationInput() *TfWorkgroup_S3LoggingConfigurationProperty {
	var returns *TfWorkgroup_S3LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"s3LoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkgroup_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkgroup_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkgroup_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkgroup_MonitoringConfigurationPropertyOutputReference_Override(t TfWorkgroup_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *TfWorkgroup_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) PutCloudWatchLoggingConfiguration(value *TfWorkgroup_CloudWatchLoggingConfigurationProperty) {
	if err := t.validatePutCloudWatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudWatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) PutManagedLoggingConfiguration(value *TfWorkgroup_ManagedLoggingConfigurationProperty) {
	if err := t.validatePutManagedLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedLoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) PutS3LoggingConfiguration(value *TfWorkgroup_S3LoggingConfigurationProperty) {
	if err := t.validatePutS3LoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3LoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetCloudWatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetManagedLoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedLoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ResetS3LoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3LoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkgroup_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

