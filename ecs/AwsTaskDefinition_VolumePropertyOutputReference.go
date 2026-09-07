package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTaskDefinition_VolumePropertyOutputReference interface {
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
	// Experimental.
	ConfigureAtLaunch() interface{}
	// Experimental.
	SetConfigureAtLaunch(val interface{})
	// Experimental.
	ConfigureAtLaunchInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DockerVolumeConfiguration() AwsTaskDefinition_DockerVolumeConfigurationPropertyOutputReference
	// Experimental.
	DockerVolumeConfigurationInput() *AwsTaskDefinition_DockerVolumeConfigurationProperty
	// Experimental.
	EfsVolumeConfiguration() AwsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference
	// Experimental.
	EfsVolumeConfigurationInput() *AwsTaskDefinition_EfsVolumeConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FsxWindowsFileServerVolumeConfiguration() AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference
	// Experimental.
	FsxWindowsFileServerVolumeConfigurationInput() *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty
	// Experimental.
	HostPath() *string
	// Experimental.
	SetHostPath(val *string)
	// Experimental.
	HostPathInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	S3FilesVolumeConfiguration() AwsTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference
	// Experimental.
	S3FilesVolumeConfigurationInput() *AwsTaskDefinition_S3filesVolumeConfigurationProperty
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
	PutDockerVolumeConfiguration(value *AwsTaskDefinition_DockerVolumeConfigurationProperty)
	// Experimental.
	PutEfsVolumeConfiguration(value *AwsTaskDefinition_EfsVolumeConfigurationProperty)
	// Experimental.
	PutFsxWindowsFileServerVolumeConfiguration(value *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty)
	// Experimental.
	PutS3FilesVolumeConfiguration(value *AwsTaskDefinition_S3filesVolumeConfigurationProperty)
	// Experimental.
	ResetConfigureAtLaunch()
	// Experimental.
	ResetDockerVolumeConfiguration()
	// Experimental.
	ResetEfsVolumeConfiguration()
	// Experimental.
	ResetFsxWindowsFileServerVolumeConfiguration()
	// Experimental.
	ResetHostPath()
	// Experimental.
	ResetS3FilesVolumeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTaskDefinition_VolumePropertyOutputReference
type jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ConfigureAtLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configureAtLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ConfigureAtLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configureAtLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) DockerVolumeConfiguration() AwsTaskDefinition_DockerVolumeConfigurationPropertyOutputReference {
	var returns AwsTaskDefinition_DockerVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) DockerVolumeConfigurationInput() *AwsTaskDefinition_DockerVolumeConfigurationProperty {
	var returns *AwsTaskDefinition_DockerVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"dockerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) EfsVolumeConfiguration() AwsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference {
	var returns AwsTaskDefinition_EfsVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"efsVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) EfsVolumeConfigurationInput() *AwsTaskDefinition_EfsVolumeConfigurationProperty {
	var returns *AwsTaskDefinition_EfsVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"efsVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) FsxWindowsFileServerVolumeConfiguration() AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference {
	var returns AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) FsxWindowsFileServerVolumeConfigurationInput() *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty {
	var returns *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) HostPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) HostPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) S3FilesVolumeConfiguration() AwsTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference {
	var returns AwsTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3FilesVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) S3FilesVolumeConfigurationInput() *AwsTaskDefinition_S3filesVolumeConfigurationProperty {
	var returns *AwsTaskDefinition_S3filesVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"s3FilesVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTaskDefinition_VolumePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTaskDefinition_VolumePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTaskDefinition_VolumePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsTaskDefinition.VolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTaskDefinition_VolumePropertyOutputReference_Override(a AwsTaskDefinition_VolumePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsTaskDefinition.VolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetConfigureAtLaunch(val interface{}) {
	if err := j.validateSetConfigureAtLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configureAtLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetHostPath(val *string) {
	if err := j.validateSetHostPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPath",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) PutDockerVolumeConfiguration(value *AwsTaskDefinition_DockerVolumeConfigurationProperty) {
	if err := a.validatePutDockerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDockerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) PutEfsVolumeConfiguration(value *AwsTaskDefinition_EfsVolumeConfigurationProperty) {
	if err := a.validatePutEfsVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEfsVolumeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) PutFsxWindowsFileServerVolumeConfiguration(value *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty) {
	if err := a.validatePutFsxWindowsFileServerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFsxWindowsFileServerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) PutS3FilesVolumeConfiguration(value *AwsTaskDefinition_S3filesVolumeConfigurationProperty) {
	if err := a.validatePutS3FilesVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3FilesVolumeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetConfigureAtLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigureAtLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetDockerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerVolumeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetEfsVolumeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEfsVolumeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetFsxWindowsFileServerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFsxWindowsFileServerVolumeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		a,
		"resetHostPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ResetS3FilesVolumeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3FilesVolumeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTaskDefinition_VolumePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

