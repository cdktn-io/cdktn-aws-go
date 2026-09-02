package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	LogGroupName() *string
	// Experimental.
	SetLogGroupName(val *string)
	// Experimental.
	LogGroupNameInput() *string
	// Experimental.
	LogStreamName() *string
	// Experimental.
	SetLogStreamName(val *string)
	// Experimental.
	LogStreamNameInput() *string
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
	ResetEnabled()
	// Experimental.
	ResetLogGroupName()
	// Experimental.
	ResetLogStreamName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference
type jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InternalValue() *TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference_Override(t TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetLogStreamName(val *string) {
	if err := j.validateSetLogStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetLogStreamName() {
	_jsii_.InvokeVoid(
		t,
		"resetLogStreamName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchserverlessConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

