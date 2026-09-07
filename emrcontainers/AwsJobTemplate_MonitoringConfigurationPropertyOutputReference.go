package emrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobTemplate_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchMonitoringConfiguration() AwsJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference
	// Experimental.
	CloudWatchMonitoringConfigurationInput() *AwsJobTemplate_CloudWatchMonitoringConfigurationProperty
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
	InternalValue() *AwsJobTemplate_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsJobTemplate_MonitoringConfigurationProperty)
	// Experimental.
	PersistentAppUi() *string
	// Experimental.
	SetPersistentAppUi(val *string)
	// Experimental.
	PersistentAppUiInput() *string
	// Experimental.
	S3MonitoringConfiguration() AwsJobTemplate_S3MonitoringConfigurationPropertyOutputReference
	// Experimental.
	S3MonitoringConfigurationInput() *AwsJobTemplate_S3MonitoringConfigurationProperty
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
	PutCloudWatchMonitoringConfiguration(value *AwsJobTemplate_CloudWatchMonitoringConfigurationProperty)
	// Experimental.
	PutS3MonitoringConfiguration(value *AwsJobTemplate_S3MonitoringConfigurationProperty)
	// Experimental.
	ResetCloudWatchMonitoringConfiguration()
	// Experimental.
	ResetPersistentAppUi()
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

// The jsii proxy struct for AwsJobTemplate_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) CloudWatchMonitoringConfiguration() AwsJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference {
	var returns AwsJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) CloudWatchMonitoringConfigurationInput() *AwsJobTemplate_CloudWatchMonitoringConfigurationProperty {
	var returns *AwsJobTemplate_CloudWatchMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"cloudWatchMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) InternalValue() *AwsJobTemplate_MonitoringConfigurationProperty {
	var returns *AwsJobTemplate_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) PersistentAppUi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentAppUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) PersistentAppUiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentAppUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfiguration() AwsJobTemplate_S3MonitoringConfigurationPropertyOutputReference {
	var returns AwsJobTemplate_S3MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3MonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfigurationInput() *AwsJobTemplate_S3MonitoringConfigurationProperty {
	var returns *AwsJobTemplate_S3MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"s3MonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobTemplate_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobTemplate_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobTemplate_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobTemplate_MonitoringConfigurationPropertyOutputReference_Override(a AwsJobTemplate_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *AwsJobTemplate_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetPersistentAppUi(val *string) {
	if err := j.validateSetPersistentAppUiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentAppUi",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) PutCloudWatchMonitoringConfiguration(value *AwsJobTemplate_CloudWatchMonitoringConfigurationProperty) {
	if err := a.validatePutCloudWatchMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) PutS3MonitoringConfiguration(value *AwsJobTemplate_S3MonitoringConfigurationProperty) {
	if err := a.validatePutS3MonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3MonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetCloudWatchMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetPersistentAppUi() {
	_jsii_.InvokeVoid(
		a,
		"resetPersistentAppUi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetS3MonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3MonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

