package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSource_DataSourceConfigurationPropertyOutputReference interface {
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
	ConfluenceConfiguration() AwsDataSource_ConfluenceConfigurationPropertyList
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
	ManagedKnowledgeBaseConnectorConfiguration() AwsDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList
	// Experimental.
	ManagedKnowledgeBaseConnectorConfigurationInput() interface{}
	// Experimental.
	S3Configuration() AwsDataSource_S3ConfigurationPropertyList
	// Experimental.
	S3ConfigurationInput() interface{}
	// Experimental.
	SalesforceConfiguration() AwsDataSource_SalesforceConfigurationPropertyList
	// Experimental.
	SalesforceConfigurationInput() interface{}
	// Experimental.
	SharePointConfiguration() AwsDataSource_SharePointConfigurationPropertyList
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
	WebConfiguration() AwsDataSource_WebConfigurationPropertyList
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

// The jsii proxy struct for AwsDataSource_DataSourceConfigurationPropertyOutputReference
type jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ConfluenceConfiguration() AwsDataSource_ConfluenceConfigurationPropertyList {
	var returns AwsDataSource_ConfluenceConfigurationPropertyList
	_jsii_.Get(
		j,
		"confluenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ConfluenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ManagedKnowledgeBaseConnectorConfiguration() AwsDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList {
	var returns AwsDataSource_ManagedKnowledgeBaseConnectorConfigurationPropertyList
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ManagedKnowledgeBaseConnectorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) S3Configuration() AwsDataSource_S3ConfigurationPropertyList {
	var returns AwsDataSource_S3ConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) SalesforceConfiguration() AwsDataSource_SalesforceConfigurationPropertyList {
	var returns AwsDataSource_SalesforceConfigurationPropertyList
	_jsii_.Get(
		j,
		"salesforceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) SalesforceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) SharePointConfiguration() AwsDataSource_SharePointConfigurationPropertyList {
	var returns AwsDataSource_SharePointConfigurationPropertyList
	_jsii_.Get(
		j,
		"sharePointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) SharePointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) WebConfiguration() AwsDataSource_WebConfigurationPropertyList {
	var returns AwsDataSource_WebConfigurationPropertyList
	_jsii_.Get(
		j,
		"webConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) WebConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSource_DataSourceConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataSource_DataSourceConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSource_DataSourceConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsDataSource.DataSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSource_DataSourceConfigurationPropertyOutputReference_Override(a AwsDataSource_DataSourceConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsDataSource.DataSourceConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutConfluenceConfiguration(value interface{}) {
	if err := a.validatePutConfluenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfluenceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutManagedKnowledgeBaseConnectorConfiguration(value interface{}) {
	if err := a.validatePutManagedKnowledgeBaseConnectorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedKnowledgeBaseConnectorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutS3Configuration(value interface{}) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutSalesforceConfiguration(value interface{}) {
	if err := a.validatePutSalesforceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutSharePointConfiguration(value interface{}) {
	if err := a.validatePutSharePointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSharePointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) PutWebConfiguration(value interface{}) {
	if err := a.validatePutWebConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetConfluenceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetConfluenceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetManagedKnowledgeBaseConnectorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedKnowledgeBaseConnectorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetSalesforceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetSharePointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSharePointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ResetWebConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetWebConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSource_DataSourceConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

