package athena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/athena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/athena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkgroup_ConfigurationPropertyOutputReference interface {
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
	CustomerContentEncryptionConfiguration() AwsWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	// Experimental.
	CustomerContentEncryptionConfigurationInput() *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty
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
	EngineVersion() AwsWorkgroup_EngineVersionPropertyOutputReference
	// Experimental.
	EngineVersionInput() *AwsWorkgroup_EngineVersionProperty
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IdentityCenterConfiguration() AwsWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	// Experimental.
	IdentityCenterConfigurationInput() *AwsWorkgroup_IdentityCenterConfigurationProperty
	// Experimental.
	InternalValue() *AwsWorkgroup_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsWorkgroup_ConfigurationProperty)
	// Experimental.
	ManagedQueryResultsConfiguration() AwsWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	// Experimental.
	ManagedQueryResultsConfigurationInput() *AwsWorkgroup_ManagedQueryResultsConfigurationProperty
	// Experimental.
	MonitoringConfiguration() AwsWorkgroup_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *AwsWorkgroup_MonitoringConfigurationProperty
	// Experimental.
	PublishCloudwatchMetricsEnabled() interface{}
	// Experimental.
	SetPublishCloudwatchMetricsEnabled(val interface{})
	// Experimental.
	PublishCloudwatchMetricsEnabledInput() interface{}
	// Experimental.
	QueryResultsS3AccessGrantsConfiguration() AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	// Experimental.
	QueryResultsS3AccessGrantsConfigurationInput() *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	// Experimental.
	RequesterPaysEnabled() interface{}
	// Experimental.
	SetRequesterPaysEnabled(val interface{})
	// Experimental.
	RequesterPaysEnabledInput() interface{}
	// Experimental.
	ResultConfiguration() AwsWorkgroup_ResultConfigurationPropertyOutputReference
	// Experimental.
	ResultConfigurationInput() *AwsWorkgroup_ResultConfigurationProperty
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
	PutCustomerContentEncryptionConfiguration(value *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty)
	// Experimental.
	PutEngineVersion(value *AwsWorkgroup_EngineVersionProperty)
	// Experimental.
	PutIdentityCenterConfiguration(value *AwsWorkgroup_IdentityCenterConfigurationProperty)
	// Experimental.
	PutManagedQueryResultsConfiguration(value *AwsWorkgroup_ManagedQueryResultsConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *AwsWorkgroup_MonitoringConfigurationProperty)
	// Experimental.
	PutQueryResultsS3AccessGrantsConfiguration(value *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty)
	// Experimental.
	PutResultConfiguration(value *AwsWorkgroup_ResultConfigurationProperty)
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

// The jsii proxy struct for AwsWorkgroup_ConfigurationPropertyOutputReference
type jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfiguration() AwsWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfigurationInput() *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty {
	var returns *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"customerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EngineVersion() AwsWorkgroup_EngineVersionPropertyOutputReference {
	var returns AwsWorkgroup_EngineVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) EngineVersionInput() *AwsWorkgroup_EngineVersionProperty {
	var returns *AwsWorkgroup_EngineVersionProperty
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfiguration() AwsWorkgroup_IdentityCenterConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfigurationInput() *AwsWorkgroup_IdentityCenterConfigurationProperty {
	var returns *AwsWorkgroup_IdentityCenterConfigurationProperty
	_jsii_.Get(
		j,
		"identityCenterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) InternalValue() *AwsWorkgroup_ConfigurationProperty {
	var returns *AwsWorkgroup_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfiguration() AwsWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfigurationInput() *AwsWorkgroup_ManagedQueryResultsConfigurationProperty {
	var returns *AwsWorkgroup_ManagedQueryResultsConfigurationProperty
	_jsii_.Get(
		j,
		"managedQueryResultsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfiguration() AwsWorkgroup_MonitoringConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfigurationInput() *AwsWorkgroup_MonitoringConfigurationProperty {
	var returns *AwsWorkgroup_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfiguration() AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfigurationInput() *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty {
	var returns *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResultConfiguration() AwsWorkgroup_ResultConfigurationPropertyOutputReference {
	var returns AwsWorkgroup_ResultConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResultConfigurationInput() *AwsWorkgroup_ResultConfigurationProperty {
	var returns *AwsWorkgroup_ResultConfigurationProperty
	_jsii_.Get(
		j,
		"resultConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkgroup_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsWorkgroup_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkgroup_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.AwsWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkgroup_ConfigurationPropertyOutputReference_Override(a AwsWorkgroup_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.AwsWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetEnableMinimumEncryptionConfiguration(val interface{}) {
	if err := j.validateSetEnableMinimumEncryptionConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMinimumEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetEnforceWorkgroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkgroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkgroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetInternalValue(val *AwsWorkgroup_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutCustomerContentEncryptionConfiguration(value *AwsWorkgroup_CustomerContentEncryptionConfigurationProperty) {
	if err := a.validatePutCustomerContentEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerContentEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutEngineVersion(value *AwsWorkgroup_EngineVersionProperty) {
	if err := a.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutIdentityCenterConfiguration(value *AwsWorkgroup_IdentityCenterConfigurationProperty) {
	if err := a.validatePutIdentityCenterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentityCenterConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutManagedQueryResultsConfiguration(value *AwsWorkgroup_ManagedQueryResultsConfigurationProperty) {
	if err := a.validatePutManagedQueryResultsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedQueryResultsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *AwsWorkgroup_MonitoringConfigurationProperty) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutQueryResultsS3AccessGrantsConfiguration(value *AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty) {
	if err := a.validatePutQueryResultsS3AccessGrantsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryResultsS3AccessGrantsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) PutResultConfiguration(value *AwsWorkgroup_ResultConfigurationProperty) {
	if err := a.validatePutResultConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResultConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetEnableMinimumEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMinimumEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetEnforceWorkgroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceWorkgroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetIdentityCenterConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityCenterConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetManagedQueryResultsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedQueryResultsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetQueryResultsS3AccessGrantsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryResultsS3AccessGrantsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ResetResultConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetResultConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkgroup_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

