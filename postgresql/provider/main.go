// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-postgresql.provider.PostgresqlProvider",
		reflect.TypeOf((*PostgresqlProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamAuth", GoGetter: "AwsRdsIamAuth"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamAuthInput", GoGetter: "AwsRdsIamAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamProfile", GoGetter: "AwsRdsIamProfile"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamProfileInput", GoGetter: "AwsRdsIamProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamRegion", GoGetter: "AwsRdsIamRegion"},
			_jsii_.MemberProperty{JsiiProperty: "awsRdsIamRegionInput", GoGetter: "AwsRdsIamRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureIdentityAuth", GoGetter: "AzureIdentityAuth"},
			_jsii_.MemberProperty{JsiiProperty: "azureIdentityAuthInput", GoGetter: "AzureIdentityAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantId", GoGetter: "AzureTenantId"},
			_jsii_.MemberProperty{JsiiProperty: "azureTenantIdInput", GoGetter: "AzureTenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientcert", GoGetter: "Clientcert"},
			_jsii_.MemberProperty{JsiiProperty: "clientcertInput", GoGetter: "ClientcertInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeout", GoGetter: "ConnectTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeoutInput", GoGetter: "ConnectTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseUsername", GoGetter: "DatabaseUsername"},
			_jsii_.MemberProperty{JsiiProperty: "databaseUsernameInput", GoGetter: "DatabaseUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "expectedVersion", GoGetter: "ExpectedVersion"},
			_jsii_.MemberProperty{JsiiProperty: "expectedVersionInput", GoGetter: "ExpectedVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnections", GoGetter: "MaxConnections"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsInput", GoGetter: "MaxConnectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsRdsIamAuth", GoMethod: "ResetAwsRdsIamAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsRdsIamProfile", GoMethod: "ResetAwsRdsIamProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsRdsIamRegion", GoMethod: "ResetAwsRdsIamRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureIdentityAuth", GoMethod: "ResetAzureIdentityAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureTenantId", GoMethod: "ResetAzureTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientcert", GoMethod: "ResetClientcert"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectTimeout", GoMethod: "ResetConnectTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabase", GoMethod: "ResetDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabaseUsername", GoMethod: "ResetDatabaseUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpectedVersion", GoMethod: "ResetExpectedVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConnections", GoMethod: "ResetMaxConnections"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheme", GoMethod: "ResetScheme"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslmode", GoMethod: "ResetSslmode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslMode", GoMethod: "ResetSslMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslrootcert", GoMethod: "ResetSslrootcert"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuperuser", GoMethod: "ResetSuperuser"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslmode", GoGetter: "Sslmode"},
			_jsii_.MemberProperty{JsiiProperty: "sslMode", GoGetter: "SslMode"},
			_jsii_.MemberProperty{JsiiProperty: "sslmodeInput", GoGetter: "SslmodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslModeInput", GoGetter: "SslModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslrootcert", GoGetter: "Sslrootcert"},
			_jsii_.MemberProperty{JsiiProperty: "sslrootcertInput", GoGetter: "SslrootcertInput"},
			_jsii_.MemberProperty{JsiiProperty: "superuser", GoGetter: "Superuser"},
			_jsii_.MemberProperty{JsiiProperty: "superuserInput", GoGetter: "SuperuserInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_PostgresqlProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-postgresql.provider.PostgresqlProviderClientcert",
		reflect.TypeOf((*PostgresqlProviderClientcert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-postgresql.provider.PostgresqlProviderConfig",
		reflect.TypeOf((*PostgresqlProviderConfig)(nil)).Elem(),
	)
}
