package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference interface {
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
	EncryptionConflictResolutionStrategy() *string
	// Experimental.
	SetEncryptionConflictResolutionStrategy(val *string)
	// Experimental.
	EncryptionConflictResolutionStrategyInput() *string
	// Experimental.
	EncryptionScope() *string
	// Experimental.
	SetEncryptionScope(val *string)
	// Experimental.
	EncryptionScopeInput() *string
	// Experimental.
	EncryptionStrategy() *string
	// Experimental.
	SetEncryptionStrategy(val *string)
	// Experimental.
	EncryptionStrategyInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
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
	ResetEncryptionConflictResolutionStrategy()
	// Experimental.
	ResetEncryptionScope()
	// Experimental.
	ResetKmsKeyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference
type jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionConflictResolutionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionConflictResolutionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionConflictResolutionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionConflictResolutionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) EncryptionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.LogsEncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference_Override(t TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfCentralizationRuleForOrganization.LogsEncryptionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetEncryptionConflictResolutionStrategy(val *string) {
	if err := j.validateSetEncryptionConflictResolutionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionConflictResolutionStrategy",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetEncryptionScope(val *string) {
	if err := j.validateSetEncryptionScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionScope",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetEncryptionStrategy(val *string) {
	if err := j.validateSetEncryptionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionStrategy",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ResetEncryptionConflictResolutionStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionConflictResolutionStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ResetEncryptionScope() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionScope",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCentralizationRuleForOrganization_LogsEncryptionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

