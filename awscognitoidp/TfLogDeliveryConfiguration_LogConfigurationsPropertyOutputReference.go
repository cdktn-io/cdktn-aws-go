package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchLogsConfiguration() TfLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList
	// Experimental.
	CloudWatchLogsConfigurationInput() interface{}
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
	EventSource() *string
	// Experimental.
	SetEventSource(val *string)
	// Experimental.
	EventSourceInput() *string
	// Experimental.
	FirehoseConfiguration() TfLogDeliveryConfiguration_FirehoseConfigurationPropertyList
	// Experimental.
	FirehoseConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogLevel() *string
	// Experimental.
	SetLogLevel(val *string)
	// Experimental.
	LogLevelInput() *string
	// Experimental.
	S3Configuration() TfLogDeliveryConfiguration_S3ConfigurationPropertyList
	// Experimental.
	S3ConfigurationInput() interface{}
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
	PutCloudWatchLogsConfiguration(value interface{})
	// Experimental.
	PutFirehoseConfiguration(value interface{})
	// Experimental.
	PutS3Configuration(value interface{})
	// Experimental.
	ResetCloudWatchLogsConfiguration()
	// Experimental.
	ResetFirehoseConfiguration()
	// Experimental.
	ResetS3Configuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference
type jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CloudWatchLogsConfiguration() TfLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList {
	var returns TfLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList
	_jsii_.Get(
		j,
		"cloudWatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CloudWatchLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) EventSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) EventSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) FirehoseConfiguration() TfLogDeliveryConfiguration_FirehoseConfigurationPropertyList {
	var returns TfLogDeliveryConfiguration_FirehoseConfigurationPropertyList
	_jsii_.Get(
		j,
		"firehoseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) FirehoseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) S3Configuration() TfLogDeliveryConfiguration_S3ConfigurationPropertyList {
	var returns TfLogDeliveryConfiguration_S3ConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfLogDeliveryConfiguration.LogConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference_Override(t TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfLogDeliveryConfiguration.LogConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetEventSource(val *string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutCloudWatchLogsConfiguration(value interface{}) {
	if err := t.validatePutCloudWatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudWatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutFirehoseConfiguration(value interface{}) {
	if err := t.validatePutFirehoseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirehoseConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutS3Configuration(value interface{}) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetCloudWatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchLogsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetFirehoseConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehoseConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

