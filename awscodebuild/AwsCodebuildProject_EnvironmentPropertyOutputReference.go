package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodebuildProject_EnvironmentPropertyOutputReference interface {
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
	DockerServer() AwsCodebuildProject_DockerServerPropertyOutputReference
	// Experimental.
	DockerServerInput() *AwsCodebuildProject_DockerServerProperty
	// Experimental.
	EnvironmentVariable() AwsCodebuildProject_EnvironmentVariablePropertyList
	// Experimental.
	EnvironmentVariableInput() interface{}
	// Experimental.
	Fleet() AwsCodebuildProject_FleetPropertyOutputReference
	// Experimental.
	FleetInput() *AwsCodebuildProject_FleetProperty
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
	InternalValue() *AwsCodebuildProject_EnvironmentProperty
	// Experimental.
	SetInternalValue(val *AwsCodebuildProject_EnvironmentProperty)
	// Experimental.
	PrivilegedMode() interface{}
	// Experimental.
	SetPrivilegedMode(val interface{})
	// Experimental.
	PrivilegedModeInput() interface{}
	// Experimental.
	RegistryCredential() AwsCodebuildProject_RegistryCredentialPropertyOutputReference
	// Experimental.
	RegistryCredentialInput() *AwsCodebuildProject_RegistryCredentialProperty
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
	PutDockerServer(value *AwsCodebuildProject_DockerServerProperty)
	// Experimental.
	PutEnvironmentVariable(value interface{})
	// Experimental.
	PutFleet(value *AwsCodebuildProject_FleetProperty)
	// Experimental.
	PutRegistryCredential(value *AwsCodebuildProject_RegistryCredentialProperty)
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

// The jsii proxy struct for AwsCodebuildProject_EnvironmentPropertyOutputReference
type jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) DockerServer() AwsCodebuildProject_DockerServerPropertyOutputReference {
	var returns AwsCodebuildProject_DockerServerPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) DockerServerInput() *AwsCodebuildProject_DockerServerProperty {
	var returns *AwsCodebuildProject_DockerServerProperty
	_jsii_.Get(
		j,
		"dockerServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) EnvironmentVariable() AwsCodebuildProject_EnvironmentVariablePropertyList {
	var returns AwsCodebuildProject_EnvironmentVariablePropertyList
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Fleet() AwsCodebuildProject_FleetPropertyOutputReference {
	var returns AwsCodebuildProject_FleetPropertyOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) FleetInput() *AwsCodebuildProject_FleetProperty {
	var returns *AwsCodebuildProject_FleetProperty
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) HostKernel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKernel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) HostKernelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKernelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ImagePullCredentialsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ImagePullCredentialsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) InternalValue() *AwsCodebuildProject_EnvironmentProperty {
	var returns *AwsCodebuildProject_EnvironmentProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PrivilegedMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PrivilegedModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) RegistryCredential() AwsCodebuildProject_RegistryCredentialPropertyOutputReference {
	var returns AwsCodebuildProject_RegistryCredentialPropertyOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) RegistryCredentialInput() *AwsCodebuildProject_RegistryCredentialProperty {
	var returns *AwsCodebuildProject_RegistryCredentialProperty
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodebuildProject_EnvironmentPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCodebuildProject_EnvironmentPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodebuildProject_EnvironmentPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsCodebuildProject.EnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodebuildProject_EnvironmentPropertyOutputReference_Override(a AwsCodebuildProject_EnvironmentPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsCodebuildProject.EnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetComputeType(val *string) {
	if err := j.validateSetComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetHostKernel(val *string) {
	if err := j.validateSetHostKernelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKernel",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetImagePullCredentialsType(val *string) {
	if err := j.validateSetImagePullCredentialsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullCredentialsType",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetInternalValue(val *AwsCodebuildProject_EnvironmentProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetPrivilegedMode(val interface{}) {
	if err := j.validateSetPrivilegedModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedMode",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PutDockerServer(value *AwsCodebuildProject_DockerServerProperty) {
	if err := a.validatePutDockerServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDockerServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PutEnvironmentVariable(value interface{}) {
	if err := a.validatePutEnvironmentVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentVariable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PutFleet(value *AwsCodebuildProject_FleetProperty) {
	if err := a.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFleet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) PutRegistryCredential(value *AwsCodebuildProject_RegistryCredentialProperty) {
	if err := a.validatePutRegistryCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetDockerServer() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetFleet() {
	_jsii_.InvokeVoid(
		a,
		"resetFleet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetHostKernel() {
	_jsii_.InvokeVoid(
		a,
		"resetHostKernel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetImagePullCredentialsType() {
	_jsii_.InvokeVoid(
		a,
		"resetImagePullCredentialsType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetPrivilegedMode() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivilegedMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodebuildProject_EnvironmentPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

