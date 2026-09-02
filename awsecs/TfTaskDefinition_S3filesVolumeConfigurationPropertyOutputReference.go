package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessPointArn() *string
	// Experimental.
	SetAccessPointArn(val *string)
	// Experimental.
	AccessPointArnInput() *string
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
	FileSystemArn() *string
	// Experimental.
	SetFileSystemArn(val *string)
	// Experimental.
	FileSystemArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTaskDefinition_S3filesVolumeConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfTaskDefinition_S3filesVolumeConfigurationProperty)
	// Experimental.
	RootDirectory() *string
	// Experimental.
	SetRootDirectory(val *string)
	// Experimental.
	RootDirectoryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitEncryptionPort() *float64
	// Experimental.
	SetTransitEncryptionPort(val *float64)
	// Experimental.
	TransitEncryptionPortInput() *float64
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
	ResetAccessPointArn()
	// Experimental.
	ResetRootDirectory()
	// Experimental.
	ResetTransitEncryptionPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference
type jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) AccessPointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) FileSystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) FileSystemArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) InternalValue() *TfTaskDefinition_S3filesVolumeConfigurationProperty {
	var returns *TfTaskDefinition_S3filesVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) RootDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) RootDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) TransitEncryptionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) TransitEncryptionPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPortInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.S3filesVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference_Override(t TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.S3filesVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetAccessPointArn(val *string) {
	if err := j.validateSetAccessPointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPointArn",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetFileSystemArn(val *string) {
	if err := j.validateSetFileSystemArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemArn",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetInternalValue(val *TfTaskDefinition_S3filesVolumeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetRootDirectory(val *string) {
	if err := j.validateSetRootDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDirectory",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference)SetTransitEncryptionPort(val *float64) {
	if err := j.validateSetTransitEncryptionPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionPort",
		val,
	)
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ResetAccessPointArn() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessPointArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ResetRootDirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetRootDirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ResetTransitEncryptionPort() {
	_jsii_.InvokeVoid(
		t,
		"resetTransitEncryptionPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

