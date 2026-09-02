package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsathena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsathena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWorkgroup_ConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BytesScannedCutoffPerQuery() *float64
	// Experimental.
	SetBytesScannedCutoffPerQuery(val *float64)
	// Experimental.
	BytesScannedCutoffPerQueryInput() *float64
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
	CustomerContentEncryptionConfiguration() TfWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	// Experimental.
	CustomerContentEncryptionConfigurationInput() *TfWorkgroup_CustomerContentEncryptionConfigurationProperty
	// Experimental.
	EnableMinimumEncryptionConfiguration() interface{}
	// Experimental.
	SetEnableMinimumEncryptionConfiguration(val interface{})
	// Experimental.
	EnableMinimumEncryptionConfigurationInput() interface{}
	// Experimental.
	EnforceWorkgroupConfiguration() interface{}
	// Experimental.
	SetEnforceWorkgroupConfiguration(val interface{})
	// Experimental.
	EnforceWorkgroupConfigurationInput() interface{}
	// Experimental.
	EngineVersion() TfWorkgroup_EngineVersionPropertyOutputReference
	// Experimental.
	EngineVersionInput() *TfWorkgroup_EngineVersionProperty
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IdentityCenterConfiguration() TfWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	// Experimental.
	IdentityCenterConfigurationInput() *TfWorkgroup_IdentityCenterConfigurationProperty
	// Experimental.
	InternalValue() *TfWorkgroup_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfWorkgroup_ConfigurationProperty)
	// Experimental.
	ManagedQueryResultsConfiguration() TfWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	// Experimental.
	ManagedQueryResultsConfigurationInput() *TfWorkgroup_ManagedQueryResultsConfigurationProperty
	// Experimental.
	MonitoringConfiguration() TfWorkgroup_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *TfWorkgroup_MonitoringConfigurationProperty
	// Experimental.
	PublishCloudwatchMetricsEnabled() interface{}
	// Experimental.
	SetPublishCloudwatchMetricsEnabled(val interface{})
	// Experimental.
	PublishCloudwatchMetricsEnabledInput() interface{}
	// Experimental.
	QueryResultsS3AccessGrantsConfiguration() TfWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	// Experimental.
	QueryResultsS3AccessGrantsConfigurationInput() *TfWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	// Experimental.
	RequesterPaysEnabled() interface{}
	// Experimental.
	SetRequesterPaysEnabled(val interface{})
	// Experimental.
	RequesterPaysEnabledInput() interface{}
	// Experimental.
	ResultConfiguration() TfWorkgroup_ResultConfigurationPropertyOutputReference
	// Experimental.
	ResultConfigurationInput() *TfWorkgroup_ResultConfigurationProperty
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
	PutCustomerContentEncryptionConfiguration(value *TfWorkgroup_CustomerContentEncryptionConfigurationProperty)
	// Experimental.
	PutEngineVersion(value *TfWorkgroup_EngineVersionProperty)
	// Experimental.
	PutIdentityCenterConfiguration(value *TfWorkgroup_IdentityCenterConfigurationProperty)
	// Experimental.
	PutManagedQueryResultsConfiguration(value *TfWorkgroup_ManagedQueryResultsConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *TfWorkgroup_MonitoringConfigurationProperty)
	// Experimental.
	PutQueryResultsS3AccessGrantsConfiguration(value *TfWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty)
	// Experimental.
	PutResultConfiguration(value *TfWorkgroup_ResultConfigurationProperty)
	// Experimental.
	ResetBytesScannedCutoffPerQuery()
	// Experimental.
	ResetCustomerContentEncryptionConfiguration()
	// Experimental.
	ResetEnableMinimumEncryptionConfiguration()
	// Experimental.
	ResetEnforceWorkgroupConfiguration()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetExecutionRole()
	// Experimental.
	ResetIdentityCenterConfiguration()
	// Experimental.
	ResetManagedQueryResultsConfiguration()
	// Experimental.
	ResetMonitoringConfiguration()
	// Experimental.
	ResetPublishCloudwatchMetricsEnabled()
	// Experimental.
	ResetQueryResultsS3AccessGrantsConfiguration()
	// Experimental.
	ResetRequesterPaysEnabled()
	// Experimental.
	ResetResultConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWorkgroup_ConfigurationPropertyOutputReference
type jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfiguration() TfWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference {
	var returns TfWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfigurationInput() *TfWorkgroup_CustomerContentEncryptionConfigurationProperty {
	var returns *TfWorkgroup_CustomerContentEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"customerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EngineVersion() TfWorkgroup_EngineVersionPropertyOutputReference {
	var returns TfWorkgroup_EngineVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) EngineVersionInput() *TfWorkgroup_EngineVersionProperty {
	var returns *TfWorkgroup_EngineVersionProperty
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfiguration() TfWorkgroup_IdentityCenterConfigurationPropertyOutputReference {
	var returns TfWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfigurationInput() *TfWorkgroup_IdentityCenterConfigurationProperty {
	var returns *TfWorkgroup_IdentityCenterConfigurationProperty
	_jsii_.Get(
		j,
		"identityCenterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) InternalValue() *TfWorkgroup_ConfigurationProperty {
	var returns *TfWorkgroup_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfiguration() TfWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference {
	var returns TfWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfigurationInput() *TfWorkgroup_ManagedQueryResultsConfigurationProperty {
	var returns *TfWorkgroup_ManagedQueryResultsConfigurationProperty
	_jsii_.Get(
		j,
		"managedQueryResultsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfiguration() TfWorkgroup_MonitoringConfigurationPropertyOutputReference {
	var returns TfWorkgroup_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfigurationInput() *TfWorkgroup_MonitoringConfigurationProperty {
	var returns *TfWorkgroup_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfiguration() TfWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference {
	var returns TfWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfigurationInput() *TfWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty {
	var returns *TfWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResultConfiguration() TfWorkgroup_ResultConfigurationPropertyOutputReference {
	var returns TfWorkgroup_ResultConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResultConfigurationInput() *TfWorkgroup_ResultConfigurationProperty {
	var returns *TfWorkgroup_ResultConfigurationProperty
	_jsii_.Get(
		j,
		"resultConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWorkgroup_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWorkgroup_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWorkgroup_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWorkgroup_ConfigurationPropertyOutputReference_Override(t TfWorkgroup_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.TfWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetEnableMinimumEncryptionConfiguration(val interface{}) {
	if err := j.validateSetEnableMinimumEncryptionConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMinimumEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetEnforceWorkgroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkgroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkgroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetInternalValue(val *TfWorkgroup_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutCustomerContentEncryptionConfiguration(value *TfWorkgroup_CustomerContentEncryptionConfigurationProperty) {
	if err := t.validatePutCustomerContentEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomerContentEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutEngineVersion(value *TfWorkgroup_EngineVersionProperty) {
	if err := t.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutIdentityCenterConfiguration(value *TfWorkgroup_IdentityCenterConfigurationProperty) {
	if err := t.validatePutIdentityCenterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIdentityCenterConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutManagedQueryResultsConfiguration(value *TfWorkgroup_ManagedQueryResultsConfigurationProperty) {
	if err := t.validatePutManagedQueryResultsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedQueryResultsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *TfWorkgroup_MonitoringConfigurationProperty) {
	if err := t.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutQueryResultsS3AccessGrantsConfiguration(value *TfWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty) {
	if err := t.validatePutQueryResultsS3AccessGrantsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryResultsS3AccessGrantsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) PutResultConfiguration(value *TfWorkgroup_ResultConfigurationProperty) {
	if err := t.validatePutResultConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResultConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		t,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetEnableMinimumEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableMinimumEncryptionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetEnforceWorkgroupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetEnforceWorkgroupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetIdentityCenterConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityCenterConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetManagedQueryResultsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedQueryResultsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetQueryResultsS3AccessGrantsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryResultsS3AccessGrantsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ResetResultConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetResultConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWorkgroup_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

