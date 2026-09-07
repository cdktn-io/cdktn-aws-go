package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/batch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/batch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsJobDefinition_ContainerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Command() *[]*string
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
	Environment() DataAwsJobDefinition_EnvironmentPropertyList
	// Experimental.
	EphemeralStorage() DataAwsJobDefinition_EphemeralStoragePropertyList
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	FargatePlatformConfiguration() DataAwsJobDefinition_FargatePlatformConfigurationPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	Image() *string
	// Experimental.
	InstanceType() *string
	// Experimental.
	InternalValue() *DataAwsJobDefinition_ContainerProperty
	// Experimental.
	SetInternalValue(val *DataAwsJobDefinition_ContainerProperty)
	// Experimental.
	JobRoleArn() *string
	// Experimental.
	LinuxParameters() DataAwsJobDefinition_LinuxParametersPropertyList
	// Experimental.
	LogConfiguration() DataAwsJobDefinition_LogConfigurationPropertyList
	// Experimental.
	MountPoints() DataAwsJobDefinition_MountPointsPropertyList
	// Experimental.
	NetworkConfiguration() DataAwsJobDefinition_NetworkConfigurationPropertyList
	// Experimental.
	Privileged() cdktn.IResolvable
	// Experimental.
	ReadonlyRootFilesystem() cdktn.IResolvable
	// Experimental.
	ResourceRequirements() DataAwsJobDefinition_ResourceRequirementsPropertyList
	// Experimental.
	RuntimePlatform() DataAwsJobDefinition_RuntimePlatformPropertyList
	// Experimental.
	Secrets() DataAwsJobDefinition_SecretsPropertyList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Ulimits() DataAwsJobDefinition_UlimitsPropertyList
	// Experimental.
	User() *string
	// Experimental.
	Volumes() DataAwsJobDefinition_NodePropertiesNodeRangePropertiesContainerVolumesPropertyList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsJobDefinition_ContainerPropertyOutputReference
type jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Environment() DataAwsJobDefinition_EnvironmentPropertyList {
	var returns DataAwsJobDefinition_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) EphemeralStorage() DataAwsJobDefinition_EphemeralStoragePropertyList {
	var returns DataAwsJobDefinition_EphemeralStoragePropertyList
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) FargatePlatformConfiguration() DataAwsJobDefinition_FargatePlatformConfigurationPropertyList {
	var returns DataAwsJobDefinition_FargatePlatformConfigurationPropertyList
	_jsii_.Get(
		j,
		"fargatePlatformConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) InternalValue() *DataAwsJobDefinition_ContainerProperty {
	var returns *DataAwsJobDefinition_ContainerProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) JobRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) LinuxParameters() DataAwsJobDefinition_LinuxParametersPropertyList {
	var returns DataAwsJobDefinition_LinuxParametersPropertyList
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) LogConfiguration() DataAwsJobDefinition_LogConfigurationPropertyList {
	var returns DataAwsJobDefinition_LogConfigurationPropertyList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) MountPoints() DataAwsJobDefinition_MountPointsPropertyList {
	var returns DataAwsJobDefinition_MountPointsPropertyList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) NetworkConfiguration() DataAwsJobDefinition_NetworkConfigurationPropertyList {
	var returns DataAwsJobDefinition_NetworkConfigurationPropertyList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Privileged() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ReadonlyRootFilesystem() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ResourceRequirements() DataAwsJobDefinition_ResourceRequirementsPropertyList {
	var returns DataAwsJobDefinition_ResourceRequirementsPropertyList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) RuntimePlatform() DataAwsJobDefinition_RuntimePlatformPropertyList {
	var returns DataAwsJobDefinition_RuntimePlatformPropertyList
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Secrets() DataAwsJobDefinition_SecretsPropertyList {
	var returns DataAwsJobDefinition_SecretsPropertyList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Ulimits() DataAwsJobDefinition_UlimitsPropertyList {
	var returns DataAwsJobDefinition_UlimitsPropertyList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Volumes() DataAwsJobDefinition_NodePropertiesNodeRangePropertiesContainerVolumesPropertyList {
	var returns DataAwsJobDefinition_NodePropertiesNodeRangePropertiesContainerVolumesPropertyList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsJobDefinition_ContainerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsJobDefinition_ContainerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsJobDefinition_ContainerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.DataAwsJobDefinition.ContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsJobDefinition_ContainerPropertyOutputReference_Override(d DataAwsJobDefinition_ContainerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.DataAwsJobDefinition.ContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference)SetInternalValue(val *DataAwsJobDefinition_ContainerProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsJobDefinition_ContainerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

