package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketReplicationConfiguration_DestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlTranslation() TfBucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference
	// Experimental.
	AccessControlTranslationInput() *TfBucketReplicationConfiguration_AccessControlTranslationProperty
	// Experimental.
	Account() *string
	// Experimental.
	SetAccount(val *string)
	// Experimental.
	AccountInput() *string
	// Experimental.
	Bucket() *string
	// Experimental.
	SetBucket(val *string)
	// Experimental.
	BucketInput() *string
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
	EncryptionConfiguration() TfBucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference
	// Experimental.
	EncryptionConfigurationInput() *TfBucketReplicationConfiguration_EncryptionConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfBucketReplicationConfiguration_DestinationProperty
	// Experimental.
	SetInternalValue(val *TfBucketReplicationConfiguration_DestinationProperty)
	// Experimental.
	Metrics() TfBucketReplicationConfiguration_MetricsPropertyOutputReference
	// Experimental.
	MetricsInput() *TfBucketReplicationConfiguration_MetricsProperty
	// Experimental.
	ReplicationTime() TfBucketReplicationConfiguration_ReplicationTimePropertyOutputReference
	// Experimental.
	ReplicationTimeInput() *TfBucketReplicationConfiguration_ReplicationTimeProperty
	// Experimental.
	StorageClass() *string
	// Experimental.
	SetStorageClass(val *string)
	// Experimental.
	StorageClassInput() *string
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
	PutAccessControlTranslation(value *TfBucketReplicationConfiguration_AccessControlTranslationProperty)
	// Experimental.
	PutEncryptionConfiguration(value *TfBucketReplicationConfiguration_EncryptionConfigurationProperty)
	// Experimental.
	PutMetrics(value *TfBucketReplicationConfiguration_MetricsProperty)
	// Experimental.
	PutReplicationTime(value *TfBucketReplicationConfiguration_ReplicationTimeProperty)
	// Experimental.
	ResetAccessControlTranslation()
	// Experimental.
	ResetAccount()
	// Experimental.
	ResetEncryptionConfiguration()
	// Experimental.
	ResetMetrics()
	// Experimental.
	ResetReplicationTime()
	// Experimental.
	ResetStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBucketReplicationConfiguration_DestinationPropertyOutputReference
type jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) AccessControlTranslation() TfBucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) AccessControlTranslationInput() *TfBucketReplicationConfiguration_AccessControlTranslationProperty {
	var returns *TfBucketReplicationConfiguration_AccessControlTranslationProperty
	_jsii_.Get(
		j,
		"accessControlTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) EncryptionConfiguration() TfBucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) EncryptionConfigurationInput() *TfBucketReplicationConfiguration_EncryptionConfigurationProperty {
	var returns *TfBucketReplicationConfiguration_EncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) InternalValue() *TfBucketReplicationConfiguration_DestinationProperty {
	var returns *TfBucketReplicationConfiguration_DestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) Metrics() TfBucketReplicationConfiguration_MetricsPropertyOutputReference {
	var returns TfBucketReplicationConfiguration_MetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) MetricsInput() *TfBucketReplicationConfiguration_MetricsProperty {
	var returns *TfBucketReplicationConfiguration_MetricsProperty
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ReplicationTime() TfBucketReplicationConfiguration_ReplicationTimePropertyOutputReference {
	var returns TfBucketReplicationConfiguration_ReplicationTimePropertyOutputReference
	_jsii_.Get(
		j,
		"replicationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ReplicationTimeInput() *TfBucketReplicationConfiguration_ReplicationTimeProperty {
	var returns *TfBucketReplicationConfiguration_ReplicationTimeProperty
	_jsii_.Get(
		j,
		"replicationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketReplicationConfiguration_DestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBucketReplicationConfiguration_DestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketReplicationConfiguration_DestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketReplicationConfiguration_DestinationPropertyOutputReference_Override(t TfBucketReplicationConfiguration_DestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketReplicationConfiguration.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetAccount(val *string) {
	if err := j.validateSetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetInternalValue(val *TfBucketReplicationConfiguration_DestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) PutAccessControlTranslation(value *TfBucketReplicationConfiguration_AccessControlTranslationProperty) {
	if err := t.validatePutAccessControlTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlTranslation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) PutEncryptionConfiguration(value *TfBucketReplicationConfiguration_EncryptionConfigurationProperty) {
	if err := t.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) PutMetrics(value *TfBucketReplicationConfiguration_MetricsProperty) {
	if err := t.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) PutReplicationTime(value *TfBucketReplicationConfiguration_ReplicationTimeProperty) {
	if err := t.validatePutReplicationTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReplicationTime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetAccessControlTranslation() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessControlTranslation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		t,
		"resetAccount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetReplicationTime() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicationTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketReplicationConfiguration_DestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

