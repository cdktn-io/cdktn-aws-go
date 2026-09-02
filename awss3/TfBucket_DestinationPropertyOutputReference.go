package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucket_DestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlTranslation() TfBucket_AccessControlTranslationPropertyOutputReference
	// Experimental.
	AccessControlTranslationInput() *TfBucket_AccessControlTranslationProperty
	// Experimental.
	AccountId() *string
	// Experimental.
	SetAccountId(val *string)
	// Experimental.
	AccountIdInput() *string
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfBucket_DestinationProperty
	// Experimental.
	SetInternalValue(val *TfBucket_DestinationProperty)
	// Experimental.
	Metrics() TfBucket_MetricsPropertyOutputReference
	// Experimental.
	MetricsInput() *TfBucket_MetricsProperty
	// Experimental.
	ReplicaKmsKeyId() *string
	// Experimental.
	SetReplicaKmsKeyId(val *string)
	// Experimental.
	ReplicaKmsKeyIdInput() *string
	// Experimental.
	ReplicationTime() TfBucket_ReplicationTimePropertyOutputReference
	// Experimental.
	ReplicationTimeInput() *TfBucket_ReplicationTimeProperty
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
	PutAccessControlTranslation(value *TfBucket_AccessControlTranslationProperty)
	// Experimental.
	PutMetrics(value *TfBucket_MetricsProperty)
	// Experimental.
	PutReplicationTime(value *TfBucket_ReplicationTimeProperty)
	// Experimental.
	ResetAccessControlTranslation()
	// Experimental.
	ResetAccountId()
	// Experimental.
	ResetMetrics()
	// Experimental.
	ResetReplicaKmsKeyId()
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

// The jsii proxy struct for TfBucket_DestinationPropertyOutputReference
type jsiiProxy_TfBucket_DestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) AccessControlTranslation() TfBucket_AccessControlTranslationPropertyOutputReference {
	var returns TfBucket_AccessControlTranslationPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) AccessControlTranslationInput() *TfBucket_AccessControlTranslationProperty {
	var returns *TfBucket_AccessControlTranslationProperty
	_jsii_.Get(
		j,
		"accessControlTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) InternalValue() *TfBucket_DestinationProperty {
	var returns *TfBucket_DestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) Metrics() TfBucket_MetricsPropertyOutputReference {
	var returns TfBucket_MetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) MetricsInput() *TfBucket_MetricsProperty {
	var returns *TfBucket_MetricsProperty
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ReplicaKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ReplicaKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ReplicationTime() TfBucket_ReplicationTimePropertyOutputReference {
	var returns TfBucket_ReplicationTimePropertyOutputReference
	_jsii_.Get(
		j,
		"replicationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ReplicationTimeInput() *TfBucket_ReplicationTimeProperty {
	var returns *TfBucket_ReplicationTimeProperty
	_jsii_.Get(
		j,
		"replicationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucket_DestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBucket_DestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucket_DestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucket_DestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucket.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucket_DestinationPropertyOutputReference_Override(t TfBucket_DestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucket.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetInternalValue(val *TfBucket_DestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetReplicaKmsKeyId(val *string) {
	if err := j.validateSetReplicaKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucket_DestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) PutAccessControlTranslation(value *TfBucket_AccessControlTranslationProperty) {
	if err := t.validatePutAccessControlTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlTranslation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) PutMetrics(value *TfBucket_MetricsProperty) {
	if err := t.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) PutReplicationTime(value *TfBucket_ReplicationTimeProperty) {
	if err := t.validatePutReplicationTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReplicationTime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetAccessControlTranslation() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessControlTranslation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetReplicaKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetReplicationTime() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicationTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucket_DestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

