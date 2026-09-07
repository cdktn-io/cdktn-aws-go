//go:build no_runtime_type_checking

package networkfirewall

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsFirewall_AttachmentPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsFirewall_AttachmentPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

