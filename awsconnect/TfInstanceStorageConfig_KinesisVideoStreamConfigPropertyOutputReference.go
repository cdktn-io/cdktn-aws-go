package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference interface {
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
	EncryptionConfig() TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigPropertyOutputReference
	// Experimental.
	EncryptionConfigInput() *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty
	// Experimental.
	SetInternalValue(val *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty)
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	RetentionPeriodHours() *float64
	// Experimental.
	SetRetentionPeriodHours(val *float64)
	// Experimental.
	RetentionPeriodHoursInput() *float64
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
	PutEncryptionConfig(value *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference
type jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) EncryptionConfig() TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigPropertyOutputReference {
	var returns TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) EncryptionConfigInput() *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty {
	var returns *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) InternalValue() *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty {
	var returns *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) RetentionPeriodHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) RetentionPeriodHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstanceStorageConfig.KinesisVideoStreamConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference_Override(t TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstanceStorageConfig.KinesisVideoStreamConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetInternalValue(val *TfInstanceStorageConfig_KinesisVideoStreamConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetRetentionPeriodHours(val *float64) {
	if err := j.validateSetRetentionPeriodHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodHours",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) PutEncryptionConfig(value *TfInstanceStorageConfig_StorageConfigKinesisVideoStreamConfigEncryptionConfigProperty) {
	if err := t.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfInstanceStorageConfig_KinesisVideoStreamConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

