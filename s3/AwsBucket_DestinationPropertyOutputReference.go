package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucket_DestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlTranslation() AwsBucket_AccessControlTranslationPropertyOutputReference
	// Experimental.
	AccessControlTranslationInput() *AwsBucket_AccessControlTranslationProperty
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
	InternalValue() *AwsBucket_DestinationProperty
	// Experimental.
	SetInternalValue(val *AwsBucket_DestinationProperty)
	// Experimental.
	Metrics() AwsBucket_MetricsPropertyOutputReference
	// Experimental.
	MetricsInput() *AwsBucket_MetricsProperty
	// Experimental.
	ReplicaKmsKeyId() *string
	// Experimental.
	SetReplicaKmsKeyId(val *string)
	// Experimental.
	ReplicaKmsKeyIdInput() *string
	// Experimental.
	ReplicationTime() AwsBucket_ReplicationTimePropertyOutputReference
	// Experimental.
	ReplicationTimeInput() *AwsBucket_ReplicationTimeProperty
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
	PutAccessControlTranslation(value *AwsBucket_AccessControlTranslationProperty)
	// Experimental.
	PutMetrics(value *AwsBucket_MetricsProperty)
	// Experimental.
	PutReplicationTime(value *AwsBucket_ReplicationTimeProperty)
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

// The jsii proxy struct for AwsBucket_DestinationPropertyOutputReference
type jsiiProxy_AwsBucket_DestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) AccessControlTranslation() AwsBucket_AccessControlTranslationPropertyOutputReference {
	var returns AwsBucket_AccessControlTranslationPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) AccessControlTranslationInput() *AwsBucket_AccessControlTranslationProperty {
	var returns *AwsBucket_AccessControlTranslationProperty
	_jsii_.Get(
		j,
		"accessControlTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) InternalValue() *AwsBucket_DestinationProperty {
	var returns *AwsBucket_DestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) Metrics() AwsBucket_MetricsPropertyOutputReference {
	var returns AwsBucket_MetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) MetricsInput() *AwsBucket_MetricsProperty {
	var returns *AwsBucket_MetricsProperty
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ReplicaKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ReplicaKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ReplicationTime() AwsBucket_ReplicationTimePropertyOutputReference {
	var returns AwsBucket_ReplicationTimePropertyOutputReference
	_jsii_.Get(
		j,
		"replicationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ReplicationTimeInput() *AwsBucket_ReplicationTimeProperty {
	var returns *AwsBucket_ReplicationTimeProperty
	_jsii_.Get(
		j,
		"replicationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucket_DestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBucket_DestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucket_DestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucket_DestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucket.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucket_DestinationPropertyOutputReference_Override(a AwsBucket_DestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucket.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetInternalValue(val *AwsBucket_DestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetReplicaKmsKeyId(val *string) {
	if err := j.validateSetReplicaKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucket_DestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) PutAccessControlTranslation(value *AwsBucket_AccessControlTranslationProperty) {
	if err := a.validatePutAccessControlTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlTranslation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) PutMetrics(value *AwsBucket_MetricsProperty) {
	if err := a.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) PutReplicationTime(value *AwsBucket_ReplicationTimeProperty) {
	if err := a.validatePutReplicationTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplicationTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetAccessControlTranslation() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessControlTranslation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetReplicaKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetReplicationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucket_DestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

