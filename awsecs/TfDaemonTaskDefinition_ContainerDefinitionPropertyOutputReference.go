package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Command() *[]*string
	// Experimental.
	SetCommand(val *[]*string)
	// Experimental.
	CommandInput() *[]*string
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
	Cpu() *float64
	// Experimental.
	SetCpu(val *float64)
	// Experimental.
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DependsOn() TfDaemonTaskDefinition_DependsOnPropertyList
	// Experimental.
	DependsOnInput() interface{}
	// Experimental.
	EntryPoint() *[]*string
	// Experimental.
	SetEntryPoint(val *[]*string)
	// Experimental.
	EntryPointInput() *[]*string
	// Experimental.
	Environment() TfDaemonTaskDefinition_EnvironmentPropertyList
	// Experimental.
	EnvironmentFile() TfDaemonTaskDefinition_EnvironmentFilePropertyList
	// Experimental.
	EnvironmentFileInput() interface{}
	// Experimental.
	EnvironmentInput() interface{}
	// Experimental.
	Essential() interface{}
	// Experimental.
	SetEssential(val interface{})
	// Experimental.
	EssentialInput() interface{}
	// Experimental.
	FirelensConfiguration() TfDaemonTaskDefinition_FirelensConfigurationPropertyList
	// Experimental.
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HealthCheck() TfDaemonTaskDefinition_HealthCheckPropertyList
	// Experimental.
	HealthCheckInput() interface{}
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageInput() *string
	// Experimental.
	Interactive() interface{}
	// Experimental.
	SetInteractive(val interface{})
	// Experimental.
	InteractiveInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LinuxParameters() TfDaemonTaskDefinition_LinuxParametersPropertyList
	// Experimental.
	LinuxParametersInput() interface{}
	// Experimental.
	LogConfiguration() TfDaemonTaskDefinition_LogConfigurationPropertyList
	// Experimental.
	LogConfigurationInput() interface{}
	// Experimental.
	Memory() *float64
	// Experimental.
	SetMemory(val *float64)
	// Experimental.
	MemoryInput() *float64
	// Experimental.
	MemoryReservation() *float64
	// Experimental.
	SetMemoryReservation(val *float64)
	// Experimental.
	MemoryReservationInput() *float64
	// Experimental.
	MountPoint() TfDaemonTaskDefinition_MountPointPropertyList
	// Experimental.
	MountPointInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Privileged() interface{}
	// Experimental.
	SetPrivileged(val interface{})
	// Experimental.
	PrivilegedInput() interface{}
	// Experimental.
	PseudoTerminal() interface{}
	// Experimental.
	SetPseudoTerminal(val interface{})
	// Experimental.
	PseudoTerminalInput() interface{}
	// Experimental.
	ReadonlyRootFilesystem() interface{}
	// Experimental.
	SetReadonlyRootFilesystem(val interface{})
	// Experimental.
	ReadonlyRootFilesystemInput() interface{}
	// Experimental.
	RepositoryCredentials() TfDaemonTaskDefinition_RepositoryCredentialsPropertyList
	// Experimental.
	RepositoryCredentialsInput() interface{}
	// Experimental.
	RestartPolicy() TfDaemonTaskDefinition_RestartPolicyPropertyList
	// Experimental.
	RestartPolicyInput() interface{}
	// Experimental.
	Secret() TfDaemonTaskDefinition_SecretPropertyList
	// Experimental.
	SecretInput() interface{}
	// Experimental.
	StartTimeout() *float64
	// Experimental.
	SetStartTimeout(val *float64)
	// Experimental.
	StartTimeoutInput() *float64
	// Experimental.
	StopTimeout() *float64
	// Experimental.
	SetStopTimeout(val *float64)
	// Experimental.
	StopTimeoutInput() *float64
	// Experimental.
	SystemControl() TfDaemonTaskDefinition_SystemControlPropertyList
	// Experimental.
	SystemControlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Ulimit() TfDaemonTaskDefinition_UlimitPropertyList
	// Experimental.
	UlimitInput() interface{}
	// Experimental.
	User() *string
	// Experimental.
	SetUser(val *string)
	// Experimental.
	UserInput() *string
	// Experimental.
	WorkingDirectory() *string
	// Experimental.
	SetWorkingDirectory(val *string)
	// Experimental.
	WorkingDirectoryInput() *string
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
	PutDependsOn(value interface{})
	// Experimental.
	PutEnvironment(value interface{})
	// Experimental.
	PutEnvironmentFile(value interface{})
	// Experimental.
	PutFirelensConfiguration(value interface{})
	// Experimental.
	PutHealthCheck(value interface{})
	// Experimental.
	PutLinuxParameters(value interface{})
	// Experimental.
	PutLogConfiguration(value interface{})
	// Experimental.
	PutMountPoint(value interface{})
	// Experimental.
	PutRepositoryCredentials(value interface{})
	// Experimental.
	PutRestartPolicy(value interface{})
	// Experimental.
	PutSecret(value interface{})
	// Experimental.
	PutSystemControl(value interface{})
	// Experimental.
	PutUlimit(value interface{})
	// Experimental.
	ResetCommand()
	// Experimental.
	ResetCpu()
	// Experimental.
	ResetDependsOn()
	// Experimental.
	ResetEntryPoint()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetEnvironmentFile()
	// Experimental.
	ResetEssential()
	// Experimental.
	ResetFirelensConfiguration()
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetInteractive()
	// Experimental.
	ResetLinuxParameters()
	// Experimental.
	ResetLogConfiguration()
	// Experimental.
	ResetMemory()
	// Experimental.
	ResetMemoryReservation()
	// Experimental.
	ResetMountPoint()
	// Experimental.
	ResetName()
	// Experimental.
	ResetPrivileged()
	// Experimental.
	ResetPseudoTerminal()
	// Experimental.
	ResetReadonlyRootFilesystem()
	// Experimental.
	ResetRepositoryCredentials()
	// Experimental.
	ResetRestartPolicy()
	// Experimental.
	ResetSecret()
	// Experimental.
	ResetStartTimeout()
	// Experimental.
	ResetStopTimeout()
	// Experimental.
	ResetSystemControl()
	// Experimental.
	ResetUlimit()
	// Experimental.
	ResetUser()
	// Experimental.
	ResetWorkingDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference
type jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) DependsOn() TfDaemonTaskDefinition_DependsOnPropertyList {
	var returns TfDaemonTaskDefinition_DependsOnPropertyList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Environment() TfDaemonTaskDefinition_EnvironmentPropertyList {
	var returns TfDaemonTaskDefinition_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentFile() TfDaemonTaskDefinition_EnvironmentFilePropertyList {
	var returns TfDaemonTaskDefinition_EnvironmentFilePropertyList
	_jsii_.Get(
		j,
		"environmentFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) FirelensConfiguration() TfDaemonTaskDefinition_FirelensConfigurationPropertyList {
	var returns TfDaemonTaskDefinition_FirelensConfigurationPropertyList
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) HealthCheck() TfDaemonTaskDefinition_HealthCheckPropertyList {
	var returns TfDaemonTaskDefinition_HealthCheckPropertyList
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Interactive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InteractiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LinuxParameters() TfDaemonTaskDefinition_LinuxParametersPropertyList {
	var returns TfDaemonTaskDefinition_LinuxParametersPropertyList
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LogConfiguration() TfDaemonTaskDefinition_LogConfigurationPropertyList {
	var returns TfDaemonTaskDefinition_LogConfigurationPropertyList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MountPoint() TfDaemonTaskDefinition_MountPointPropertyList {
	var returns TfDaemonTaskDefinition_MountPointPropertyList
	_jsii_.Get(
		j,
		"mountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MountPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PseudoTerminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PseudoTerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RepositoryCredentials() TfDaemonTaskDefinition_RepositoryCredentialsPropertyList {
	var returns TfDaemonTaskDefinition_RepositoryCredentialsPropertyList
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RestartPolicy() TfDaemonTaskDefinition_RestartPolicyPropertyList {
	var returns TfDaemonTaskDefinition_RestartPolicyPropertyList
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RestartPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Secret() TfDaemonTaskDefinition_SecretPropertyList {
	var returns TfDaemonTaskDefinition_SecretPropertyList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StopTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SystemControl() TfDaemonTaskDefinition_SystemControlPropertyList {
	var returns TfDaemonTaskDefinition_SystemControlPropertyList
	_jsii_.Get(
		j,
		"systemControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SystemControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Ulimit() TfDaemonTaskDefinition_UlimitPropertyList {
	var returns TfDaemonTaskDefinition_UlimitPropertyList
	_jsii_.Get(
		j,
		"ulimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) UlimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfDaemonTaskDefinition.ContainerDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference_Override(t TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfDaemonTaskDefinition.ContainerDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetEntryPoint(val *[]*string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetInteractive(val interface{}) {
	if err := j.validateSetInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactive",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetPseudoTerminal(val interface{}) {
	if err := j.validateSetPseudoTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoTerminal",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetStartTimeout(val *float64) {
	if err := j.validateSetStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetStopTimeout(val *float64) {
	if err := j.validateSetStopTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutDependsOn(value interface{}) {
	if err := t.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutEnvironment(value interface{}) {
	if err := t.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutEnvironmentFile(value interface{}) {
	if err := t.validatePutEnvironmentFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironmentFile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutFirelensConfiguration(value interface{}) {
	if err := t.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutHealthCheck(value interface{}) {
	if err := t.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutLinuxParameters(value interface{}) {
	if err := t.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutLogConfiguration(value interface{}) {
	if err := t.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutMountPoint(value interface{}) {
	if err := t.validatePutMountPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMountPoint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutRepositoryCredentials(value interface{}) {
	if err := t.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutRestartPolicy(value interface{}) {
	if err := t.validatePutRestartPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRestartPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutSecret(value interface{}) {
	if err := t.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecret",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutSystemControl(value interface{}) {
	if err := t.validatePutSystemControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSystemControl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutUlimit(value interface{}) {
	if err := t.validatePutUlimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUlimit",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		t,
		"resetCpu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		t,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		t,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEnvironmentFile() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		t,
		"resetEssential",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetInteractive() {
	_jsii_.InvokeVoid(
		t,
		"resetInteractive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		t,
		"resetMemory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		t,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMountPoint() {
	_jsii_.InvokeVoid(
		t,
		"resetMountPoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetPseudoTerminal() {
	_jsii_.InvokeVoid(
		t,
		"resetPseudoTerminal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		t,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetStartTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetStartTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetStopTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetStopTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetSystemControl() {
	_jsii_.InvokeVoid(
		t,
		"resetSystemControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetUlimit() {
	_jsii_.InvokeVoid(
		t,
		"resetUlimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		t,
		"resetUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

