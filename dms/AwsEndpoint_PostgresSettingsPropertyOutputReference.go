package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_PostgresSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AfterConnectScript() *string
	// Experimental.
	SetAfterConnectScript(val *string)
	// Experimental.
	AfterConnectScriptInput() *string
	// Experimental.
	AuthenticationMethod() *string
	// Experimental.
	SetAuthenticationMethod(val *string)
	// Experimental.
	AuthenticationMethodInput() *string
	// Experimental.
	BabelfishDatabaseName() *string
	// Experimental.
	SetBabelfishDatabaseName(val *string)
	// Experimental.
	BabelfishDatabaseNameInput() *string
	// Experimental.
	CaptureDdls() interface{}
	// Experimental.
	SetCaptureDdls(val interface{})
	// Experimental.
	CaptureDdlsInput() interface{}
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
	DatabaseMode() *string
	// Experimental.
	SetDatabaseMode(val *string)
	// Experimental.
	DatabaseModeInput() *string
	// Experimental.
	DdlArtifactsSchema() *string
	// Experimental.
	SetDdlArtifactsSchema(val *string)
	// Experimental.
	DdlArtifactsSchemaInput() *string
	// Experimental.
	ExecuteTimeout() *float64
	// Experimental.
	SetExecuteTimeout(val *float64)
	// Experimental.
	ExecuteTimeoutInput() *float64
	// Experimental.
	FailTasksOnLobTruncation() interface{}
	// Experimental.
	SetFailTasksOnLobTruncation(val interface{})
	// Experimental.
	FailTasksOnLobTruncationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HeartbeatEnable() interface{}
	// Experimental.
	SetHeartbeatEnable(val interface{})
	// Experimental.
	HeartbeatEnableInput() interface{}
	// Experimental.
	HeartbeatFrequency() *float64
	// Experimental.
	SetHeartbeatFrequency(val *float64)
	// Experimental.
	HeartbeatFrequencyInput() *float64
	// Experimental.
	HeartbeatSchema() *string
	// Experimental.
	SetHeartbeatSchema(val *string)
	// Experimental.
	HeartbeatSchemaInput() *string
	// Experimental.
	InternalValue() *AwsEndpoint_PostgresSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_PostgresSettingsProperty)
	// Experimental.
	MapBooleanAsBoolean() interface{}
	// Experimental.
	SetMapBooleanAsBoolean(val interface{})
	// Experimental.
	MapBooleanAsBooleanInput() interface{}
	// Experimental.
	MapJsonbAsClob() interface{}
	// Experimental.
	SetMapJsonbAsClob(val interface{})
	// Experimental.
	MapJsonbAsClobInput() interface{}
	// Experimental.
	MapLongVarcharAs() *string
	// Experimental.
	SetMapLongVarcharAs(val *string)
	// Experimental.
	MapLongVarcharAsInput() *string
	// Experimental.
	MaxFileSize() *float64
	// Experimental.
	SetMaxFileSize(val *float64)
	// Experimental.
	MaxFileSizeInput() *float64
	// Experimental.
	PluginName() *string
	// Experimental.
	SetPluginName(val *string)
	// Experimental.
	PluginNameInput() *string
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	SetServiceAccessRoleArn(val *string)
	// Experimental.
	ServiceAccessRoleArnInput() *string
	// Experimental.
	SlotName() *string
	// Experimental.
	SetSlotName(val *string)
	// Experimental.
	SlotNameInput() *string
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
	ResetAfterConnectScript()
	// Experimental.
	ResetAuthenticationMethod()
	// Experimental.
	ResetBabelfishDatabaseName()
	// Experimental.
	ResetCaptureDdls()
	// Experimental.
	ResetDatabaseMode()
	// Experimental.
	ResetDdlArtifactsSchema()
	// Experimental.
	ResetExecuteTimeout()
	// Experimental.
	ResetFailTasksOnLobTruncation()
	// Experimental.
	ResetHeartbeatEnable()
	// Experimental.
	ResetHeartbeatFrequency()
	// Experimental.
	ResetHeartbeatSchema()
	// Experimental.
	ResetMapBooleanAsBoolean()
	// Experimental.
	ResetMapJsonbAsClob()
	// Experimental.
	ResetMapLongVarcharAs()
	// Experimental.
	ResetMaxFileSize()
	// Experimental.
	ResetPluginName()
	// Experimental.
	ResetServiceAccessRoleArn()
	// Experimental.
	ResetSlotName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEndpoint_PostgresSettingsPropertyOutputReference
type jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) BabelfishDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) BabelfishDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) CaptureDdls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) CaptureDdlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) DatabaseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) DatabaseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) DdlArtifactsSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) DdlArtifactsSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) FailTasksOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) FailTasksOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) InternalValue() *AwsEndpoint_PostgresSettingsProperty {
	var returns *AwsEndpoint_PostgresSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapBooleanAsBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapBooleanAsBooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBooleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapJsonbAsClob() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapJsonbAsClobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapLongVarcharAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MapLongVarcharAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) SlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) SlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_PostgresSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_PostgresSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_PostgresSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.PostgresSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_PostgresSettingsPropertyOutputReference_Override(a AwsEndpoint_PostgresSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.PostgresSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetBabelfishDatabaseName(val *string) {
	if err := j.validateSetBabelfishDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"babelfishDatabaseName",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetCaptureDdls(val interface{}) {
	if err := j.validateSetCaptureDdlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captureDdls",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetDatabaseMode(val *string) {
	if err := j.validateSetDatabaseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseMode",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetDdlArtifactsSchema(val *string) {
	if err := j.validateSetDdlArtifactsSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddlArtifactsSchema",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetFailTasksOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTasksOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTasksOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatEnable(val interface{}) {
	if err := j.validateSetHeartbeatEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatEnable",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatFrequency(val *float64) {
	if err := j.validateSetHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatSchema(val *string) {
	if err := j.validateSetHeartbeatSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatSchema",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_PostgresSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetMapBooleanAsBoolean(val interface{}) {
	if err := j.validateSetMapBooleanAsBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBooleanAsBoolean",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetMapJsonbAsClob(val interface{}) {
	if err := j.validateSetMapJsonbAsClobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapJsonbAsClob",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetMapLongVarcharAs(val *string) {
	if err := j.validateSetMapLongVarcharAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapLongVarcharAs",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetSlotName(val *string) {
	if err := j.validateSetSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotName",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		a,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetBabelfishDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetBabelfishDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetCaptureDdls() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptureDdls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetDatabaseMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetDdlArtifactsSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetDdlArtifactsSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetFailTasksOnLobTruncation() {
	_jsii_.InvokeVoid(
		a,
		"resetFailTasksOnLobTruncation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatEnable() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatEnable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetMapBooleanAsBoolean() {
	_jsii_.InvokeVoid(
		a,
		"resetMapBooleanAsBoolean",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetMapJsonbAsClob() {
	_jsii_.InvokeVoid(
		a,
		"resetMapJsonbAsClob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetMapLongVarcharAs() {
	_jsii_.InvokeVoid(
		a,
		"resetMapLongVarcharAs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		a,
		"resetPluginName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ResetSlotName() {
	_jsii_.InvokeVoid(
		a,
		"resetSlotName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_PostgresSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

