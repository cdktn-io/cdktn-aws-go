package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTaskDefinition_VolumePropertyOutputReference interface {
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
	DockerVolumeConfiguration() TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference
	// Experimental.
	DockerVolumeConfigurationInput() *TfTaskDefinition_DockerVolumeConfigurationProperty
	// Experimental.
	EfsVolumeConfiguration() TfTaskDefinition_EfsVolumeConfigurationPropertyOutputReference
	// Experimental.
	EfsVolumeConfigurationInput() *TfTaskDefinition_EfsVolumeConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FsxWindowsFileServerVolumeConfiguration() TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference
	// Experimental.
	FsxWindowsFileServerVolumeConfigurationInput() *TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty
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
	S3FilesVolumeConfiguration() TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference
	// Experimental.
	S3FilesVolumeConfigurationInput() *TfTaskDefinition_S3filesVolumeConfigurationProperty
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
	PutDockerVolumeConfiguration(value *TfTaskDefinition_DockerVolumeConfigurationProperty)
	// Experimental.
	PutEfsVolumeConfiguration(value *TfTaskDefinition_EfsVolumeConfigurationProperty)
	// Experimental.
	PutFsxWindowsFileServerVolumeConfiguration(value *TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty)
	// Experimental.
	PutS3FilesVolumeConfiguration(value *TfTaskDefinition_S3filesVolumeConfigurationProperty)
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

// The jsii proxy struct for TfTaskDefinition_VolumePropertyOutputReference
type jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ConfigureAtLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configureAtLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ConfigureAtLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configureAtLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) DockerVolumeConfiguration() TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference {
	var returns TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) DockerVolumeConfigurationInput() *TfTaskDefinition_DockerVolumeConfigurationProperty {
	var returns *TfTaskDefinition_DockerVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"dockerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) EfsVolumeConfiguration() TfTaskDefinition_EfsVolumeConfigurationPropertyOutputReference {
	var returns TfTaskDefinition_EfsVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"efsVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) EfsVolumeConfigurationInput() *TfTaskDefinition_EfsVolumeConfigurationProperty {
	var returns *TfTaskDefinition_EfsVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"efsVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) FsxWindowsFileServerVolumeConfiguration() TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference {
	var returns TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) FsxWindowsFileServerVolumeConfigurationInput() *TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty {
	var returns *TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) HostPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) HostPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) S3FilesVolumeConfiguration() TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference {
	var returns TfTaskDefinition_S3filesVolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3FilesVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) S3FilesVolumeConfigurationInput() *TfTaskDefinition_S3filesVolumeConfigurationProperty {
	var returns *TfTaskDefinition_S3filesVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"s3FilesVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTaskDefinition_VolumePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTaskDefinition_VolumePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTaskDefinition_VolumePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.VolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTaskDefinition_VolumePropertyOutputReference_Override(t TfTaskDefinition_VolumePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.VolumePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetConfigureAtLaunch(val interface{}) {
	if err := j.validateSetConfigureAtLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configureAtLaunch",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetHostPath(val *string) {
	if err := j.validateSetHostPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPath",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) PutDockerVolumeConfiguration(value *TfTaskDefinition_DockerVolumeConfigurationProperty) {
	if err := t.validatePutDockerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDockerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) PutEfsVolumeConfiguration(value *TfTaskDefinition_EfsVolumeConfigurationProperty) {
	if err := t.validatePutEfsVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEfsVolumeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) PutFsxWindowsFileServerVolumeConfiguration(value *TfTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty) {
	if err := t.validatePutFsxWindowsFileServerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFsxWindowsFileServerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) PutS3FilesVolumeConfiguration(value *TfTaskDefinition_S3filesVolumeConfigurationProperty) {
	if err := t.validatePutS3FilesVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3FilesVolumeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetConfigureAtLaunch() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigureAtLaunch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetDockerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDockerVolumeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetEfsVolumeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEfsVolumeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetFsxWindowsFileServerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetFsxWindowsFileServerVolumeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		t,
		"resetHostPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ResetS3FilesVolumeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3FilesVolumeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_VolumePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

