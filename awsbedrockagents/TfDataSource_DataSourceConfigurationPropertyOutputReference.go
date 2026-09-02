package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_DataSourceConfigurationPropertyOutputReference interface {
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
	ConfluenceConfiguration() TfDataSource_ConfluenceConfigurationPropertyList
	// Experimental.
	ConfluenceConfigurationInput() interface{}
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
	ManagedKnowledgeBaseConnectorConfiguration() TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList
	// Experimental.
	ManagedKnowledgeBaseConnectorConfigurationInput() interface{}
	// Experimental.
	S3Configuration() TfDataSource_S3ConfigurationPropertyList
	// Experimental.
	S3ConfigurationInput() interface{}
	// Experimental.
	SalesforceConfiguration() TfDataSource_SalesforceConfigurationPropertyList
	// Experimental.
	SalesforceConfigurationInput() interface{}
	// Experimental.
	SharePointConfiguration() TfDataSource_SharePointConfigurationPropertyList
	// Experimental.
	SharePointConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	WebConfiguration() TfDataSource_WebConfigurationPropertyList
	// Experimental.
	WebConfigurationInput() interface{}
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
	PutConfluenceConfiguration(value interface{})
	// Experimental.
	PutManagedKnowledgeBaseConnectorConfiguration(value interface{})
	// Experimental.
	PutS3Configuration(value interface{})
	// Experimental.
	PutSalesforceConfiguration(value interface{})
	// Experimental.
	PutSharePointConfiguration(value interface{})
	// Experimental.
	PutWebConfiguration(value interface{})
	// Experimental.
	ResetConfluenceConfiguration()
	// Experimental.
	ResetManagedKnowledgeBaseConnectorConfiguration()
	// Experimental.
	ResetS3Configuration()
	// Experimental.
	ResetSalesforceConfiguration()
	// Experimental.
	ResetSharePointConfiguration()
	// Experimental.
	ResetWebConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_DataSourceConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ConfluenceConfiguration() TfDataSource_ConfluenceConfigurationPropertyList {
	var returns TfDataSource_ConfluenceConfigurationPropertyList
	_jsii_.Get(
		j,
		"confluenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ConfluenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ManagedKnowledgeBaseConnectorConfiguration() TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList {
	var returns TfDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ManagedKnowledgeBaseConnectorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) S3Configuration() TfDataSource_S3ConfigurationPropertyList {
	var returns TfDataSource_S3ConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) SalesforceConfiguration() TfDataSource_SalesforceConfigurationPropertyList {
	var returns TfDataSource_SalesforceConfigurationPropertyList
	_jsii_.Get(
		j,
		"salesforceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) SalesforceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) SharePointConfiguration() TfDataSource_SharePointConfigurationPropertyList {
	var returns TfDataSource_SharePointConfigurationPropertyList
	_jsii_.Get(
		j,
		"sharePointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) SharePointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) WebConfiguration() TfDataSource_WebConfigurationPropertyList {
	var returns TfDataSource_WebConfigurationPropertyList
	_jsii_.Get(
		j,
		"webConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) WebConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_DataSourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSource_DataSourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_DataSourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.DataSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_DataSourceConfigurationPropertyOutputReference_Override(t TfDataSource_DataSourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.DataSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutConfluenceConfiguration(value interface{}) {
	if err := t.validatePutConfluenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfluenceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutManagedKnowledgeBaseConnectorConfiguration(value interface{}) {
	if err := t.validatePutManagedKnowledgeBaseConnectorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedKnowledgeBaseConnectorConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutS3Configuration(value interface{}) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutSalesforceConfiguration(value interface{}) {
	if err := t.validatePutSalesforceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSalesforceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutSharePointConfiguration(value interface{}) {
	if err := t.validatePutSharePointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSharePointConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) PutWebConfiguration(value interface{}) {
	if err := t.validatePutWebConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWebConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetConfluenceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfluenceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetManagedKnowledgeBaseConnectorConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedKnowledgeBaseConnectorConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetSalesforceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSalesforceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetSharePointConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSharePointConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ResetWebConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetWebConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_DataSourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

