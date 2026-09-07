package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDistribution_DefaultCacheBehaviorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedMethods() *[]*string
	// Experimental.
	SetAllowedMethods(val *[]*string)
	// Experimental.
	AllowedMethodsInput() *[]*string
	// Experimental.
	CachedMethods() *[]*string
	// Experimental.
	SetCachedMethods(val *[]*string)
	// Experimental.
	CachedMethodsInput() *[]*string
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
	DefaultTtl() *float64
	// Experimental.
	SetDefaultTtl(val *float64)
	// Experimental.
	DefaultTtlInput() *float64
	// Experimental.
	FieldLevelEncryptionId() *string
	// Experimental.
	SetFieldLevelEncryptionId(val *string)
	// Experimental.
	FieldLevelEncryptionIdInput() *string
	// Experimental.
	ForwardedValues() AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference
	// Experimental.
	ForwardedValuesInput() *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	FunctionAssociation() AwsDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList
	// Experimental.
	FunctionAssociationInput() interface{}
	// Experimental.
	GrpcConfig() AwsDistribution_DefaultCacheBehaviorGrpcConfigPropertyOutputReference
	// Experimental.
	GrpcConfigInput() *AwsDistribution_DefaultCacheBehaviorGrpcConfigProperty
	// Experimental.
	InternalValue() *AwsDistribution_DefaultCacheBehaviorProperty
	// Experimental.
	SetInternalValue(val *AwsDistribution_DefaultCacheBehaviorProperty)
	// Experimental.
	LambdaFunctionAssociation() AwsDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList
	// Experimental.
	LambdaFunctionAssociationInput() interface{}
	// Experimental.
	MaxTtl() *float64
	// Experimental.
	SetMaxTtl(val *float64)
	// Experimental.
	MaxTtlInput() *float64
	// Experimental.
	MinTtl() *float64
	// Experimental.
	SetMinTtl(val *float64)
	// Experimental.
	MinTtlInput() *float64
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
	SmoothStreaming() interface{}
	// Experimental.
	SetSmoothStreaming(val interface{})
	// Experimental.
	SmoothStreamingInput() interface{}
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
	TrustedKeyGroups() *[]*string
	// Experimental.
	SetTrustedKeyGroups(val *[]*string)
	// Experimental.
	TrustedKeyGroupsInput() *[]*string
	// Experimental.
	TrustedSigners() *[]*string
	// Experimental.
	SetTrustedSigners(val *[]*string)
	// Experimental.
	TrustedSignersInput() *[]*string
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
	PutForwardedValues(value *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty)
	// Experimental.
	PutFunctionAssociation(value interface{})
	// Experimental.
	PutGrpcConfig(value *AwsDistribution_DefaultCacheBehaviorGrpcConfigProperty)
	// Experimental.
	PutLambdaFunctionAssociation(value interface{})
	// Experimental.
	ResetCachePolicyId()
	// Experimental.
	ResetCompress()
	// Experimental.
	ResetDefaultTtl()
	// Experimental.
	ResetFieldLevelEncryptionId()
	// Experimental.
	ResetForwardedValues()
	// Experimental.
	ResetFunctionAssociation()
	// Experimental.
	ResetGrpcConfig()
	// Experimental.
	ResetLambdaFunctionAssociation()
	// Experimental.
	ResetMaxTtl()
	// Experimental.
	ResetMinTtl()
	// Experimental.
	ResetOriginRequestPolicyId()
	// Experimental.
	ResetRealtimeLogConfigArn()
	// Experimental.
	ResetResponseHeadersPolicyId()
	// Experimental.
	ResetSmoothStreaming()
	// Experimental.
	ResetTrustedKeyGroups()
	// Experimental.
	ResetTrustedSigners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDistribution_DefaultCacheBehaviorPropertyOutputReference
type jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CachedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ForwardedValues() AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference {
	var returns AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ForwardedValuesInput() *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty
	_jsii_.Get(
		j,
		"forwardedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) FunctionAssociation() AwsDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList {
	var returns AwsDistribution_DefaultCacheBehaviorFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) FunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GrpcConfig() AwsDistribution_DefaultCacheBehaviorGrpcConfigPropertyOutputReference {
	var returns AwsDistribution_DefaultCacheBehaviorGrpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"grpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GrpcConfigInput() *AwsDistribution_DefaultCacheBehaviorGrpcConfigProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorGrpcConfigProperty
	_jsii_.Get(
		j,
		"grpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) InternalValue() *AwsDistribution_DefaultCacheBehaviorProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) LambdaFunctionAssociation() AwsDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList {
	var returns AwsDistribution_DefaultCacheBehaviorLambdaFunctionAssociationPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) LambdaFunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedKeyGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) TrustedSignersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSignersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDistribution_DefaultCacheBehaviorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDistribution_DefaultCacheBehaviorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDistribution_DefaultCacheBehaviorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.DefaultCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDistribution_DefaultCacheBehaviorPropertyOutputReference_Override(a AwsDistribution_DefaultCacheBehaviorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.DefaultCacheBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetCachedMethods(val *[]*string) {
	if err := j.validateSetCachedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedMethods",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetInternalValue(val *AwsDistribution_DefaultCacheBehaviorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetMinTtl(val *float64) {
	if err := j.validateSetMinTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetSmoothStreaming(val interface{}) {
	if err := j.validateSetSmoothStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTrustedKeyGroups(val *[]*string) {
	if err := j.validateSetTrustedKeyGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedKeyGroups",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetTrustedSigners(val *[]*string) {
	if err := j.validateSetTrustedSignersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedSigners",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) PutForwardedValues(value *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty) {
	if err := a.validatePutForwardedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedValues",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) PutFunctionAssociation(value interface{}) {
	if err := a.validatePutFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) PutGrpcConfig(value *AwsDistribution_DefaultCacheBehaviorGrpcConfigProperty) {
	if err := a.validatePutGrpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) PutLambdaFunctionAssociation(value interface{}) {
	if err := a.validatePutLambdaFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		a,
		"resetCompress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetForwardedValues() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetGrpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetMinTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		a,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ResetTrustedSigners() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedSigners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

