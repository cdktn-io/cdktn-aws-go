package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlTranslation() AwsS3BucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference
	// Experimental.
	AccessControlTranslationInput() *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty
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
	EncryptionConfiguration() AwsS3BucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference
	// Experimental.
	EncryptionConfigurationInput() *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsS3BucketReplicationConfiguration_DestinationProperty
	// Experimental.
	SetInternalValue(val *AwsS3BucketReplicationConfiguration_DestinationProperty)
	// Experimental.
	Metrics() AwsS3BucketReplicationConfiguration_MetricsPropertyOutputReference
	// Experimental.
	MetricsInput() *AwsS3BucketReplicationConfiguration_MetricsProperty
	// Experimental.
	ReplicationTime() AwsS3BucketReplicationConfiguration_ReplicationTimePropertyOutputReference
	// Experimental.
	ReplicationTimeInput() *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty
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
	PutAccessControlTranslation(value *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty)
	// Experimental.
	PutEncryptionConfiguration(value *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty)
	// Experimental.
	PutMetrics(value *AwsS3BucketReplicationConfiguration_MetricsProperty)
	// Experimental.
	PutReplicationTime(value *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty)
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

// The jsii proxy struct for AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference
type jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) AccessControlTranslation() AwsS3BucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_AccessControlTranslationPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlTranslation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) AccessControlTranslationInput() *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty {
	var returns *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty
	_jsii_.Get(
		j,
		"accessControlTranslationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) EncryptionConfiguration() AwsS3BucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_EncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) EncryptionConfigurationInput() *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty {
	var returns *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) InternalValue() *AwsS3BucketReplicationConfiguration_DestinationProperty {
	var returns *AwsS3BucketReplicationConfiguration_DestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) Metrics() AwsS3BucketReplicationConfiguration_MetricsPropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_MetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) MetricsInput() *AwsS3BucketReplicationConfiguration_MetricsProperty {
	var returns *AwsS3BucketReplicationConfiguration_MetricsProperty
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ReplicationTime() AwsS3BucketReplicationConfiguration_ReplicationTimePropertyOutputReference {
	var returns AwsS3BucketReplicationConfiguration_ReplicationTimePropertyOutputReference
	_jsii_.Get(
		j,
		"replicationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ReplicationTimeInput() *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty {
	var returns *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty
	_jsii_.Get(
		j,
		"replicationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3BucketReplicationConfiguration_DestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketReplicationConfiguration.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference_Override(a AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketReplicationConfiguration.DestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetAccount(val *string) {
	if err := j.validateSetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetInternalValue(val *AwsS3BucketReplicationConfiguration_DestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) PutAccessControlTranslation(value *AwsS3BucketReplicationConfiguration_AccessControlTranslationProperty) {
	if err := a.validatePutAccessControlTranslationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlTranslation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) PutEncryptionConfiguration(value *AwsS3BucketReplicationConfiguration_EncryptionConfigurationProperty) {
	if err := a.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) PutMetrics(value *AwsS3BucketReplicationConfiguration_MetricsProperty) {
	if err := a.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) PutReplicationTime(value *AwsS3BucketReplicationConfiguration_ReplicationTimeProperty) {
	if err := a.validatePutReplicationTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplicationTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetAccessControlTranslation() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessControlTranslation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetReplicationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3BucketReplicationConfiguration_DestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

