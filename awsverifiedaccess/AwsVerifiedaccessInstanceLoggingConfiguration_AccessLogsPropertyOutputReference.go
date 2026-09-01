package awsverifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsProperty
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
	IncludeTrustContext() interface{}
	// Experimental.
	SetIncludeTrustContext(val interface{})
	// Experimental.
	IncludeTrustContextInput() interface{}
	// Experimental.
	InternalValue() *AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsProperty
	// Experimental.
	SetInternalValue(val *AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsProperty)
	// Experimental.
	KinesisDataFirehose() AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	// Experimental.
	KinesisDataFirehoseInput() *AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	// Experimental.
	LogVersion() *string
	// Experimental.
	SetLogVersion(val *string)
	// Experimental.
	LogVersionInput() *string
	// Experimental.
	S3() AwsVerifiedaccessInstanceLoggingConfiguration_S3PropertyOutputReference
	// Experimental.
	S3Input() *AwsVerifiedaccessInstanceLoggingConfiguration_S3Property
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
	PutCloudwatchLogs(value *AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsProperty)
	// Experimental.
	PutKinesisDataFirehose(value *AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehoseProperty)
	// Experimental.
	PutS3(value *AwsVerifiedaccessInstanceLoggingConfiguration_S3Property)
	// Experimental.
	ResetCloudwatchLogs()
	// Experimental.
	ResetIncludeTrustContext()
	// Experimental.
	ResetKinesisDataFirehose()
	// Experimental.
	ResetLogVersion()
	// Experimental.
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference
type jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogs() AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference {
	var returns AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogsInput() *AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsProperty {
	var returns *AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InternalValue() *AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsProperty {
	var returns *AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehose() AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference {
	var returns AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehoseInput() *AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehoseProperty {
	var returns *AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3() AwsVerifiedaccessInstanceLoggingConfiguration_S3PropertyOutputReference {
	var returns AwsVerifiedaccessInstanceLoggingConfiguration_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3Input() *AwsVerifiedaccessInstanceLoggingConfiguration_S3Property {
	var returns *AwsVerifiedaccessInstanceLoggingConfiguration_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsVerifiedaccessInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference_Override(a AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsVerifiedaccessInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetIncludeTrustContext(val interface{}) {
	if err := j.validateSetIncludeTrustContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTrustContext",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetInternalValue(val *AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetLogVersion(val *string) {
	if err := j.validateSetLogVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVersion",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutCloudwatchLogs(value *AwsVerifiedaccessInstanceLoggingConfiguration_CloudwatchLogsProperty) {
	if err := a.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutKinesisDataFirehose(value *AwsVerifiedaccessInstanceLoggingConfiguration_KinesisDataFirehoseProperty) {
	if err := a.validatePutKinesisDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisDataFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutS3(value *AwsVerifiedaccessInstanceLoggingConfiguration_S3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetIncludeTrustContext() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTrustContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetKinesisDataFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisDataFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetLogVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetLogVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVerifiedaccessInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

