package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsService_ManagedEbsVolumePropertyOutputReference interface {
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
	Encrypted() interface{}
	// Experimental.
	SetEncrypted(val interface{})
	// Experimental.
	EncryptedInput() interface{}
	// Experimental.
	FileSystemType() *string
	// Experimental.
	SetFileSystemType(val *string)
	// Experimental.
	FileSystemTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEcsService_ManagedEbsVolumeProperty
	// Experimental.
	SetInternalValue(val *AwsEcsService_ManagedEbsVolumeProperty)
	// Experimental.
	Iops() *float64
	// Experimental.
	SetIops(val *float64)
	// Experimental.
	IopsInput() *float64
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	SizeInGb() *float64
	// Experimental.
	SetSizeInGb(val *float64)
	// Experimental.
	SizeInGbInput() *float64
	// Experimental.
	SnapshotId() *string
	// Experimental.
	SetSnapshotId(val *string)
	// Experimental.
	SnapshotIdInput() *string
	// Experimental.
	TagSpecifications() AwsEcsService_TagSpecificationsPropertyList
	// Experimental.
	TagSpecificationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Throughput() *float64
	// Experimental.
	SetThroughput(val *float64)
	// Experimental.
	ThroughputInput() *float64
	// Experimental.
	VolumeInitializationRate() *float64
	// Experimental.
	SetVolumeInitializationRate(val *float64)
	// Experimental.
	VolumeInitializationRateInput() *float64
	// Experimental.
	VolumeType() *string
	// Experimental.
	SetVolumeType(val *string)
	// Experimental.
	VolumeTypeInput() *string
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
	PutTagSpecifications(value interface{})
	// Experimental.
	ResetEncrypted()
	// Experimental.
	ResetFileSystemType()
	// Experimental.
	ResetIops()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetSizeInGb()
	// Experimental.
	ResetSnapshotId()
	// Experimental.
	ResetTagSpecifications()
	// Experimental.
	ResetThroughput()
	// Experimental.
	ResetVolumeInitializationRate()
	// Experimental.
	ResetVolumeType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcsService_ManagedEbsVolumePropertyOutputReference
type jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) FileSystemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) FileSystemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) InternalValue() *AwsEcsService_ManagedEbsVolumeProperty {
	var returns *AwsEcsService_ManagedEbsVolumeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) SizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) SizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) TagSpecifications() AwsEcsService_TagSpecificationsPropertyList {
	var returns AwsEcsService_TagSpecificationsPropertyList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) VolumeInitializationRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeInitializationRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) VolumeInitializationRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeInitializationRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsService_ManagedEbsVolumePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcsService_ManagedEbsVolumePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsService_ManagedEbsVolumePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsService.ManagedEbsVolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsService_ManagedEbsVolumePropertyOutputReference_Override(a AwsEcsService_ManagedEbsVolumePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsService.ManagedEbsVolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetFileSystemType(val *string) {
	if err := j.validateSetFileSystemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemType",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetInternalValue(val *AwsEcsService_ManagedEbsVolumeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetSizeInGb(val *float64) {
	if err := j.validateSetSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInGb",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetVolumeInitializationRate(val *float64) {
	if err := j.validateSetVolumeInitializationRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeInitializationRate",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) PutTagSpecifications(value interface{}) {
	if err := a.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetFileSystemType() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		a,
		"resetIops",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetSizeInGb() {
	_jsii_.InvokeVoid(
		a,
		"resetSizeInGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		a,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetVolumeInitializationRate() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeInitializationRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsService_ManagedEbsVolumePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

