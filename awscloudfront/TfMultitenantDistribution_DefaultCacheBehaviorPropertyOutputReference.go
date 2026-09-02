package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedMethods() TfMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsPropertyList
	// Experimental.
	AllowedMethodsInput() interface{}
	// Experimental.
	CachePolicyId() *string
	// Experimental.
	SetCachePolicyId(val *string)
	// Experimental.
	CachePolicyIdInput() *string
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
	Compress() interface{}
	// Experimental.
	SetCompress(val interface{})
	// Experimental.
	CompressInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FieldLevelEncryptionId() *string
	// Experimental.
	SetFieldLevelEncryptionId(val *string)
	// Experimental.
	FieldLevelEncryptionIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FunctionAssociation() TfMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList
	// Experimental.
	FunctionAssociationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaFunctionAssociation() TfMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList
	// Experimental.
	LambdaFunctionAssociationInput() interface{}
	// Experimental.
	OriginRequestPolicyId() *string
	// Experimental.
	SetOriginRequestPolicyId(val *string)
	// Experimental.
	OriginRequestPolicyIdInput() *string
	// Experimental.
	RealtimeLogConfigArn() *string
	// Experimental.
	SetRealtimeLogConfigArn(val *string)
	// Experimental.
	RealtimeLogConfigArnInput() *string
	// Experimental.
	ResponseHeadersPolicyId() *string
	// Experimental.
	SetResponseHeadersPolicyId(val *string)
	// Experimental.
	ResponseHeadersPolicyIdInput() *string
	// Experimental.
	TargetOriginId() *string
	// Experimental.
	SetTargetOriginId(val *string)
	// Experimental.
	TargetOriginIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrustedKeyGroups() TfMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsPropertyList
	// Experimental.
	TrustedKeyGroupsInput() interface{}
	// Experimental.
	ViewerProtocolPolicy() *string
	// Experimental.
	SetViewerProtocolPolicy(val *string)
	// Experimental.
	ViewerProtocolPolicyInput() *string
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
	PutAllowedMethods(value interface{})
	// Experimental.
	PutFunctionAssociation(value interface{})
	// Experimental.
	PutLambdaFunctionAssociation(value interface{})
	// Experimental.
	PutTrustedKeyGroups(value interface{})
	// Experimental.
	ResetAllowedMethods()
	// Experimental.
	ResetCachePolicyId()
	// Experimental.
	ResetCompress()
	// Experimental.
	ResetFieldLevelEncryptionId()
	// Experimental.
	ResetFunctionAssociation()
	// Experimental.
	ResetLambdaFunctionAssociation()
	// Experimental.
	ResetOriginRequestPolicyId()
	// Experimental.
	ResetRealtimeLogConfigArn()
	// Experimental.
	ResetResponseHeadersPolicyId()
	// Experimental.
	ResetTrustedKeyGroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference
type jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) AllowedMethods() TfMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsPropertyList {
	var returns TfMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsPropertyList
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) AllowedMethodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) FunctionAssociation() TfMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList {
	var returns TfMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) FunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) LambdaFunctionAssociation() TfMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList {
	var returns TfMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) LambdaFunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedKeyGroups() TfMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsPropertyList {
	var returns TfMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedKeyGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfMultitenantDistribution.DefaultCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference_Override(t TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfMultitenantDistribution.DefaultCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) PutAllowedMethods(value interface{}) {
	if err := t.validatePutAllowedMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllowedMethods",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) PutFunctionAssociation(value interface{}) {
	if err := t.validatePutFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFunctionAssociation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) PutLambdaFunctionAssociation(value interface{}) {
	if err := t.validatePutLambdaFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionAssociation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) PutTrustedKeyGroups(value interface{}) {
	if err := t.validatePutTrustedKeyGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrustedKeyGroups",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		t,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		t,
		"resetCompress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		t,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		t,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		t,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

