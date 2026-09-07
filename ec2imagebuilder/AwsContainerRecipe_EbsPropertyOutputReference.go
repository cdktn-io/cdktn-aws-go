package ec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsContainerRecipe_EbsPropertyOutputReference interface {
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
	DeleteOnTermination() *string
	// Experimental.
	SetDeleteOnTermination(val *string)
	// Experimental.
	DeleteOnTerminationInput() *string
	// Experimental.
	Encrypted() *string
	// Experimental.
	SetEncrypted(val *string)
	// Experimental.
	EncryptedInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsContainerRecipe_EbsProperty
	// Experimental.
	SetInternalValue(val *AwsContainerRecipe_EbsProperty)
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
	SnapshotId() *string
	// Experimental.
	SetSnapshotId(val *string)
	// Experimental.
	SnapshotIdInput() *string
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
	VolumeSize() *float64
	// Experimental.
	SetVolumeSize(val *float64)
	// Experimental.
	VolumeSizeInput() *float64
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
	ResetDeleteOnTermination()
	// Experimental.
	ResetEncrypted()
	// Experimental.
	ResetIops()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetSnapshotId()
	// Experimental.
	ResetThroughput()
	// Experimental.
	ResetVolumeSize()
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

// The jsii proxy struct for AwsContainerRecipe_EbsPropertyOutputReference
type jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) DeleteOnTerminationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) Encrypted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) EncryptedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) InternalValue() *AwsContainerRecipe_EbsProperty {
	var returns *AwsContainerRecipe_EbsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsContainerRecipe_EbsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsContainerRecipe_EbsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsContainerRecipe_EbsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsContainerRecipe.EbsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsContainerRecipe_EbsPropertyOutputReference_Override(a AwsContainerRecipe_EbsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsContainerRecipe.EbsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetDeleteOnTermination(val *string) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetEncrypted(val *string) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetInternalValue(val *AwsContainerRecipe_EbsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		a,
		"resetIops",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsContainerRecipe_EbsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

