package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGateway_CustomJwtAuthorizerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedAudience() *[]*string
	// Experimental.
	SetAllowedAudience(val *[]*string)
	// Experimental.
	AllowedAudienceInput() *[]*string
	// Experimental.
	AllowedClients() *[]*string
	// Experimental.
	SetAllowedClients(val *[]*string)
	// Experimental.
	AllowedClientsInput() *[]*string
	// Experimental.
	AllowedScopes() *[]*string
	// Experimental.
	SetAllowedScopes(val *[]*string)
	// Experimental.
	AllowedScopesInput() *[]*string
	// Experimental.
	AllowedWorkloadConfiguration() TfGateway_AllowedWorkloadConfigurationPropertyList
	// Experimental.
	AllowedWorkloadConfigurationInput() interface{}
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
	CustomClaim() TfGateway_CustomClaimPropertyList
	// Experimental.
	CustomClaimInput() interface{}
	// Experimental.
	DiscoveryUrl() *string
	// Experimental.
	SetDiscoveryUrl(val *string)
	// Experimental.
	DiscoveryUrlInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PrivateEndpoint() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyList
	// Experimental.
	PrivateEndpointInput() interface{}
	// Experimental.
	PrivateEndpointOverrides() TfGateway_PrivateEndpointOverridesPropertyList
	// Experimental.
	PrivateEndpointOverridesInput() interface{}
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
	PutAllowedWorkloadConfiguration(value interface{})
	// Experimental.
	PutCustomClaim(value interface{})
	// Experimental.
	PutPrivateEndpoint(value interface{})
	// Experimental.
	PutPrivateEndpointOverrides(value interface{})
	// Experimental.
	ResetAllowedAudience()
	// Experimental.
	ResetAllowedClients()
	// Experimental.
	ResetAllowedScopes()
	// Experimental.
	ResetAllowedWorkloadConfiguration()
	// Experimental.
	ResetCustomClaim()
	// Experimental.
	ResetPrivateEndpoint()
	// Experimental.
	ResetPrivateEndpointOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGateway_CustomJwtAuthorizerPropertyOutputReference
type jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedAudience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedAudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedWorkloadConfiguration() TfGateway_AllowedWorkloadConfigurationPropertyList {
	var returns TfGateway_AllowedWorkloadConfigurationPropertyList
	_jsii_.Get(
		j,
		"allowedWorkloadConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) AllowedWorkloadConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedWorkloadConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) CustomClaim() TfGateway_CustomClaimPropertyList {
	var returns TfGateway_CustomClaimPropertyList
	_jsii_.Get(
		j,
		"customClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) CustomClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) DiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PrivateEndpoint() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyList {
	var returns TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyList
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PrivateEndpointOverrides() TfGateway_PrivateEndpointOverridesPropertyList {
	var returns TfGateway_PrivateEndpointOverridesPropertyList
	_jsii_.Get(
		j,
		"privateEndpointOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PrivateEndpointOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGateway_CustomJwtAuthorizerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGateway_CustomJwtAuthorizerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGateway_CustomJwtAuthorizerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.CustomJwtAuthorizerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGateway_CustomJwtAuthorizerPropertyOutputReference_Override(t TfGateway_CustomJwtAuthorizerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.CustomJwtAuthorizerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetAllowedAudience(val *[]*string) {
	if err := j.validateSetAllowedAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAudience",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetAllowedClients(val *[]*string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetAllowedScopes(val *[]*string) {
	if err := j.validateSetAllowedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedScopes",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetDiscoveryUrl(val *string) {
	if err := j.validateSetDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryUrl",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PutAllowedWorkloadConfiguration(value interface{}) {
	if err := t.validatePutAllowedWorkloadConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllowedWorkloadConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PutCustomClaim(value interface{}) {
	if err := t.validatePutCustomClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomClaim",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PutPrivateEndpoint(value interface{}) {
	if err := t.validatePutPrivateEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrivateEndpoint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) PutPrivateEndpointOverrides(value interface{}) {
	if err := t.validatePutPrivateEndpointOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrivateEndpointOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetAllowedAudience() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedAudience",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetAllowedScopes() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedScopes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetAllowedWorkloadConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedWorkloadConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetCustomClaim() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomClaim",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetPrivateEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ResetPrivateEndpointOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateEndpointOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGateway_CustomJwtAuthorizerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

