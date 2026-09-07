package verifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/verifiedaccess/jsii"

	"github.com/cdktn-io/cdktn-aws-go/verifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() AwsInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty
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
	InternalValue() *AwsInstanceLoggingConfiguration_AccessLogsProperty
	// Experimental.
	SetInternalValue(val *AwsInstanceLoggingConfiguration_AccessLogsProperty)
	// Experimental.
	KinesisDataFirehose() AwsInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	// Experimental.
	KinesisDataFirehoseInput() *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	// Experimental.
	LogVersion() *string
	// Experimental.
	SetLogVersion(val *string)
	// Experimental.
	LogVersionInput() *string
	// Experimental.
	S3() AwsInstanceLoggingConfiguration_S3PropertyOutputReference
	// Experimental.
	S3Input() *AwsInstanceLoggingConfiguration_S3Property
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
	PutCloudwatchLogs(value *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty)
	// Experimental.
	PutKinesisDataFirehose(value *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty)
	// Experimental.
	PutS3(value *AwsInstanceLoggingConfiguration_S3Property)
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

// The jsii proxy struct for AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference
type jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogs() AwsInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference {
	var returns AwsInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogsInput() *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty {
	var returns *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InternalValue() *AwsInstanceLoggingConfiguration_AccessLogsProperty {
	var returns *AwsInstanceLoggingConfiguration_AccessLogsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehose() AwsInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference {
	var returns AwsInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehoseInput() *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty {
	var returns *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3() AwsInstanceLoggingConfiguration_S3PropertyOutputReference {
	var returns AwsInstanceLoggingConfiguration_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3Input() *AwsInstanceLoggingConfiguration_S3Property {
	var returns *AwsInstanceLoggingConfiguration_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference_Override(a AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetIncludeTrustContext(val interface{}) {
	if err := j.validateSetIncludeTrustContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTrustContext",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetInternalValue(val *AwsInstanceLoggingConfiguration_AccessLogsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetLogVersion(val *string) {
	if err := j.validateSetLogVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVersion",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutCloudwatchLogs(value *AwsInstanceLoggingConfiguration_CloudwatchLogsProperty) {
	if err := a.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutKinesisDataFirehose(value *AwsInstanceLoggingConfiguration_KinesisDataFirehoseProperty) {
	if err := a.validatePutKinesisDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisDataFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutS3(value *AwsInstanceLoggingConfiguration_S3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetIncludeTrustContext() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTrustContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetKinesisDataFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisDataFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetLogVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetLogVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

