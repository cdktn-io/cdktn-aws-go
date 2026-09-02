package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference interface {
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
	// Experimental.
	ConnectorParameters() *string
	// Experimental.
	SetConnectorParameters(val *string)
	// Experimental.
	ConnectorParametersInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeletionProtectionConfiguration() TfDataSource_DeletionProtectionConfigurationPropertyList
	// Experimental.
	DeletionProtectionConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MediaExtractionConfiguration() TfDataSource_MediaExtractionConfigurationPropertyList
	// Experimental.
	MediaExtractionConfigurationInput() interface{}
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
	PutDeletionProtectionConfiguration(value interface{})
	// Experimental.
	PutMediaExtractionConfiguration(value interface{})
	// Experimental.
	ResetConnectorParameters()
	// Experimental.
	ResetDeletionProtectionConfiguration()
	// Experimental.
	ResetMediaExtractionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ConnectorParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ConnectorParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) DeletionProtectionConfiguration() TfDataSource_DeletionProtectionConfigurationPropertyList {
	var returns TfDataSource_DeletionProtectionConfigurationPropertyList
	_jsii_.Get(
		j,
		"deletionProtectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) DeletionProtectionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) MediaExtractionConfiguration() TfDataSource_MediaExtractionConfigurationPropertyList {
	var returns TfDataSource_MediaExtractionConfigurationPropertyList
	_jsii_.Get(
		j,
		"mediaExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) MediaExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference_Override(t TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetConnectorParameters(val *string) {
	if err := j.validateSetConnectorParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorParameters",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) PutDeletionProtectionConfiguration(value interface{}) {
	if err := t.validatePutDeletionProtectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeletionProtectionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) PutMediaExtractionConfiguration(value interface{}) {
	if err := t.validatePutMediaExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMediaExtractionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ResetConnectorParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectorParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ResetDeletionProtectionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtectionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ResetMediaExtractionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMediaExtractionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

