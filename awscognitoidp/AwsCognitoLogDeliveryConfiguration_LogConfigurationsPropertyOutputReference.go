package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchLogsConfiguration() AwsCognitoLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList
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
	FirehoseConfiguration() AwsCognitoLogDeliveryConfiguration_FirehoseConfigurationPropertyList
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
	S3Configuration() AwsCognitoLogDeliveryConfiguration_S3ConfigurationPropertyList
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

// The jsii proxy struct for AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference
type jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CloudWatchLogsConfiguration() AwsCognitoLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList {
	var returns AwsCognitoLogDeliveryConfiguration_CloudWatchLogsConfigurationPropertyList
	_jsii_.Get(
		j,
		"cloudWatchLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CloudWatchLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) EventSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) EventSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) FirehoseConfiguration() AwsCognitoLogDeliveryConfiguration_FirehoseConfigurationPropertyList {
	var returns AwsCognitoLogDeliveryConfiguration_FirehoseConfigurationPropertyList
	_jsii_.Get(
		j,
		"firehoseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) FirehoseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) S3Configuration() AwsCognitoLogDeliveryConfiguration_S3ConfigurationPropertyList {
	var returns AwsCognitoLogDeliveryConfiguration_S3ConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoLogDeliveryConfiguration.LogConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference_Override(a AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoLogDeliveryConfiguration.LogConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetEventSource(val *string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutCloudWatchLogsConfiguration(value interface{}) {
	if err := a.validatePutCloudWatchLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchLogsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutFirehoseConfiguration(value interface{}) {
	if err := a.validatePutFirehoseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehoseConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) PutS3Configuration(value interface{}) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetCloudWatchLogsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchLogsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetFirehoseConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoLogDeliveryConfiguration_LogConfigurationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

