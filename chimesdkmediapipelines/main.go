// cdk-terrain bindings for the Chime SDK Media Pipelines group of terraform-provider-aws 6.62.0
package chimesdkmediapipelines

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "elements", GoGetter: "Elements"},
			_jsii_.MemberProperty{JsiiProperty: "elementsInput", GoGetter: "ElementsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putElements", GoMethod: "PutElements"},
			_jsii_.MemberMethod{JsiiMethod: "putRealTimeAlertConfiguration", GoMethod: "PutRealTimeAlertConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "realTimeAlertConfiguration", GoGetter: "RealTimeAlertConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "realTimeAlertConfigurationInput", GoGetter: "RealTimeAlertConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealTimeAlertConfiguration", GoMethod: "ResetRealTimeAlertConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAccessRoleArn", GoGetter: "ResourceAccessRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAccessRoleArnInput", GoGetter: "ResourceAccessRoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "callAnalyticsStreamCategories", GoGetter: "CallAnalyticsStreamCategories"},
			_jsii_.MemberProperty{JsiiProperty: "callAnalyticsStreamCategoriesInput", GoGetter: "CallAnalyticsStreamCategoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentIdentificationType", GoGetter: "ContentIdentificationType"},
			_jsii_.MemberProperty{JsiiProperty: "contentIdentificationTypeInput", GoGetter: "ContentIdentificationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionType", GoGetter: "ContentRedactionType"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionTypeInput", GoGetter: "ContentRedactionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartialResultsStabilization", GoGetter: "EnablePartialResultsStabilization"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartialResultsStabilizationInput", GoGetter: "EnablePartialResultsStabilizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "filterPartialResults", GoGetter: "FilterPartialResults"},
			_jsii_.MemberProperty{JsiiProperty: "filterPartialResultsInput", GoGetter: "FilterPartialResultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "languageModelName", GoGetter: "LanguageModelName"},
			_jsii_.MemberProperty{JsiiProperty: "languageModelNameInput", GoGetter: "LanguageModelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "partialResultsStability", GoGetter: "PartialResultsStability"},
			_jsii_.MemberProperty{JsiiProperty: "partialResultsStabilityInput", GoGetter: "PartialResultsStabilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "piiEntityTypes", GoGetter: "PiiEntityTypes"},
			_jsii_.MemberProperty{JsiiProperty: "piiEntityTypesInput", GoGetter: "PiiEntityTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "postCallAnalyticsSettings", GoGetter: "PostCallAnalyticsSettings"},
			_jsii_.MemberProperty{JsiiProperty: "postCallAnalyticsSettingsInput", GoGetter: "PostCallAnalyticsSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostCallAnalyticsSettings", GoMethod: "PutPostCallAnalyticsSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetCallAnalyticsStreamCategories", GoMethod: "ResetCallAnalyticsStreamCategories"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentIdentificationType", GoMethod: "ResetContentIdentificationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentRedactionType", GoMethod: "ResetContentRedactionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePartialResultsStabilization", GoMethod: "ResetEnablePartialResultsStabilization"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilterPartialResults", GoMethod: "ResetFilterPartialResults"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageModelName", GoMethod: "ResetLanguageModelName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPartialResultsStability", GoMethod: "ResetPartialResultsStability"},
			_jsii_.MemberMethod{JsiiMethod: "resetPiiEntityTypes", GoMethod: "ResetPiiEntityTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostCallAnalyticsSettings", GoMethod: "ResetPostCallAnalyticsSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyFilterMethod", GoMethod: "ResetVocabularyFilterMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyFilterName", GoMethod: "ResetVocabularyFilterName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyName", GoMethod: "ResetVocabularyName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterMethod", GoGetter: "VocabularyFilterMethod"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterMethodInput", GoGetter: "VocabularyFilterMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterName", GoGetter: "VocabularyFilterName"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterNameInput", GoGetter: "VocabularyFilterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyName", GoGetter: "VocabularyName"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyNameInput", GoGetter: "VocabularyNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeProcessorConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.AmazonTranscribeProcessorConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentIdentificationType", GoGetter: "ContentIdentificationType"},
			_jsii_.MemberProperty{JsiiProperty: "contentIdentificationTypeInput", GoGetter: "ContentIdentificationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionType", GoGetter: "ContentRedactionType"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionTypeInput", GoGetter: "ContentRedactionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartialResultsStabilization", GoGetter: "EnablePartialResultsStabilization"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartialResultsStabilizationInput", GoGetter: "EnablePartialResultsStabilizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "filterPartialResults", GoGetter: "FilterPartialResults"},
			_jsii_.MemberProperty{JsiiProperty: "filterPartialResultsInput", GoGetter: "FilterPartialResultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "languageModelName", GoGetter: "LanguageModelName"},
			_jsii_.MemberProperty{JsiiProperty: "languageModelNameInput", GoGetter: "LanguageModelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "partialResultsStability", GoGetter: "PartialResultsStability"},
			_jsii_.MemberProperty{JsiiProperty: "partialResultsStabilityInput", GoGetter: "PartialResultsStabilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "piiEntityTypes", GoGetter: "PiiEntityTypes"},
			_jsii_.MemberProperty{JsiiProperty: "piiEntityTypesInput", GoGetter: "PiiEntityTypesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentIdentificationType", GoMethod: "ResetContentIdentificationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentRedactionType", GoMethod: "ResetContentRedactionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePartialResultsStabilization", GoMethod: "ResetEnablePartialResultsStabilization"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilterPartialResults", GoMethod: "ResetFilterPartialResults"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageModelName", GoMethod: "ResetLanguageModelName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPartialResultsStability", GoMethod: "ResetPartialResultsStability"},
			_jsii_.MemberMethod{JsiiMethod: "resetPiiEntityTypes", GoMethod: "ResetPiiEntityTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetShowSpeakerLabel", GoMethod: "ResetShowSpeakerLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyFilterMethod", GoMethod: "ResetVocabularyFilterMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyFilterName", GoMethod: "ResetVocabularyFilterName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVocabularyName", GoMethod: "ResetVocabularyName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "showSpeakerLabel", GoGetter: "ShowSpeakerLabel"},
			_jsii_.MemberProperty{JsiiProperty: "showSpeakerLabelInput", GoGetter: "ShowSpeakerLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterMethod", GoGetter: "VocabularyFilterMethod"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterMethodInput", GoGetter: "VocabularyFilterMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterName", GoGetter: "VocabularyFilterName"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyFilterNameInput", GoGetter: "VocabularyFilterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyName", GoGetter: "VocabularyName"},
			_jsii_.MemberProperty{JsiiProperty: "vocabularyNameInput", GoGetter: "VocabularyNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.ElementsProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_ElementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.ElementsPropertyList",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_ElementsPropertyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.ElementsPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amazonTranscribeCallAnalyticsProcessorConfiguration", GoGetter: "AmazonTranscribeCallAnalyticsProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonTranscribeCallAnalyticsProcessorConfigurationInput", GoGetter: "AmazonTranscribeCallAnalyticsProcessorConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "amazonTranscribeProcessorConfiguration", GoGetter: "AmazonTranscribeProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonTranscribeProcessorConfigurationInput", GoGetter: "AmazonTranscribeProcessorConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisDataStreamSinkConfiguration", GoGetter: "KinesisDataStreamSinkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisDataStreamSinkConfigurationInput", GoGetter: "KinesisDataStreamSinkConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionSinkConfiguration", GoGetter: "LambdaFunctionSinkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionSinkConfigurationInput", GoGetter: "LambdaFunctionSinkConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonTranscribeCallAnalyticsProcessorConfiguration", GoMethod: "PutAmazonTranscribeCallAnalyticsProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonTranscribeProcessorConfiguration", GoMethod: "PutAmazonTranscribeProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putKinesisDataStreamSinkConfiguration", GoMethod: "PutKinesisDataStreamSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaFunctionSinkConfiguration", GoMethod: "PutLambdaFunctionSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putS3RecordingSinkConfiguration", GoMethod: "PutS3RecordingSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSnsTopicSinkConfiguration", GoMethod: "PutSnsTopicSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSqsQueueSinkConfiguration", GoMethod: "PutSqsQueueSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putVoiceAnalyticsProcessorConfiguration", GoMethod: "PutVoiceAnalyticsProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonTranscribeCallAnalyticsProcessorConfiguration", GoMethod: "ResetAmazonTranscribeCallAnalyticsProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonTranscribeProcessorConfiguration", GoMethod: "ResetAmazonTranscribeProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetKinesisDataStreamSinkConfiguration", GoMethod: "ResetKinesisDataStreamSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionSinkConfiguration", GoMethod: "ResetLambdaFunctionSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3RecordingSinkConfiguration", GoMethod: "ResetS3RecordingSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnsTopicSinkConfiguration", GoMethod: "ResetSnsTopicSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqsQueueSinkConfiguration", GoMethod: "ResetSqsQueueSinkConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetVoiceAnalyticsProcessorConfiguration", GoMethod: "ResetVoiceAnalyticsProcessorConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3RecordingSinkConfiguration", GoGetter: "S3RecordingSinkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "s3RecordingSinkConfigurationInput", GoGetter: "S3RecordingSinkConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicSinkConfiguration", GoGetter: "SnsTopicSinkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicSinkConfigurationInput", GoGetter: "SnsTopicSinkConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "sqsQueueSinkConfiguration", GoGetter: "SqsQueueSinkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "sqsQueueSinkConfigurationInput", GoGetter: "SqsQueueSinkConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "voiceAnalyticsProcessorConfiguration", GoGetter: "VoiceAnalyticsProcessorConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "voiceAnalyticsProcessorConfigurationInput", GoGetter: "VoiceAnalyticsProcessorConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.IssueDetectionConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.IssueDetectionConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ruleName", GoGetter: "RuleName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNameInput", GoGetter: "RuleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.KeywordMatchConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.KeywordMatchConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keywords", GoGetter: "Keywords"},
			_jsii_.MemberProperty{JsiiProperty: "keywordsInput", GoGetter: "KeywordsInput"},
			_jsii_.MemberProperty{JsiiProperty: "negate", GoGetter: "Negate"},
			_jsii_.MemberProperty{JsiiProperty: "negateInput", GoGetter: "NegateInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNegate", GoMethod: "ResetNegate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ruleName", GoGetter: "RuleName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNameInput", GoGetter: "RuleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.KinesisDataStreamSinkConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.KinesisDataStreamSinkConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTarget", GoGetter: "InsightsTarget"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTargetInput", GoGetter: "InsightsTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.LambdaFunctionSinkConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.LambdaFunctionSinkConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTarget", GoGetter: "InsightsTarget"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTargetInput", GoGetter: "InsightsTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.PostCallAnalyticsSettingsProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.PostCallAnalyticsSettingsPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionOutput", GoGetter: "ContentRedactionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "contentRedactionOutputInput", GoGetter: "ContentRedactionOutputInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccessRoleArn", GoGetter: "DataAccessRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccessRoleArnInput", GoGetter: "DataAccessRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "outputEncryptionKmsKeyId", GoGetter: "OutputEncryptionKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "outputEncryptionKmsKeyIdInput", GoGetter: "OutputEncryptionKmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputLocation", GoGetter: "OutputLocation"},
			_jsii_.MemberProperty{JsiiProperty: "outputLocationInput", GoGetter: "OutputLocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentRedactionOutput", GoMethod: "ResetContentRedactionOutput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputEncryptionKmsKeyId", GoMethod: "ResetOutputEncryptionKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RealTimeAlertConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_RealTimeAlertConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RealTimeAlertConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_RealTimeAlertConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disabled", GoGetter: "Disabled"},
			_jsii_.MemberProperty{JsiiProperty: "disabledInput", GoGetter: "DisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putRules", GoMethod: "PutRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabled", GoMethod: "ResetDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "rulesInput", GoGetter: "RulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_RealTimeAlertConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RulesProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_RulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RulesPropertyList",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_RulesPropertyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.RulesPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issueDetectionConfiguration", GoGetter: "IssueDetectionConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "issueDetectionConfigurationInput", GoGetter: "IssueDetectionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "keywordMatchConfiguration", GoGetter: "KeywordMatchConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "keywordMatchConfigurationInput", GoGetter: "KeywordMatchConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIssueDetectionConfiguration", GoMethod: "PutIssueDetectionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putKeywordMatchConfiguration", GoMethod: "PutKeywordMatchConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSentimentConfiguration", GoMethod: "PutSentimentConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssueDetectionConfiguration", GoMethod: "ResetIssueDetectionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeywordMatchConfiguration", GoMethod: "ResetKeywordMatchConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSentimentConfiguration", GoMethod: "ResetSentimentConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sentimentConfiguration", GoGetter: "SentimentConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "sentimentConfigurationInput", GoGetter: "SentimentConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_RulesPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.S3RecordingSinkConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.S3RecordingSinkConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SentimentConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SentimentConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ruleName", GoGetter: "RuleName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNameInput", GoGetter: "RuleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sentimentType", GoGetter: "SentimentType"},
			_jsii_.MemberProperty{JsiiProperty: "sentimentTypeInput", GoGetter: "SentimentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriod", GoGetter: "TimePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodInput", GoGetter: "TimePeriodInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_SentimentConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SnsTopicSinkConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SnsTopicSinkConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTarget", GoGetter: "InsightsTarget"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTargetInput", GoGetter: "InsightsTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SqsQueueSinkConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.SqsQueueSinkConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTarget", GoGetter: "InsightsTarget"},
			_jsii_.MemberProperty{JsiiProperty: "insightsTargetInput", GoGetter: "InsightsTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.TimeoutsProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_TimeoutsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.TimeoutsPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_TimeoutsPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_TimeoutsPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.VoiceAnalyticsProcessorConfigurationProperty",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.VoiceAnalyticsProcessorConfigurationPropertyOutputReference",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "speakerSearchStatus", GoGetter: "SpeakerSearchStatus"},
			_jsii_.MemberProperty{JsiiProperty: "speakerSearchStatusInput", GoGetter: "SpeakerSearchStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "voiceToneAnalysisStatus", GoGetter: "VoiceToneAnalysisStatus"},
			_jsii_.MemberProperty{JsiiProperty: "voiceToneAnalysisStatusInput", GoGetter: "VoiceToneAnalysisStatusInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfigurationConfig",
		reflect.TypeOf((*AwsMediaInsightsPipelineConfigurationConfig)(nil)).Elem(),
	)
}
