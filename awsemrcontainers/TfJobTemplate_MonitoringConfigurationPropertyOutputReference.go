package awsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobTemplate_MonitoringConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchMonitoringConfiguration() TfJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference
	// Experimental.
	CloudWatchMonitoringConfigurationInput() *TfJobTemplate_CloudWatchMonitoringConfigurationProperty
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
	InternalValue() *TfJobTemplate_MonitoringConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfJobTemplate_MonitoringConfigurationProperty)
	// Experimental.
	PersistentAppUi() *string
	// Experimental.
	SetPersistentAppUi(val *string)
	// Experimental.
	PersistentAppUiInput() *string
	// Experimental.
	S3MonitoringConfiguration() TfJobTemplate_S3MonitoringConfigurationPropertyOutputReference
	// Experimental.
	S3MonitoringConfigurationInput() *TfJobTemplate_S3MonitoringConfigurationProperty
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
	PutCloudWatchMonitoringConfiguration(value *TfJobTemplate_CloudWatchMonitoringConfigurationProperty)
	// Experimental.
	PutS3MonitoringConfiguration(value *TfJobTemplate_S3MonitoringConfigurationProperty)
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

// The jsii proxy struct for TfJobTemplate_MonitoringConfigurationPropertyOutputReference
type jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) CloudWatchMonitoringConfiguration() TfJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference {
	var returns TfJobTemplate_CloudWatchMonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchMonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) CloudWatchMonitoringConfigurationInput() *TfJobTemplate_CloudWatchMonitoringConfigurationProperty {
	var returns *TfJobTemplate_CloudWatchMonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"cloudWatchMonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) InternalValue() *TfJobTemplate_MonitoringConfigurationProperty {
	var returns *TfJobTemplate_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) PersistentAppUi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentAppUi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) PersistentAppUiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentAppUiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfiguration() TfJobTemplate_S3MonitoringConfigurationPropertyOutputReference {
	var returns TfJobTemplate_S3MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3MonitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) S3MonitoringConfigurationInput() *TfJobTemplate_S3MonitoringConfigurationProperty {
	var returns *TfJobTemplate_S3MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"s3MonitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobTemplate_MonitoringConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobTemplate_MonitoringConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobTemplate_MonitoringConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobTemplate_MonitoringConfigurationPropertyOutputReference_Override(t TfJobTemplate_MonitoringConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.MonitoringConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetInternalValue(val *TfJobTemplate_MonitoringConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetPersistentAppUi(val *string) {
	if err := j.validateSetPersistentAppUiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentAppUi",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) PutCloudWatchMonitoringConfiguration(value *TfJobTemplate_CloudWatchMonitoringConfigurationProperty) {
	if err := t.validatePutCloudWatchMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudWatchMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) PutS3MonitoringConfiguration(value *TfJobTemplate_S3MonitoringConfigurationProperty) {
	if err := t.validatePutS3MonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3MonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetCloudWatchMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchMonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetPersistentAppUi() {
	_jsii_.InvokeVoid(
		t,
		"resetPersistentAppUi",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ResetS3MonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3MonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobTemplate_MonitoringConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

