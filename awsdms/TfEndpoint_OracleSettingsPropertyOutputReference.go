package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_OracleSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessAlternateDirectly() interface{}
	// Experimental.
	SetAccessAlternateDirectly(val interface{})
	// Experimental.
	AccessAlternateDirectlyInput() interface{}
	// Experimental.
	AdditionalArchivedLogDestId() *float64
	// Experimental.
	SetAdditionalArchivedLogDestId(val *float64)
	// Experimental.
	AdditionalArchivedLogDestIdInput() *float64
	// Experimental.
	AddSupplementalLogging() interface{}
	// Experimental.
	SetAddSupplementalLogging(val interface{})
	// Experimental.
	AddSupplementalLoggingInput() interface{}
	// Experimental.
	AllowSelectedNestedTables() interface{}
	// Experimental.
	SetAllowSelectedNestedTables(val interface{})
	// Experimental.
	AllowSelectedNestedTablesInput() interface{}
	// Experimental.
	ArchivedLogDestId() *float64
	// Experimental.
	SetArchivedLogDestId(val *float64)
	// Experimental.
	ArchivedLogDestIdInput() *float64
	// Experimental.
	ArchivedLogsOnly() interface{}
	// Experimental.
	SetArchivedLogsOnly(val interface{})
	// Experimental.
	ArchivedLogsOnlyInput() interface{}
	// Experimental.
	AsmPassword() *string
	// Experimental.
	SetAsmPassword(val *string)
	// Experimental.
	AsmPasswordInput() *string
	// Experimental.
	AsmServer() *string
	// Experimental.
	SetAsmServer(val *string)
	// Experimental.
	AsmServerInput() *string
	// Experimental.
	AsmUser() *string
	// Experimental.
	SetAsmUser(val *string)
	// Experimental.
	AsmUserInput() *string
	// Experimental.
	AuthenticationMethod() *string
	// Experimental.
	SetAuthenticationMethod(val *string)
	// Experimental.
	AuthenticationMethodInput() *string
	// Experimental.
	CharLengthSemantics() *string
	// Experimental.
	SetCharLengthSemantics(val *string)
	// Experimental.
	CharLengthSemanticsInput() *string
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
	ConvertTimestampWithZoneToUtc() interface{}
	// Experimental.
	SetConvertTimestampWithZoneToUtc(val interface{})
	// Experimental.
	ConvertTimestampWithZoneToUtcInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DirectPathNoLog() interface{}
	// Experimental.
	SetDirectPathNoLog(val interface{})
	// Experimental.
	DirectPathNoLogInput() interface{}
	// Experimental.
	DirectPathParallelLoad() interface{}
	// Experimental.
	SetDirectPathParallelLoad(val interface{})
	// Experimental.
	DirectPathParallelLoadInput() interface{}
	// Experimental.
	EnableHomogenousTablespace() interface{}
	// Experimental.
	SetEnableHomogenousTablespace(val interface{})
	// Experimental.
	EnableHomogenousTablespaceInput() interface{}
	// Experimental.
	ExtraArchivedLogDestIds() *[]*float64
	// Experimental.
	SetExtraArchivedLogDestIds(val *[]*float64)
	// Experimental.
	ExtraArchivedLogDestIdsInput() *[]*float64
	// Experimental.
	FailTaskOnLobTruncation() interface{}
	// Experimental.
	SetFailTaskOnLobTruncation(val interface{})
	// Experimental.
	FailTaskOnLobTruncationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEndpoint_OracleSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_OracleSettingsProperty)
	// Experimental.
	NumberDatatypeScale() *float64
	// Experimental.
	SetNumberDatatypeScale(val *float64)
	// Experimental.
	NumberDatatypeScaleInput() *float64
	// Experimental.
	OpenTransactionWindow() *float64
	// Experimental.
	SetOpenTransactionWindow(val *float64)
	// Experimental.
	OpenTransactionWindowInput() *float64
	// Experimental.
	OraclePathPrefix() *string
	// Experimental.
	SetOraclePathPrefix(val *string)
	// Experimental.
	OraclePathPrefixInput() *string
	// Experimental.
	ParallelAsmReadThreads() *float64
	// Experimental.
	SetParallelAsmReadThreads(val *float64)
	// Experimental.
	ParallelAsmReadThreadsInput() *float64
	// Experimental.
	ReadAheadBlocks() *float64
	// Experimental.
	SetReadAheadBlocks(val *float64)
	// Experimental.
	ReadAheadBlocksInput() *float64
	// Experimental.
	ReadTableSpaceName() interface{}
	// Experimental.
	SetReadTableSpaceName(val interface{})
	// Experimental.
	ReadTableSpaceNameInput() interface{}
	// Experimental.
	ReplacePathPrefix() interface{}
	// Experimental.
	SetReplacePathPrefix(val interface{})
	// Experimental.
	ReplacePathPrefixInput() interface{}
	// Experimental.
	RetryInterval() *float64
	// Experimental.
	SetRetryInterval(val *float64)
	// Experimental.
	RetryIntervalInput() *float64
	// Experimental.
	SecretsManagerOracleAsmAccessRoleArn() *string
	// Experimental.
	SetSecretsManagerOracleAsmAccessRoleArn(val *string)
	// Experimental.
	SecretsManagerOracleAsmAccessRoleArnInput() *string
	// Experimental.
	SecretsManagerOracleAsmSecretId() *string
	// Experimental.
	SetSecretsManagerOracleAsmSecretId(val *string)
	// Experimental.
	SecretsManagerOracleAsmSecretIdInput() *string
	// Experimental.
	SecurityDbEncryption() *string
	// Experimental.
	SetSecurityDbEncryption(val *string)
	// Experimental.
	SecurityDbEncryptionInput() *string
	// Experimental.
	SecurityDbEncryptionName() *string
	// Experimental.
	SetSecurityDbEncryptionName(val *string)
	// Experimental.
	SecurityDbEncryptionNameInput() *string
	// Experimental.
	SpatialDataOptionToGeoJsonFunctionName() *string
	// Experimental.
	SetSpatialDataOptionToGeoJsonFunctionName(val *string)
	// Experimental.
	SpatialDataOptionToGeoJsonFunctionNameInput() *string
	// Experimental.
	StandbyDelayTime() *float64
	// Experimental.
	SetStandbyDelayTime(val *float64)
	// Experimental.
	StandbyDelayTimeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrimSpaceInChar() interface{}
	// Experimental.
	SetTrimSpaceInChar(val interface{})
	// Experimental.
	TrimSpaceInCharInput() interface{}
	// Experimental.
	UseAlternateFolderForOnline() interface{}
	// Experimental.
	SetUseAlternateFolderForOnline(val interface{})
	// Experimental.
	UseAlternateFolderForOnlineInput() interface{}
	// Experimental.
	UseBfile() interface{}
	// Experimental.
	SetUseBfile(val interface{})
	// Experimental.
	UseBfileInput() interface{}
	// Experimental.
	UseDirectPathFullLoad() interface{}
	// Experimental.
	SetUseDirectPathFullLoad(val interface{})
	// Experimental.
	UseDirectPathFullLoadInput() interface{}
	// Experimental.
	UseLogminerReader() interface{}
	// Experimental.
	SetUseLogminerReader(val interface{})
	// Experimental.
	UseLogminerReaderInput() interface{}
	// Experimental.
	UsePathPrefix() *string
	// Experimental.
	SetUsePathPrefix(val *string)
	// Experimental.
	UsePathPrefixInput() *string
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
	ResetAccessAlternateDirectly()
	// Experimental.
	ResetAdditionalArchivedLogDestId()
	// Experimental.
	ResetAddSupplementalLogging()
	// Experimental.
	ResetAllowSelectedNestedTables()
	// Experimental.
	ResetArchivedLogDestId()
	// Experimental.
	ResetArchivedLogsOnly()
	// Experimental.
	ResetAsmPassword()
	// Experimental.
	ResetAsmServer()
	// Experimental.
	ResetAsmUser()
	// Experimental.
	ResetAuthenticationMethod()
	// Experimental.
	ResetCharLengthSemantics()
	// Experimental.
	ResetConvertTimestampWithZoneToUtc()
	// Experimental.
	ResetDirectPathNoLog()
	// Experimental.
	ResetDirectPathParallelLoad()
	// Experimental.
	ResetEnableHomogenousTablespace()
	// Experimental.
	ResetExtraArchivedLogDestIds()
	// Experimental.
	ResetFailTaskOnLobTruncation()
	// Experimental.
	ResetNumberDatatypeScale()
	// Experimental.
	ResetOpenTransactionWindow()
	// Experimental.
	ResetOraclePathPrefix()
	// Experimental.
	ResetParallelAsmReadThreads()
	// Experimental.
	ResetReadAheadBlocks()
	// Experimental.
	ResetReadTableSpaceName()
	// Experimental.
	ResetReplacePathPrefix()
	// Experimental.
	ResetRetryInterval()
	// Experimental.
	ResetSecretsManagerOracleAsmAccessRoleArn()
	// Experimental.
	ResetSecretsManagerOracleAsmSecretId()
	// Experimental.
	ResetSecurityDbEncryption()
	// Experimental.
	ResetSecurityDbEncryptionName()
	// Experimental.
	ResetSpatialDataOptionToGeoJsonFunctionName()
	// Experimental.
	ResetStandbyDelayTime()
	// Experimental.
	ResetTrimSpaceInChar()
	// Experimental.
	ResetUseAlternateFolderForOnline()
	// Experimental.
	ResetUseBfile()
	// Experimental.
	ResetUseDirectPathFullLoad()
	// Experimental.
	ResetUseLogminerReader()
	// Experimental.
	ResetUsePathPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_OracleSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AccessAlternateDirectly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAlternateDirectly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AccessAlternateDirectlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessAlternateDirectlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AdditionalArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalArchivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AdditionalArchivedLogDestIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalArchivedLogDestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AddSupplementalLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addSupplementalLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AddSupplementalLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addSupplementalLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AllowSelectedNestedTables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelectedNestedTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AllowSelectedNestedTablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSelectedNestedTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ArchivedLogDestId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archivedLogDestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ArchivedLogDestIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"archivedLogDestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ArchivedLogsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedLogsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ArchivedLogsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedLogsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AsmUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asmUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) CharLengthSemantics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"charLengthSemantics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) CharLengthSemanticsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"charLengthSemanticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ConvertTimestampWithZoneToUtc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convertTimestampWithZoneToUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ConvertTimestampWithZoneToUtcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convertTimestampWithZoneToUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) DirectPathNoLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathNoLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) DirectPathNoLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathNoLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) DirectPathParallelLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathParallelLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) DirectPathParallelLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPathParallelLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) EnableHomogenousTablespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHomogenousTablespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) EnableHomogenousTablespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHomogenousTablespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ExtraArchivedLogDestIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"extraArchivedLogDestIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ExtraArchivedLogDestIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"extraArchivedLogDestIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) FailTaskOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTaskOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) FailTaskOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTaskOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) InternalValue() *TfEndpoint_OracleSettingsProperty {
	var returns *TfEndpoint_OracleSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) NumberDatatypeScale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberDatatypeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) NumberDatatypeScaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberDatatypeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) OpenTransactionWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"openTransactionWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) OpenTransactionWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"openTransactionWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) OraclePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oraclePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) OraclePathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oraclePathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ParallelAsmReadThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelAsmReadThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ParallelAsmReadThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelAsmReadThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReadAheadBlocks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readAheadBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReadAheadBlocksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readAheadBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReadTableSpaceName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readTableSpaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReadTableSpaceNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readTableSpaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReplacePathPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replacePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ReplacePathPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replacePathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) RetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) RetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecretsManagerOracleAsmAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecretsManagerOracleAsmAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecretsManagerOracleAsmSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecretsManagerOracleAsmSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerOracleAsmSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecurityDbEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecurityDbEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecurityDbEncryptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SecurityDbEncryptionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDbEncryptionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SpatialDataOptionToGeoJsonFunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialDataOptionToGeoJsonFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) SpatialDataOptionToGeoJsonFunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialDataOptionToGeoJsonFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) StandbyDelayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"standbyDelayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) StandbyDelayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"standbyDelayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) TrimSpaceInChar() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpaceInChar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) TrimSpaceInCharInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimSpaceInCharInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseAlternateFolderForOnline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAlternateFolderForOnline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseAlternateFolderForOnlineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAlternateFolderForOnlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseBfile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseBfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseDirectPathFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDirectPathFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseDirectPathFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDirectPathFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseLogminerReader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogminerReader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UseLogminerReaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLogminerReaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UsePathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) UsePathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePathPrefixInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_OracleSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_OracleSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_OracleSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.OracleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_OracleSettingsPropertyOutputReference_Override(t TfEndpoint_OracleSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.OracleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAccessAlternateDirectly(val interface{}) {
	if err := j.validateSetAccessAlternateDirectlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessAlternateDirectly",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAdditionalArchivedLogDestId(val *float64) {
	if err := j.validateSetAdditionalArchivedLogDestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalArchivedLogDestId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAddSupplementalLogging(val interface{}) {
	if err := j.validateSetAddSupplementalLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addSupplementalLogging",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAllowSelectedNestedTables(val interface{}) {
	if err := j.validateSetAllowSelectedNestedTablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSelectedNestedTables",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetArchivedLogDestId(val *float64) {
	if err := j.validateSetArchivedLogDestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivedLogDestId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetArchivedLogsOnly(val interface{}) {
	if err := j.validateSetArchivedLogsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivedLogsOnly",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAsmPassword(val *string) {
	if err := j.validateSetAsmPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmPassword",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAsmServer(val *string) {
	if err := j.validateSetAsmServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmServer",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAsmUser(val *string) {
	if err := j.validateSetAsmUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asmUser",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetCharLengthSemantics(val *string) {
	if err := j.validateSetCharLengthSemanticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"charLengthSemantics",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetConvertTimestampWithZoneToUtc(val interface{}) {
	if err := j.validateSetConvertTimestampWithZoneToUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"convertTimestampWithZoneToUtc",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetDirectPathNoLog(val interface{}) {
	if err := j.validateSetDirectPathNoLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPathNoLog",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetDirectPathParallelLoad(val interface{}) {
	if err := j.validateSetDirectPathParallelLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPathParallelLoad",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetEnableHomogenousTablespace(val interface{}) {
	if err := j.validateSetEnableHomogenousTablespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHomogenousTablespace",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetExtraArchivedLogDestIds(val *[]*float64) {
	if err := j.validateSetExtraArchivedLogDestIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraArchivedLogDestIds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetFailTaskOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTaskOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTaskOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_OracleSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetNumberDatatypeScale(val *float64) {
	if err := j.validateSetNumberDatatypeScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberDatatypeScale",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetOpenTransactionWindow(val *float64) {
	if err := j.validateSetOpenTransactionWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openTransactionWindow",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetOraclePathPrefix(val *string) {
	if err := j.validateSetOraclePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oraclePathPrefix",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetParallelAsmReadThreads(val *float64) {
	if err := j.validateSetParallelAsmReadThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelAsmReadThreads",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetReadAheadBlocks(val *float64) {
	if err := j.validateSetReadAheadBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAheadBlocks",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetReadTableSpaceName(val interface{}) {
	if err := j.validateSetReadTableSpaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readTableSpaceName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetReplacePathPrefix(val interface{}) {
	if err := j.validateSetReplacePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacePathPrefix",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetRetryInterval(val *float64) {
	if err := j.validateSetRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryInterval",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetSecretsManagerOracleAsmAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetSecretsManagerOracleAsmSecretId(val *string) {
	if err := j.validateSetSecretsManagerOracleAsmSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerOracleAsmSecretId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetSecurityDbEncryption(val *string) {
	if err := j.validateSetSecurityDbEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDbEncryption",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetSecurityDbEncryptionName(val *string) {
	if err := j.validateSetSecurityDbEncryptionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDbEncryptionName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetSpatialDataOptionToGeoJsonFunctionName(val *string) {
	if err := j.validateSetSpatialDataOptionToGeoJsonFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spatialDataOptionToGeoJsonFunctionName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetStandbyDelayTime(val *float64) {
	if err := j.validateSetStandbyDelayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyDelayTime",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetTrimSpaceInChar(val interface{}) {
	if err := j.validateSetTrimSpaceInCharParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpaceInChar",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetUseAlternateFolderForOnline(val interface{}) {
	if err := j.validateSetUseAlternateFolderForOnlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAlternateFolderForOnline",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetUseBfile(val interface{}) {
	if err := j.validateSetUseBfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBfile",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetUseDirectPathFullLoad(val interface{}) {
	if err := j.validateSetUseDirectPathFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDirectPathFullLoad",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetUseLogminerReader(val interface{}) {
	if err := j.validateSetUseLogminerReaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLogminerReader",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference)SetUsePathPrefix(val *string) {
	if err := j.validateSetUsePathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePathPrefix",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAccessAlternateDirectly() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessAlternateDirectly",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAdditionalArchivedLogDestId() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalArchivedLogDestId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAddSupplementalLogging() {
	_jsii_.InvokeVoid(
		t,
		"resetAddSupplementalLogging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAllowSelectedNestedTables() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowSelectedNestedTables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetArchivedLogDestId() {
	_jsii_.InvokeVoid(
		t,
		"resetArchivedLogDestId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetArchivedLogsOnly() {
	_jsii_.InvokeVoid(
		t,
		"resetArchivedLogsOnly",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAsmPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetAsmPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAsmServer() {
	_jsii_.InvokeVoid(
		t,
		"resetAsmServer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAsmUser() {
	_jsii_.InvokeVoid(
		t,
		"resetAsmUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetCharLengthSemantics() {
	_jsii_.InvokeVoid(
		t,
		"resetCharLengthSemantics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetConvertTimestampWithZoneToUtc() {
	_jsii_.InvokeVoid(
		t,
		"resetConvertTimestampWithZoneToUtc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetDirectPathNoLog() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectPathNoLog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetDirectPathParallelLoad() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectPathParallelLoad",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetEnableHomogenousTablespace() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableHomogenousTablespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetExtraArchivedLogDestIds() {
	_jsii_.InvokeVoid(
		t,
		"resetExtraArchivedLogDestIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetFailTaskOnLobTruncation() {
	_jsii_.InvokeVoid(
		t,
		"resetFailTaskOnLobTruncation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetNumberDatatypeScale() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberDatatypeScale",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetOpenTransactionWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenTransactionWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetOraclePathPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetOraclePathPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetParallelAsmReadThreads() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelAsmReadThreads",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetReadAheadBlocks() {
	_jsii_.InvokeVoid(
		t,
		"resetReadAheadBlocks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetReadTableSpaceName() {
	_jsii_.InvokeVoid(
		t,
		"resetReadTableSpaceName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetReplacePathPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetReplacePathPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetRetryInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetSecretsManagerOracleAsmAccessRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerOracleAsmAccessRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetSecretsManagerOracleAsmSecretId() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerOracleAsmSecretId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetSecurityDbEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityDbEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetSecurityDbEncryptionName() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityDbEncryptionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetSpatialDataOptionToGeoJsonFunctionName() {
	_jsii_.InvokeVoid(
		t,
		"resetSpatialDataOptionToGeoJsonFunctionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetStandbyDelayTime() {
	_jsii_.InvokeVoid(
		t,
		"resetStandbyDelayTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetTrimSpaceInChar() {
	_jsii_.InvokeVoid(
		t,
		"resetTrimSpaceInChar",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetUseAlternateFolderForOnline() {
	_jsii_.InvokeVoid(
		t,
		"resetUseAlternateFolderForOnline",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetUseBfile() {
	_jsii_.InvokeVoid(
		t,
		"resetUseBfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetUseDirectPathFullLoad() {
	_jsii_.InvokeVoid(
		t,
		"resetUseDirectPathFullLoad",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetUseLogminerReader() {
	_jsii_.InvokeVoid(
		t,
		"resetUseLogminerReader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ResetUsePathPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetUsePathPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_OracleSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

