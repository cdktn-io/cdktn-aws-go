package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProject_EnvironmentPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Certificate() *string
	// Experimental.
	SetCertificate(val *string)
	// Experimental.
	CertificateInput() *string
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
	// Experimental.
	ComputeType() *string
	// Experimental.
	SetComputeType(val *string)
	// Experimental.
	ComputeTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DockerServer() TfProject_DockerServerPropertyOutputReference
	// Experimental.
	DockerServerInput() *TfProject_DockerServerProperty
	// Experimental.
	EnvironmentVariable() TfProject_EnvironmentVariablePropertyList
	// Experimental.
	EnvironmentVariableInput() interface{}
	// Experimental.
	Fleet() TfProject_FleetPropertyOutputReference
	// Experimental.
	FleetInput() *TfProject_FleetProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	HostKernel() *string
	// Experimental.
	SetHostKernel(val *string)
	// Experimental.
	HostKernelInput() *string
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageInput() *string
	// Experimental.
	ImagePullCredentialsType() *string
	// Experimental.
	SetImagePullCredentialsType(val *string)
	// Experimental.
	ImagePullCredentialsTypeInput() *string
	// Experimental.
	InternalValue() *TfProject_EnvironmentProperty
	// Experimental.
	SetInternalValue(val *TfProject_EnvironmentProperty)
	// Experimental.
	PrivilegedMode() interface{}
	// Experimental.
	SetPrivilegedMode(val interface{})
	// Experimental.
	PrivilegedModeInput() interface{}
	// Experimental.
	RegistryCredential() TfProject_RegistryCredentialPropertyOutputReference
	// Experimental.
	RegistryCredentialInput() *TfProject_RegistryCredentialProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutDockerServer(value *TfProject_DockerServerProperty)
	// Experimental.
	PutEnvironmentVariable(value interface{})
	// Experimental.
	PutFleet(value *TfProject_FleetProperty)
	// Experimental.
	PutRegistryCredential(value *TfProject_RegistryCredentialProperty)
	// Experimental.
	ResetCertificate()
	// Experimental.
	ResetDockerServer()
	// Experimental.
	ResetEnvironmentVariable()
	// Experimental.
	ResetFleet()
	// Experimental.
	ResetHostKernel()
	// Experimental.
	ResetImagePullCredentialsType()
	// Experimental.
	ResetPrivilegedMode()
	// Experimental.
	ResetRegistryCredential()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProject_EnvironmentPropertyOutputReference
type jsiiProxy_TfProject_EnvironmentPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) DockerServer() TfProject_DockerServerPropertyOutputReference {
	var returns TfProject_DockerServerPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) DockerServerInput() *TfProject_DockerServerProperty {
	var returns *TfProject_DockerServerProperty
	_jsii_.Get(
		j,
		"dockerServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) EnvironmentVariable() TfProject_EnvironmentVariablePropertyList {
	var returns TfProject_EnvironmentVariablePropertyList
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Fleet() TfProject_FleetPropertyOutputReference {
	var returns TfProject_FleetPropertyOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) FleetInput() *TfProject_FleetProperty {
	var returns *TfProject_FleetProperty
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) HostKernel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKernel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) HostKernelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKernelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ImagePullCredentialsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ImagePullCredentialsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) InternalValue() *TfProject_EnvironmentProperty {
	var returns *TfProject_EnvironmentProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PrivilegedMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PrivilegedModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) RegistryCredential() TfProject_RegistryCredentialPropertyOutputReference {
	var returns TfProject_RegistryCredentialPropertyOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) RegistryCredentialInput() *TfProject_RegistryCredentialProperty {
	var returns *TfProject_RegistryCredentialProperty
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProject_EnvironmentPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProject_EnvironmentPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProject_EnvironmentPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject_EnvironmentPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.EnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProject_EnvironmentPropertyOutputReference_Override(t TfProject_EnvironmentPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.EnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetComputeType(val *string) {
	if err := j.validateSetComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetHostKernel(val *string) {
	if err := j.validateSetHostKernelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKernel",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetImagePullCredentialsType(val *string) {
	if err := j.validateSetImagePullCredentialsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullCredentialsType",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetInternalValue(val *TfProject_EnvironmentProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetPrivilegedMode(val interface{}) {
	if err := j.validateSetPrivilegedModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedMode",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfProject_EnvironmentPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PutDockerServer(value *TfProject_DockerServerProperty) {
	if err := t.validatePutDockerServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDockerServer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PutEnvironmentVariable(value interface{}) {
	if err := t.validatePutEnvironmentVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironmentVariable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PutFleet(value *TfProject_FleetProperty) {
	if err := t.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFleet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) PutRegistryCredential(value *TfProject_RegistryCredentialProperty) {
	if err := t.validatePutRegistryCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetDockerServer() {
	_jsii_.InvokeVoid(
		t,
		"resetDockerServer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetFleet() {
	_jsii_.InvokeVoid(
		t,
		"resetFleet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetHostKernel() {
	_jsii_.InvokeVoid(
		t,
		"resetHostKernel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetImagePullCredentialsType() {
	_jsii_.InvokeVoid(
		t,
		"resetImagePullCredentialsType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetPrivilegedMode() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivilegedMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		t,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProject_EnvironmentPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

