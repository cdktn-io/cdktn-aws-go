package awsverifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogs() TfInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *TfInstanceLoggingConfiguration_CloudwatchLogsProperty
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
	InternalValue() *TfInstanceLoggingConfiguration_AccessLogsProperty
	// Experimental.
	SetInternalValue(val *TfInstanceLoggingConfiguration_AccessLogsProperty)
	// Experimental.
	KinesisDataFirehose() TfInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	// Experimental.
	KinesisDataFirehoseInput() *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	// Experimental.
	LogVersion() *string
	// Experimental.
	SetLogVersion(val *string)
	// Experimental.
	LogVersionInput() *string
	// Experimental.
	S3() TfInstanceLoggingConfiguration_S3PropertyOutputReference
	// Experimental.
	S3Input() *TfInstanceLoggingConfiguration_S3Property
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
	PutCloudwatchLogs(value *TfInstanceLoggingConfiguration_CloudwatchLogsProperty)
	// Experimental.
	PutKinesisDataFirehose(value *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty)
	// Experimental.
	PutS3(value *TfInstanceLoggingConfiguration_S3Property)
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

// The jsii proxy struct for TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference
type jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogs() TfInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference {
	var returns TfInstanceLoggingConfiguration_CloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CloudwatchLogsInput() *TfInstanceLoggingConfiguration_CloudwatchLogsProperty {
	var returns *TfInstanceLoggingConfiguration_CloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) IncludeTrustContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InternalValue() *TfInstanceLoggingConfiguration_AccessLogsProperty {
	var returns *TfInstanceLoggingConfiguration_AccessLogsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehose() TfInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference {
	var returns TfInstanceLoggingConfiguration_KinesisDataFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) KinesisDataFirehoseInput() *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty {
	var returns *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) LogVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3() TfInstanceLoggingConfiguration_S3PropertyOutputReference {
	var returns TfInstanceLoggingConfiguration_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) S3Input() *TfInstanceLoggingConfiguration_S3Property {
	var returns *TfInstanceLoggingConfiguration_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfInstanceLoggingConfiguration_AccessLogsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.TfInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference_Override(t TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.TfInstanceLoggingConfiguration.AccessLogsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetIncludeTrustContext(val interface{}) {
	if err := j.validateSetIncludeTrustContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTrustContext",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetInternalValue(val *TfInstanceLoggingConfiguration_AccessLogsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetLogVersion(val *string) {
	if err := j.validateSetLogVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVersion",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutCloudwatchLogs(value *TfInstanceLoggingConfiguration_CloudwatchLogsProperty) {
	if err := t.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutKinesisDataFirehose(value *TfInstanceLoggingConfiguration_KinesisDataFirehoseProperty) {
	if err := t.validatePutKinesisDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisDataFirehose",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) PutS3(value *TfInstanceLoggingConfiguration_S3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetIncludeTrustContext() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeTrustContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetKinesisDataFirehose() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisDataFirehose",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetLogVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetLogVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfInstanceLoggingConfiguration_AccessLogsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

