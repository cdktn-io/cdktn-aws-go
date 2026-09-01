package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDatasyncTask_OptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Atime() *string
	// Experimental.
	SetAtime(val *string)
	// Experimental.
	AtimeInput() *string
	// Experimental.
	BytesPerSecond() *float64
	// Experimental.
	SetBytesPerSecond(val *float64)
	// Experimental.
	BytesPerSecondInput() *float64
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
	Fqn() *string
	// Experimental.
	Gid() *string
	// Experimental.
	SetGid(val *string)
	// Experimental.
	GidInput() *string
	// Experimental.
	InternalValue() *AwsDatasyncTask_OptionsProperty
	// Experimental.
	SetInternalValue(val *AwsDatasyncTask_OptionsProperty)
	// Experimental.
	LogLevel() *string
	// Experimental.
	SetLogLevel(val *string)
	// Experimental.
	LogLevelInput() *string
	// Experimental.
	Mtime() *string
	// Experimental.
	SetMtime(val *string)
	// Experimental.
	MtimeInput() *string
	// Experimental.
	ObjectTags() *string
	// Experimental.
	SetObjectTags(val *string)
	// Experimental.
	ObjectTagsInput() *string
	// Experimental.
	OverwriteMode() *string
	// Experimental.
	SetOverwriteMode(val *string)
	// Experimental.
	OverwriteModeInput() *string
	// Experimental.
	PosixPermissions() *string
	// Experimental.
	SetPosixPermissions(val *string)
	// Experimental.
	PosixPermissionsInput() *string
	// Experimental.
	PreserveDeletedFiles() *string
	// Experimental.
	SetPreserveDeletedFiles(val *string)
	// Experimental.
	PreserveDeletedFilesInput() *string
	// Experimental.
	PreserveDevices() *string
	// Experimental.
	SetPreserveDevices(val *string)
	// Experimental.
	PreserveDevicesInput() *string
	// Experimental.
	SecurityDescriptorCopyFlags() *string
	// Experimental.
	SetSecurityDescriptorCopyFlags(val *string)
	// Experimental.
	SecurityDescriptorCopyFlagsInput() *string
	// Experimental.
	TaskQueueing() *string
	// Experimental.
	SetTaskQueueing(val *string)
	// Experimental.
	TaskQueueingInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransferMode() *string
	// Experimental.
	SetTransferMode(val *string)
	// Experimental.
	TransferModeInput() *string
	// Experimental.
	Uid() *string
	// Experimental.
	SetUid(val *string)
	// Experimental.
	UidInput() *string
	// Experimental.
	VerifyMode() *string
	// Experimental.
	SetVerifyMode(val *string)
	// Experimental.
	VerifyModeInput() *string
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
	ResetAtime()
	// Experimental.
	ResetBytesPerSecond()
	// Experimental.
	ResetGid()
	// Experimental.
	ResetLogLevel()
	// Experimental.
	ResetMtime()
	// Experimental.
	ResetObjectTags()
	// Experimental.
	ResetOverwriteMode()
	// Experimental.
	ResetPosixPermissions()
	// Experimental.
	ResetPreserveDeletedFiles()
	// Experimental.
	ResetPreserveDevices()
	// Experimental.
	ResetSecurityDescriptorCopyFlags()
	// Experimental.
	ResetTaskQueueing()
	// Experimental.
	ResetTransferMode()
	// Experimental.
	ResetUid()
	// Experimental.
	ResetVerifyMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDatasyncTask_OptionsPropertyOutputReference
type jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Atime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) AtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) BytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) BytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) InternalValue() *AwsDatasyncTask_OptionsProperty {
	var returns *AwsDatasyncTask_OptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Mtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) MtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ObjectTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ObjectTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) OverwriteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) OverwriteModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PosixPermissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PosixPermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PreserveDeletedFiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PreserveDeletedFilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PreserveDevices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) PreserveDevicesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) SecurityDescriptorCopyFlags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDescriptorCopyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) SecurityDescriptorCopyFlagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDescriptorCopyFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TaskQueueing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TaskQueueingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) TransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) VerifyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) VerifyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDatasyncTask_OptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDatasyncTask_OptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDatasyncTask_OptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsDatasyncTask.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDatasyncTask_OptionsPropertyOutputReference_Override(a AwsDatasyncTask_OptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsDatasyncTask.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetAtime(val *string) {
	if err := j.validateSetAtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atime",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetBytesPerSecond(val *float64) {
	if err := j.validateSetBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetInternalValue(val *AwsDatasyncTask_OptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetMtime(val *string) {
	if err := j.validateSetMtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtime",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetObjectTags(val *string) {
	if err := j.validateSetObjectTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectTags",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetOverwriteMode(val *string) {
	if err := j.validateSetOverwriteModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteMode",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetPosixPermissions(val *string) {
	if err := j.validateSetPosixPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posixPermissions",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetPreserveDeletedFiles(val *string) {
	if err := j.validateSetPreserveDeletedFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDeletedFiles",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetPreserveDevices(val *string) {
	if err := j.validateSetPreserveDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDevices",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetSecurityDescriptorCopyFlags(val *string) {
	if err := j.validateSetSecurityDescriptorCopyFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDescriptorCopyFlags",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetTaskQueueing(val *string) {
	if err := j.validateSetTaskQueueingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskQueueing",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetTransferMode(val *string) {
	if err := j.validateSetTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferMode",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (j *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference)SetVerifyMode(val *string) {
	if err := j.validateSetVerifyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyMode",
		val,
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetAtime() {
	_jsii_.InvokeVoid(
		a,
		"resetAtime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetBytesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		a,
		"resetGid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetMtime() {
	_jsii_.InvokeVoid(
		a,
		"resetMtime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetObjectTags() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetOverwriteMode() {
	_jsii_.InvokeVoid(
		a,
		"resetOverwriteMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetPosixPermissions() {
	_jsii_.InvokeVoid(
		a,
		"resetPosixPermissions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetPreserveDeletedFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveDeletedFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetPreserveDevices() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveDevices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetSecurityDescriptorCopyFlags() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityDescriptorCopyFlags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetTaskQueueing() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskQueueing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetTransferMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTransferMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		a,
		"resetUid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ResetVerifyMode() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifyMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDatasyncTask_OptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

