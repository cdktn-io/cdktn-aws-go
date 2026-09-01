package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference interface {
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
	DependsOn() AwsEcsDaemonTaskDefinition_DependsOnPropertyList
	// Experimental.
	DependsOnInput() interface{}
	// Experimental.
	EntryPoint() *[]*string
	// Experimental.
	SetEntryPoint(val *[]*string)
	// Experimental.
	EntryPointInput() *[]*string
	// Experimental.
	Environment() AwsEcsDaemonTaskDefinition_EnvironmentPropertyList
	// Experimental.
	EnvironmentFile() AwsEcsDaemonTaskDefinition_EnvironmentFilePropertyList
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
	FirelensConfiguration() AwsEcsDaemonTaskDefinition_FirelensConfigurationPropertyList
	// Experimental.
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HealthCheck() AwsEcsDaemonTaskDefinition_HealthCheckPropertyList
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
	LinuxParameters() AwsEcsDaemonTaskDefinition_LinuxParametersPropertyList
	// Experimental.
	LinuxParametersInput() interface{}
	// Experimental.
	LogConfiguration() AwsEcsDaemonTaskDefinition_LogConfigurationPropertyList
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
	MountPoint() AwsEcsDaemonTaskDefinition_MountPointPropertyList
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
	RepositoryCredentials() AwsEcsDaemonTaskDefinition_RepositoryCredentialsPropertyList
	// Experimental.
	RepositoryCredentialsInput() interface{}
	// Experimental.
	RestartPolicy() AwsEcsDaemonTaskDefinition_RestartPolicyPropertyList
	// Experimental.
	RestartPolicyInput() interface{}
	// Experimental.
	Secret() AwsEcsDaemonTaskDefinition_SecretPropertyList
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
	SystemControl() AwsEcsDaemonTaskDefinition_SystemControlPropertyList
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
	Ulimit() AwsEcsDaemonTaskDefinition_UlimitPropertyList
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

// The jsii proxy struct for AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference
type jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) DependsOn() AwsEcsDaemonTaskDefinition_DependsOnPropertyList {
	var returns AwsEcsDaemonTaskDefinition_DependsOnPropertyList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Environment() AwsEcsDaemonTaskDefinition_EnvironmentPropertyList {
	var returns AwsEcsDaemonTaskDefinition_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentFile() AwsEcsDaemonTaskDefinition_EnvironmentFilePropertyList {
	var returns AwsEcsDaemonTaskDefinition_EnvironmentFilePropertyList
	_jsii_.Get(
		j,
		"environmentFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) FirelensConfiguration() AwsEcsDaemonTaskDefinition_FirelensConfigurationPropertyList {
	var returns AwsEcsDaemonTaskDefinition_FirelensConfigurationPropertyList
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) HealthCheck() AwsEcsDaemonTaskDefinition_HealthCheckPropertyList {
	var returns AwsEcsDaemonTaskDefinition_HealthCheckPropertyList
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Interactive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InteractiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LinuxParameters() AwsEcsDaemonTaskDefinition_LinuxParametersPropertyList {
	var returns AwsEcsDaemonTaskDefinition_LinuxParametersPropertyList
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LogConfiguration() AwsEcsDaemonTaskDefinition_LogConfigurationPropertyList {
	var returns AwsEcsDaemonTaskDefinition_LogConfigurationPropertyList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MountPoint() AwsEcsDaemonTaskDefinition_MountPointPropertyList {
	var returns AwsEcsDaemonTaskDefinition_MountPointPropertyList
	_jsii_.Get(
		j,
		"mountPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) MountPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PseudoTerminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PseudoTerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RepositoryCredentials() AwsEcsDaemonTaskDefinition_RepositoryCredentialsPropertyList {
	var returns AwsEcsDaemonTaskDefinition_RepositoryCredentialsPropertyList
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RestartPolicy() AwsEcsDaemonTaskDefinition_RestartPolicyPropertyList {
	var returns AwsEcsDaemonTaskDefinition_RestartPolicyPropertyList
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) RestartPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Secret() AwsEcsDaemonTaskDefinition_SecretPropertyList {
	var returns AwsEcsDaemonTaskDefinition_SecretPropertyList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) StopTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SystemControl() AwsEcsDaemonTaskDefinition_SystemControlPropertyList {
	var returns AwsEcsDaemonTaskDefinition_SystemControlPropertyList
	_jsii_.Get(
		j,
		"systemControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) SystemControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Ulimit() AwsEcsDaemonTaskDefinition_UlimitPropertyList {
	var returns AwsEcsDaemonTaskDefinition_UlimitPropertyList
	_jsii_.Get(
		j,
		"ulimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) UlimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsDaemonTaskDefinition.ContainerDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference_Override(a AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsDaemonTaskDefinition.ContainerDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetEntryPoint(val *[]*string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetInteractive(val interface{}) {
	if err := j.validateSetInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactive",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetPseudoTerminal(val interface{}) {
	if err := j.validateSetPseudoTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoTerminal",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetStartTimeout(val *float64) {
	if err := j.validateSetStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetStopTimeout(val *float64) {
	if err := j.validateSetStopTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutDependsOn(value interface{}) {
	if err := a.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutEnvironment(value interface{}) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutEnvironmentFile(value interface{}) {
	if err := a.validatePutEnvironmentFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentFile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutFirelensConfiguration(value interface{}) {
	if err := a.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutHealthCheck(value interface{}) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutLinuxParameters(value interface{}) {
	if err := a.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutLogConfiguration(value interface{}) {
	if err := a.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutMountPoint(value interface{}) {
	if err := a.validatePutMountPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMountPoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutRepositoryCredentials(value interface{}) {
	if err := a.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutRestartPolicy(value interface{}) {
	if err := a.validatePutRestartPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestartPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutSecret(value interface{}) {
	if err := a.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutSystemControl(value interface{}) {
	if err := a.validatePutSystemControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSystemControl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) PutUlimit(value interface{}) {
	if err := a.validatePutUlimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUlimit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		a,
		"resetCpu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		a,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEnvironmentFile() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		a,
		"resetEssential",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetInteractive() {
	_jsii_.InvokeVoid(
		a,
		"resetInteractive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		a,
		"resetMemory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		a,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetMountPoint() {
	_jsii_.InvokeVoid(
		a,
		"resetMountPoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetPseudoTerminal() {
	_jsii_.InvokeVoid(
		a,
		"resetPseudoTerminal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		a,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetStartTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetStopTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetStopTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetSystemControl() {
	_jsii_.InvokeVoid(
		a,
		"resetSystemControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetUlimit() {
	_jsii_.InvokeVoid(
		a,
		"resetUlimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		a,
		"resetUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsDaemonTaskDefinition_ContainerDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

