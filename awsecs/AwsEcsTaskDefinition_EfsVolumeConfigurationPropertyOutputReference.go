package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationConfig() AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigPropertyOutputReference
	// Experimental.
	AuthorizationConfigInput() *AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty
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
	FileSystemId() *string
	// Experimental.
	SetFileSystemId(val *string)
	// Experimental.
	FileSystemIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEcsTaskDefinition_EfsVolumeConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsEcsTaskDefinition_EfsVolumeConfigurationProperty)
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
	TransitEncryption() *string
	// Experimental.
	SetTransitEncryption(val *string)
	// Experimental.
	TransitEncryptionInput() *string
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
	PutAuthorizationConfig(value *AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty)
	// Experimental.
	ResetAuthorizationConfig()
	// Experimental.
	ResetRootDirectory()
	// Experimental.
	ResetTransitEncryption()
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

// The jsii proxy struct for AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference
type jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) AuthorizationConfig() AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigPropertyOutputReference {
	var returns AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"authorizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) AuthorizationConfigInput() *AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty {
	var returns *AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty
	_jsii_.Get(
		j,
		"authorizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) InternalValue() *AwsEcsTaskDefinition_EfsVolumeConfigurationProperty {
	var returns *AwsEcsTaskDefinition_EfsVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) RootDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) RootDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TransitEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TransitEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TransitEncryptionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) TransitEncryptionPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPortInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsTaskDefinition.EfsVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference_Override(a AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsTaskDefinition.EfsVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetInternalValue(val *AwsEcsTaskDefinition_EfsVolumeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetRootDirectory(val *string) {
	if err := j.validateSetRootDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDirectory",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetTransitEncryption(val *string) {
	if err := j.validateSetTransitEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference)SetTransitEncryptionPort(val *float64) {
	if err := j.validateSetTransitEncryptionPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionPort",
		val,
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) PutAuthorizationConfig(value *AwsEcsTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty) {
	if err := a.validatePutAuthorizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthorizationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ResetAuthorizationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ResetRootDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetRootDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ResetTransitEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ResetTransitEncryptionPort() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryptionPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

