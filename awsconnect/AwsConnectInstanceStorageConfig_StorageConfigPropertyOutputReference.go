package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConnectInstanceStorageConfig_StorageConfigProperty
	// Experimental.
	SetInternalValue(val *AwsConnectInstanceStorageConfig_StorageConfigProperty)
	// Experimental.
	KinesisFirehoseConfig() AwsConnectInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference
	// Experimental.
	KinesisFirehoseConfigInput() *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty
	// Experimental.
	KinesisStreamConfig() AwsConnectInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference
	// Experimental.
	KinesisStreamConfigInput() *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty
	// Experimental.
	KinesisVideoStreamConfig() AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference
	// Experimental.
	KinesisVideoStreamConfigInput() *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty
	// Experimental.
	S3Config() AwsConnectInstanceStorageConfig_S3ConfigPropertyOutputReference
	// Experimental.
	S3ConfigInput() *AwsConnectInstanceStorageConfig_S3ConfigProperty
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
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
	PutKinesisFirehoseConfig(value *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty)
	// Experimental.
	PutKinesisStreamConfig(value *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty)
	// Experimental.
	PutKinesisVideoStreamConfig(value *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty)
	// Experimental.
	PutS3Config(value *AwsConnectInstanceStorageConfig_S3ConfigProperty)
	// Experimental.
	ResetKinesisFirehoseConfig()
	// Experimental.
	ResetKinesisStreamConfig()
	// Experimental.
	ResetKinesisVideoStreamConfig()
	// Experimental.
	ResetS3Config()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference
type jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) InternalValue() *AwsConnectInstanceStorageConfig_StorageConfigProperty {
	var returns *AwsConnectInstanceStorageConfig_StorageConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisFirehoseConfig() AwsConnectInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference {
	var returns AwsConnectInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisFirehoseConfigInput() *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty {
	var returns *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisStreamConfig() AwsConnectInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference {
	var returns AwsConnectInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisStreamConfigInput() *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty {
	var returns *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty
	_jsii_.Get(
		j,
		"kinesisStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisVideoStreamConfig() AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference {
	var returns AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisVideoStreamConfigInput() *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty {
	var returns *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) S3Config() AwsConnectInstanceStorageConfig_S3ConfigPropertyOutputReference {
	var returns AwsConnectInstanceStorageConfig_S3ConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) S3ConfigInput() *AwsConnectInstanceStorageConfig_S3ConfigProperty {
	var returns *AwsConnectInstanceStorageConfig_S3ConfigProperty
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.AwsConnectInstanceStorageConfig.StorageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference_Override(a AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.AwsConnectInstanceStorageConfig.StorageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetInternalValue(val *AwsConnectInstanceStorageConfig_StorageConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisFirehoseConfig(value *AwsConnectInstanceStorageConfig_KinesisFirehoseConfigProperty) {
	if err := a.validatePutKinesisFirehoseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehoseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisStreamConfig(value *AwsConnectInstanceStorageConfig_KinesisStreamConfigProperty) {
	if err := a.validatePutKinesisStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisVideoStreamConfig(value *AwsConnectInstanceStorageConfig_KinesisVideoStreamConfigProperty) {
	if err := a.validatePutKinesisVideoStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisVideoStreamConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) PutS3Config(value *AwsConnectInstanceStorageConfig_S3ConfigProperty) {
	if err := a.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Config",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisFirehoseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehoseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisStreamConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisVideoStreamConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisVideoStreamConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetS3Config() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Config",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectInstanceStorageConfig_StorageConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

