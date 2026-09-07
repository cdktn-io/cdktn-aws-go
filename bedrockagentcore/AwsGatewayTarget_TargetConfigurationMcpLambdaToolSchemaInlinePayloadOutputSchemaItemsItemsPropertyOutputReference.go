package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference interface {
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ItemsJson() *string
	// Experimental.
	SetItemsJson(val *string)
	// Experimental.
	ItemsJsonInput() *string
	// Experimental.
	PropertiesJson() *string
	// Experimental.
	SetPropertiesJson(val *string)
	// Experimental.
	PropertiesJsonInput() *string
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
	ResetDescription()
	// Experimental.
	ResetItemsJson()
	// Experimental.
	ResetPropertiesJson()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference
type jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ItemsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ItemsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) PropertiesJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) PropertiesJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference_Override(a AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsGatewayTarget.TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetItemsJson(val *string) {
	if err := j.validateSetItemsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"itemsJson",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetPropertiesJson(val *string) {
	if err := j.validateSetPropertiesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertiesJson",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ResetItemsJson() {
	_jsii_.InvokeVoid(
		a,
		"resetItemsJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ResetPropertiesJson() {
	_jsii_.InvokeVoid(
		a,
		"resetPropertiesJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadOutputSchemaItemsItemsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

