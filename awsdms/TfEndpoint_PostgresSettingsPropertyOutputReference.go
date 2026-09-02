package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_PostgresSettingsPropertyOutputReference interface {
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
	InternalValue() *TfEndpoint_PostgresSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_PostgresSettingsProperty)
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

// The jsii proxy struct for TfEndpoint_PostgresSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) AfterConnectScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) AfterConnectScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterConnectScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) AuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) AuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) BabelfishDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) BabelfishDatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"babelfishDatabaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) CaptureDdls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) CaptureDdlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureDdlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) DatabaseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) DatabaseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) DdlArtifactsSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) DdlArtifactsSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddlArtifactsSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ExecuteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ExecuteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) FailTasksOnLobTruncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) FailTasksOnLobTruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTasksOnLobTruncationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) HeartbeatSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"heartbeatSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) InternalValue() *TfEndpoint_PostgresSettingsProperty {
	var returns *TfEndpoint_PostgresSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapBooleanAsBoolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapBooleanAsBooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapBooleanAsBooleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapJsonbAsClob() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapJsonbAsClobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapJsonbAsClobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapLongVarcharAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MapLongVarcharAsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapLongVarcharAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) SlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) SlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_PostgresSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_PostgresSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_PostgresSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.PostgresSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_PostgresSettingsPropertyOutputReference_Override(t TfEndpoint_PostgresSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.PostgresSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetAfterConnectScript(val *string) {
	if err := j.validateSetAfterConnectScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterConnectScript",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetAuthenticationMethod(val *string) {
	if err := j.validateSetAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMethod",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetBabelfishDatabaseName(val *string) {
	if err := j.validateSetBabelfishDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"babelfishDatabaseName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetCaptureDdls(val interface{}) {
	if err := j.validateSetCaptureDdlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captureDdls",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetDatabaseMode(val *string) {
	if err := j.validateSetDatabaseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseMode",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetDdlArtifactsSchema(val *string) {
	if err := j.validateSetDdlArtifactsSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddlArtifactsSchema",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetExecuteTimeout(val *float64) {
	if err := j.validateSetExecuteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executeTimeout",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetFailTasksOnLobTruncation(val interface{}) {
	if err := j.validateSetFailTasksOnLobTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failTasksOnLobTruncation",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatEnable(val interface{}) {
	if err := j.validateSetHeartbeatEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatEnable",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatFrequency(val *float64) {
	if err := j.validateSetHeartbeatFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatFrequency",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetHeartbeatSchema(val *string) {
	if err := j.validateSetHeartbeatSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatSchema",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_PostgresSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetMapBooleanAsBoolean(val interface{}) {
	if err := j.validateSetMapBooleanAsBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBooleanAsBoolean",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetMapJsonbAsClob(val interface{}) {
	if err := j.validateSetMapJsonbAsClobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapJsonbAsClob",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetMapLongVarcharAs(val *string) {
	if err := j.validateSetMapLongVarcharAsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapLongVarcharAs",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetSlotName(val *string) {
	if err := j.validateSetSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetAfterConnectScript() {
	_jsii_.InvokeVoid(
		t,
		"resetAfterConnectScript",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetAuthenticationMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetBabelfishDatabaseName() {
	_jsii_.InvokeVoid(
		t,
		"resetBabelfishDatabaseName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetCaptureDdls() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptureDdls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetDatabaseMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetDdlArtifactsSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetDdlArtifactsSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetExecuteTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetExecuteTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetFailTasksOnLobTruncation() {
	_jsii_.InvokeVoid(
		t,
		"resetFailTasksOnLobTruncation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatEnable() {
	_jsii_.InvokeVoid(
		t,
		"resetHeartbeatEnable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatFrequency() {
	_jsii_.InvokeVoid(
		t,
		"resetHeartbeatFrequency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetHeartbeatSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetHeartbeatSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetMapBooleanAsBoolean() {
	_jsii_.InvokeVoid(
		t,
		"resetMapBooleanAsBoolean",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetMapJsonbAsClob() {
	_jsii_.InvokeVoid(
		t,
		"resetMapJsonbAsClob",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetMapLongVarcharAs() {
	_jsii_.InvokeVoid(
		t,
		"resetMapLongVarcharAs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		t,
		"resetPluginName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ResetSlotName() {
	_jsii_.InvokeVoid(
		t,
		"resetSlotName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_PostgresSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

