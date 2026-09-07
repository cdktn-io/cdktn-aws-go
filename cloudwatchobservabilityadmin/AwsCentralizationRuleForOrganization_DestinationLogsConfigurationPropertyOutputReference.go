package cloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BackupConfiguration() AwsCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList
	// Experimental.
	BackupConfigurationInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogGroupNameConfiguration() AwsCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList
	// Experimental.
	LogGroupNameConfigurationInput() interface{}
	// Experimental.
	LogsEncryptionConfiguration() AwsCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList
	// Experimental.
	LogsEncryptionConfigurationInput() interface{}
	// Experimental.
	TagPropagationConfiguration() AwsCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList
	// Experimental.
	TagPropagationConfigurationInput() interface{}
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
	PutBackupConfiguration(value interface{})
	// Experimental.
	PutLogGroupNameConfiguration(value interface{})
	// Experimental.
	PutLogsEncryptionConfiguration(value interface{})
	// Experimental.
	PutTagPropagationConfiguration(value interface{})
	// Experimental.
	ResetBackupConfiguration()
	// Experimental.
	ResetLogGroupNameConfiguration()
	// Experimental.
	ResetLogsEncryptionConfiguration()
	// Experimental.
	ResetTagPropagationConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference
type jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) BackupConfiguration() AwsCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList {
	var returns AwsCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) BackupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogGroupNameConfiguration() AwsCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList {
	var returns AwsCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList
	_jsii_.Get(
		j,
		"logGroupNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogGroupNameConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logGroupNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogsEncryptionConfiguration() AwsCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList {
	var returns AwsCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList
	_jsii_.Get(
		j,
		"logsEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogsEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TagPropagationConfiguration() AwsCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList {
	var returns AwsCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList
	_jsii_.Get(
		j,
		"tagPropagationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TagPropagationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPropagationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsCentralizationRuleForOrganization.DestinationLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference_Override(a AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsCentralizationRuleForOrganization.DestinationLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutBackupConfiguration(value interface{}) {
	if err := a.validatePutBackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackupConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutLogGroupNameConfiguration(value interface{}) {
	if err := a.validatePutLogGroupNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogGroupNameConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutLogsEncryptionConfiguration(value interface{}) {
	if err := a.validatePutLogsEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogsEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutTagPropagationConfiguration(value interface{}) {
	if err := a.validatePutTagPropagationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagPropagationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetBackupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetLogGroupNameConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogGroupNameConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetLogsEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogsEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetTagPropagationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTagPropagationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

