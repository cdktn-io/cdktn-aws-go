package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference interface {
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
	InternalValue() *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsProperty)
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

// The jsii proxy struct for AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference
type jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InternalValue() *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) LogStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference_Override(a AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetLogStreamName(val *string) {
	if err := j.validateSetLogStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStreamName",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ResetLogStreamName() {
	_jsii_.InvokeVoid(
		a,
		"resetLogStreamName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

