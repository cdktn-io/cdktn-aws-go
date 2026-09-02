package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInstanceStorageConfig_StorageConfigPropertyOutputReference interface {
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
	InternalValue() *TfInstanceStorageConfig_StorageConfigProperty
	// Experimental.
	SetInternalValue(val *TfInstanceStorageConfig_StorageConfigProperty)
	// Experimental.
	KinesisFirehoseConfig() TfInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference
	// Experimental.
	KinesisFirehoseConfigInput() *TfInstanceStorageConfig_KinesisFirehoseConfigProperty
	// Experimental.
	KinesisStreamConfig() TfInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference
	// Experimental.
	KinesisStreamConfigInput() *TfInstanceStorageConfig_KinesisStreamConfigProperty
	// Experimental.
	KinesisVideoStreamConfig() TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference
	// Experimental.
	KinesisVideoStreamConfigInput() *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty
	// Experimental.
	S3Config() TfInstanceStorageConfig_S3ConfigPropertyOutputReference
	// Experimental.
	S3ConfigInput() *TfInstanceStorageConfig_S3ConfigProperty
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
	PutKinesisFirehoseConfig(value *TfInstanceStorageConfig_KinesisFirehoseConfigProperty)
	// Experimental.
	PutKinesisStreamConfig(value *TfInstanceStorageConfig_KinesisStreamConfigProperty)
	// Experimental.
	PutKinesisVideoStreamConfig(value *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty)
	// Experimental.
	PutS3Config(value *TfInstanceStorageConfig_S3ConfigProperty)
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

// The jsii proxy struct for TfInstanceStorageConfig_StorageConfigPropertyOutputReference
type jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) InternalValue() *TfInstanceStorageConfig_StorageConfigProperty {
	var returns *TfInstanceStorageConfig_StorageConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisFirehoseConfig() TfInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference {
	var returns TfInstanceStorageConfig_KinesisFirehoseConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisFirehoseConfigInput() *TfInstanceStorageConfig_KinesisFirehoseConfigProperty {
	var returns *TfInstanceStorageConfig_KinesisFirehoseConfigProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisStreamConfig() TfInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference {
	var returns TfInstanceStorageConfig_KinesisStreamConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisStreamConfigInput() *TfInstanceStorageConfig_KinesisStreamConfigProperty {
	var returns *TfInstanceStorageConfig_KinesisStreamConfigProperty
	_jsii_.Get(
		j,
		"kinesisStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisVideoStreamConfig() TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference {
	var returns TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) KinesisVideoStreamConfigInput() *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty {
	var returns *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty
	_jsii_.Get(
		j,
		"kinesisVideoStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) S3Config() TfInstanceStorageConfig_S3ConfigPropertyOutputReference {
	var returns TfInstanceStorageConfig_S3ConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) S3ConfigInput() *TfInstanceStorageConfig_S3ConfigProperty {
	var returns *TfInstanceStorageConfig_S3ConfigProperty
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfInstanceStorageConfig_StorageConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfInstanceStorageConfig_StorageConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfInstanceStorageConfig_StorageConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstanceStorageConfig.StorageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfInstanceStorageConfig_StorageConfigPropertyOutputReference_Override(t TfInstanceStorageConfig_StorageConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstanceStorageConfig.StorageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetInternalValue(val *TfInstanceStorageConfig_StorageConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisFirehoseConfig(value *TfInstanceStorageConfig_KinesisFirehoseConfigProperty) {
	if err := t.validatePutKinesisFirehoseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisFirehoseConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisStreamConfig(value *TfInstanceStorageConfig_KinesisStreamConfigProperty) {
	if err := t.validatePutKinesisStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisStreamConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) PutKinesisVideoStreamConfig(value *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty) {
	if err := t.validatePutKinesisVideoStreamConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisVideoStreamConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) PutS3Config(value *TfInstanceStorageConfig_S3ConfigProperty) {
	if err := t.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Config",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisFirehoseConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisFirehoseConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisStreamConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisStreamConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetKinesisVideoStreamConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisVideoStreamConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ResetS3Config() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Config",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfInstanceStorageConfig_StorageConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

