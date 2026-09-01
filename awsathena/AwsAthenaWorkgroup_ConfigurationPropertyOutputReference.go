package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsathena/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsathena/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAthenaWorkgroup_ConfigurationPropertyOutputReference interface {
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
	CustomerContentEncryptionConfiguration() AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	// Experimental.
	CustomerContentEncryptionConfigurationInput() *AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationProperty
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
	EngineVersion() AwsAthenaWorkgroup_EngineVersionPropertyOutputReference
	// Experimental.
	EngineVersionInput() *AwsAthenaWorkgroup_EngineVersionProperty
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IdentityCenterConfiguration() AwsAthenaWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	// Experimental.
	IdentityCenterConfigurationInput() *AwsAthenaWorkgroup_IdentityCenterConfigurationProperty
	// Experimental.
	InternalValue() *AwsAthenaWorkgroup_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsAthenaWorkgroup_ConfigurationProperty)
	// Experimental.
	ManagedQueryResultsConfiguration() AwsAthenaWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	// Experimental.
	ManagedQueryResultsConfigurationInput() *AwsAthenaWorkgroup_ManagedQueryResultsConfigurationProperty
	// Experimental.
	MonitoringConfiguration() AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *AwsAthenaWorkgroup_MonitoringConfigurationProperty
	// Experimental.
	PublishCloudwatchMetricsEnabled() interface{}
	// Experimental.
	SetPublishCloudwatchMetricsEnabled(val interface{})
	// Experimental.
	PublishCloudwatchMetricsEnabledInput() interface{}
	// Experimental.
	QueryResultsS3AccessGrantsConfiguration() AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	// Experimental.
	QueryResultsS3AccessGrantsConfigurationInput() *AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	// Experimental.
	RequesterPaysEnabled() interface{}
	// Experimental.
	SetRequesterPaysEnabled(val interface{})
	// Experimental.
	RequesterPaysEnabledInput() interface{}
	// Experimental.
	ResultConfiguration() AwsAthenaWorkgroup_ResultConfigurationPropertyOutputReference
	// Experimental.
	ResultConfigurationInput() *AwsAthenaWorkgroup_ResultConfigurationProperty
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
	PutCustomerContentEncryptionConfiguration(value *AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationProperty)
	// Experimental.
	PutEngineVersion(value *AwsAthenaWorkgroup_EngineVersionProperty)
	// Experimental.
	PutIdentityCenterConfiguration(value *AwsAthenaWorkgroup_IdentityCenterConfigurationProperty)
	// Experimental.
	PutManagedQueryResultsConfiguration(value *AwsAthenaWorkgroup_ManagedQueryResultsConfigurationProperty)
	// Experimental.
	PutMonitoringConfiguration(value *AwsAthenaWorkgroup_MonitoringConfigurationProperty)
	// Experimental.
	PutQueryResultsS3AccessGrantsConfiguration(value *AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty)
	// Experimental.
	PutResultConfiguration(value *AwsAthenaWorkgroup_ResultConfigurationProperty)
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

// The jsii proxy struct for AwsAthenaWorkgroup_ConfigurationPropertyOutputReference
type jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) BytesScannedCutoffPerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesScannedCutoffPerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfiguration() AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"customerContentEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) CustomerContentEncryptionConfigurationInput() *AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationProperty {
	var returns *AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"customerContentEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EnableMinimumEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMinimumEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EnforceWorkgroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceWorkgroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EngineVersion() AwsAthenaWorkgroup_EngineVersionPropertyOutputReference {
	var returns AwsAthenaWorkgroup_EngineVersionPropertyOutputReference
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) EngineVersionInput() *AwsAthenaWorkgroup_EngineVersionProperty {
	var returns *AwsAthenaWorkgroup_EngineVersionProperty
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfiguration() AwsAthenaWorkgroup_IdentityCenterConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_IdentityCenterConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) IdentityCenterConfigurationInput() *AwsAthenaWorkgroup_IdentityCenterConfigurationProperty {
	var returns *AwsAthenaWorkgroup_IdentityCenterConfigurationProperty
	_jsii_.Get(
		j,
		"identityCenterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) InternalValue() *AwsAthenaWorkgroup_ConfigurationProperty {
	var returns *AwsAthenaWorkgroup_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfiguration() AwsAthenaWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_ManagedQueryResultsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"managedQueryResultsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ManagedQueryResultsConfigurationInput() *AwsAthenaWorkgroup_ManagedQueryResultsConfigurationProperty {
	var returns *AwsAthenaWorkgroup_ManagedQueryResultsConfigurationProperty
	_jsii_.Get(
		j,
		"managedQueryResultsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfiguration() AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) MonitoringConfigurationInput() *AwsAthenaWorkgroup_MonitoringConfigurationProperty {
	var returns *AwsAthenaWorkgroup_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PublishCloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishCloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfiguration() AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) QueryResultsS3AccessGrantsConfigurationInput() *AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty {
	var returns *AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty
	_jsii_.Get(
		j,
		"queryResultsS3AccessGrantsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) RequesterPaysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResultConfiguration() AwsAthenaWorkgroup_ResultConfigurationPropertyOutputReference {
	var returns AwsAthenaWorkgroup_ResultConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"resultConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResultConfigurationInput() *AwsAthenaWorkgroup_ResultConfigurationProperty {
	var returns *AwsAthenaWorkgroup_ResultConfigurationProperty
	_jsii_.Get(
		j,
		"resultConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAthenaWorkgroup_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAthenaWorkgroup_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAthenaWorkgroup_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-athena.AwsAthenaWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAthenaWorkgroup_ConfigurationPropertyOutputReference_Override(a AwsAthenaWorkgroup_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-athena.AwsAthenaWorkgroup.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetBytesScannedCutoffPerQuery(val *float64) {
	if err := j.validateSetBytesScannedCutoffPerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesScannedCutoffPerQuery",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetEnableMinimumEncryptionConfiguration(val interface{}) {
	if err := j.validateSetEnableMinimumEncryptionConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMinimumEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetEnforceWorkgroupConfiguration(val interface{}) {
	if err := j.validateSetEnforceWorkgroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceWorkgroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetInternalValue(val *AwsAthenaWorkgroup_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetPublishCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetPublishCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishCloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetRequesterPaysEnabled(val interface{}) {
	if err := j.validateSetRequesterPaysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPaysEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutCustomerContentEncryptionConfiguration(value *AwsAthenaWorkgroup_CustomerContentEncryptionConfigurationProperty) {
	if err := a.validatePutCustomerContentEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerContentEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutEngineVersion(value *AwsAthenaWorkgroup_EngineVersionProperty) {
	if err := a.validatePutEngineVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEngineVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutIdentityCenterConfiguration(value *AwsAthenaWorkgroup_IdentityCenterConfigurationProperty) {
	if err := a.validatePutIdentityCenterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentityCenterConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutManagedQueryResultsConfiguration(value *AwsAthenaWorkgroup_ManagedQueryResultsConfigurationProperty) {
	if err := a.validatePutManagedQueryResultsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedQueryResultsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutMonitoringConfiguration(value *AwsAthenaWorkgroup_MonitoringConfigurationProperty) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutQueryResultsS3AccessGrantsConfiguration(value *AwsAthenaWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty) {
	if err := a.validatePutQueryResultsS3AccessGrantsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryResultsS3AccessGrantsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) PutResultConfiguration(value *AwsAthenaWorkgroup_ResultConfigurationProperty) {
	if err := a.validatePutResultConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResultConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetBytesScannedCutoffPerQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetBytesScannedCutoffPerQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetCustomerContentEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerContentEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetEnableMinimumEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMinimumEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetEnforceWorkgroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceWorkgroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetIdentityCenterConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityCenterConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetManagedQueryResultsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedQueryResultsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetPublishCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetQueryResultsS3AccessGrantsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryResultsS3AccessGrantsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetRequesterPaysEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPaysEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ResetResultConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetResultConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAthenaWorkgroup_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

