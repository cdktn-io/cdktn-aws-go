package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BackupConfiguration() TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList
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
	LogGroupNameConfiguration() TfCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList
	// Experimental.
	LogGroupNameConfigurationInput() interface{}
	// Experimental.
	LogsEncryptionConfiguration() TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList
	// Experimental.
	LogsEncryptionConfigurationInput() interface{}
	// Experimental.
	TagPropagationConfiguration() TfCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList
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

// The jsii proxy struct for TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference
type jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) BackupConfiguration() TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList {
	var returns TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) BackupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogGroupNameConfiguration() TfCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList {
	var returns TfCentralizationRuleForOrganization_LogGroupNameConfigurationPropertyList
	_jsii_.Get(
		j,
		"logGroupNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogGroupNameConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logGroupNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogsEncryptionConfiguration() TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList {
	var returns TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyList
	_jsii_.Get(
		j,
		"logsEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) LogsEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TagPropagationConfiguration() TfCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList {
	var returns TfCentralizationRuleForOrganization_TagPropagationConfigurationPropertyList
	_jsii_.Get(
		j,
		"tagPropagationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TagPropagationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagPropagationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.DestinationLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference_Override(t TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.DestinationLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutBackupConfiguration(value interface{}) {
	if err := t.validatePutBackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBackupConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutLogGroupNameConfiguration(value interface{}) {
	if err := t.validatePutLogGroupNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogGroupNameConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutLogsEncryptionConfiguration(value interface{}) {
	if err := t.validatePutLogsEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogsEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) PutTagPropagationConfiguration(value interface{}) {
	if err := t.validatePutTagPropagationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagPropagationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetBackupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetLogGroupNameConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLogGroupNameConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetLogsEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLogsEncryptionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ResetTagPropagationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTagPropagationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_DestinationLogsConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

