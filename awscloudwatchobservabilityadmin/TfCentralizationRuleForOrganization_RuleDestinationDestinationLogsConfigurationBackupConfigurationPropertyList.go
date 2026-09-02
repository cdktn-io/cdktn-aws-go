package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList interface {
	cdktn.ComplexList
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList
type jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList {
	_init_.Initialize()

	if err := validateNewTfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList_Override(t TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) Get(index *float64) TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

