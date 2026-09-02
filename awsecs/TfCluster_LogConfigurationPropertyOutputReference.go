package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_LogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchEncryptionEnabled() interface{}
	// Experimental.
	SetCloudWatchEncryptionEnabled(val interface{})
	// Experimental.
	CloudWatchEncryptionEnabledInput() interface{}
	// Experimental.
	CloudWatchLogGroupName() *string
	// Experimental.
	SetCloudWatchLogGroupName(val *string)
	// Experimental.
	CloudWatchLogGroupNameInput() *string
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
	InternalValue() *TfCluster_LogConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfCluster_LogConfigurationProperty)
	// Experimental.
	S3BucketEncryptionEnabled() interface{}
	// Experimental.
	SetS3BucketEncryptionEnabled(val interface{})
	// Experimental.
	S3BucketEncryptionEnabledInput() interface{}
	// Experimental.
	S3BucketName() *string
	// Experimental.
	SetS3BucketName(val *string)
	// Experimental.
	S3BucketNameInput() *string
	// Experimental.
	S3KeyPrefix() *string
	// Experimental.
	SetS3KeyPrefix(val *string)
	// Experimental.
	S3KeyPrefixInput() *string
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
	ResetCloudWatchEncryptionEnabled()
	// Experimental.
	ResetCloudWatchLogGroupName()
	// Experimental.
	ResetS3BucketEncryptionEnabled()
	// Experimental.
	ResetS3BucketName()
	// Experimental.
	ResetS3KeyPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCluster_LogConfigurationPropertyOutputReference
type jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) CloudWatchEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) CloudWatchEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) CloudWatchLogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) CloudWatchLogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) InternalValue() *TfCluster_LogConfigurationProperty {
	var returns *TfCluster_LogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3BucketEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3BucketEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_LogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_LogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_LogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCluster.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_LogConfigurationPropertyOutputReference_Override(t TfCluster_LogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCluster.LogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetCloudWatchEncryptionEnabled(val interface{}) {
	if err := j.validateSetCloudWatchEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetCloudWatchLogGroupName(val *string) {
	if err := j.validateSetCloudWatchLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchLogGroupName",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetInternalValue(val *TfCluster_LogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetS3BucketEncryptionEnabled(val interface{}) {
	if err := j.validateSetS3BucketEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetS3KeyPrefix(val *string) {
	if err := j.validateSetS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ResetCloudWatchEncryptionEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchEncryptionEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ResetCloudWatchLogGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchLogGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ResetS3BucketEncryptionEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BucketEncryptionEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_LogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

