package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTask_OptionsPropertyOutputReference interface {
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
	InternalValue() *TfTask_OptionsProperty
	// Experimental.
	SetInternalValue(val *TfTask_OptionsProperty)
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

// The jsii proxy struct for TfTask_OptionsPropertyOutputReference
type jsiiProxy_TfTask_OptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) Atime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) AtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) BytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) BytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) InternalValue() *TfTask_OptionsProperty {
	var returns *TfTask_OptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) Mtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) MtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) ObjectTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) ObjectTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) OverwriteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) OverwriteModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PosixPermissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PosixPermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PreserveDeletedFiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PreserveDeletedFilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PreserveDevices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) PreserveDevicesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) SecurityDescriptorCopyFlags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDescriptorCopyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) SecurityDescriptorCopyFlagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDescriptorCopyFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TaskQueueing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TaskQueueingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) TransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) VerifyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference) VerifyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTask_OptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTask_OptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTask_OptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTask_OptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTask_OptionsPropertyOutputReference_Override(t TfTask_OptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.TfTask.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetAtime(val *string) {
	if err := j.validateSetAtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atime",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetBytesPerSecond(val *float64) {
	if err := j.validateSetBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetInternalValue(val *TfTask_OptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetMtime(val *string) {
	if err := j.validateSetMtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtime",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetObjectTags(val *string) {
	if err := j.validateSetObjectTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectTags",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetOverwriteMode(val *string) {
	if err := j.validateSetOverwriteModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteMode",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetPosixPermissions(val *string) {
	if err := j.validateSetPosixPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posixPermissions",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetPreserveDeletedFiles(val *string) {
	if err := j.validateSetPreserveDeletedFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDeletedFiles",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetPreserveDevices(val *string) {
	if err := j.validateSetPreserveDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDevices",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetSecurityDescriptorCopyFlags(val *string) {
	if err := j.validateSetSecurityDescriptorCopyFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDescriptorCopyFlags",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetTaskQueueing(val *string) {
	if err := j.validateSetTaskQueueingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskQueueing",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetTransferMode(val *string) {
	if err := j.validateSetTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferMode",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (j *jsiiProxy_TfTask_OptionsPropertyOutputReference)SetVerifyMode(val *string) {
	if err := j.validateSetVerifyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyMode",
		val,
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetAtime() {
	_jsii_.InvokeVoid(
		t,
		"resetAtime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetBytesPerSecond() {
	_jsii_.InvokeVoid(
		t,
		"resetBytesPerSecond",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		t,
		"resetGid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetMtime() {
	_jsii_.InvokeVoid(
		t,
		"resetMtime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetObjectTags() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetOverwriteMode() {
	_jsii_.InvokeVoid(
		t,
		"resetOverwriteMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetPosixPermissions() {
	_jsii_.InvokeVoid(
		t,
		"resetPosixPermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetPreserveDeletedFiles() {
	_jsii_.InvokeVoid(
		t,
		"resetPreserveDeletedFiles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetPreserveDevices() {
	_jsii_.InvokeVoid(
		t,
		"resetPreserveDevices",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetSecurityDescriptorCopyFlags() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityDescriptorCopyFlags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetTaskQueueing() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskQueueing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetTransferMode() {
	_jsii_.InvokeVoid(
		t,
		"resetTransferMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		t,
		"resetUid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ResetVerifyMode() {
	_jsii_.InvokeVoid(
		t,
		"resetVerifyMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTask_OptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

