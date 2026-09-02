package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DataSourceSelectionCriteria() *string
	// Experimental.
	SetDataSourceSelectionCriteria(val *string)
	// Experimental.
	DataSourceSelectionCriteriaInput() *string
	// Experimental.
	EncryptedLogGroupStrategy() *string
	// Experimental.
	SetEncryptedLogGroupStrategy(val *string)
	// Experimental.
	EncryptedLogGroupStrategyInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogGroupSelectionCriteria() *string
	// Experimental.
	SetLogGroupSelectionCriteria(val *string)
	// Experimental.
	LogGroupSelectionCriteriaInput() *string
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
	ResetDataSourceSelectionCriteria()
	// Experimental.
	ResetLogGroupSelectionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference
type jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) DataSourceSelectionCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) DataSourceSelectionCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) EncryptedLogGroupStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedLogGroupStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) EncryptedLogGroupStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedLogGroupStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) LogGroupSelectionCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupSelectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) LogGroupSelectionCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupSelectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.SourceLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference_Override(t TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.SourceLogsConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetDataSourceSelectionCriteria(val *string) {
	if err := j.validateSetDataSourceSelectionCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceSelectionCriteria",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetEncryptedLogGroupStrategy(val *string) {
	if err := j.validateSetEncryptedLogGroupStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptedLogGroupStrategy",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetLogGroupSelectionCriteria(val *string) {
	if err := j.validateSetLogGroupSelectionCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupSelectionCriteria",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ResetDataSourceSelectionCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSourceSelectionCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ResetLogGroupSelectionCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetLogGroupSelectionCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_SourceLogsConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

