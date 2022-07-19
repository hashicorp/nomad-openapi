# Nomad.Client.Model.Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinities** | [**List&lt;Affinity&gt;**](Affinity.md) |  | [optional] 
**AllAtOnce** | **bool** |  | [optional] 
**Constraints** | [**List&lt;Constraint&gt;**](Constraint.md) |  | [optional] 
**ConsulNamespace** | **string** |  | [optional] 
**ConsulToken** | **string** |  | [optional] 
**CreateIndex** | **int** |  | [optional] 
**Datacenters** | **List&lt;string&gt;** |  | [optional] 
**DispatchIdempotencyToken** | **string** |  | [optional] 
**Dispatched** | **bool** |  | [optional] 
**ID** | **string** |  | [optional] 
**JobModifyIndex** | **int** |  | [optional] 
**Meta** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] 
**ModifyIndex** | **int** |  | [optional] 
**Multiregion** | [**Multiregion**](Multiregion.md) |  | [optional] 
**Name** | **string** |  | [optional] 
**Namespace** | **string** |  | [optional] 
**NomadTokenID** | **string** |  | [optional] 
**ParameterizedJob** | [**ParameterizedJobConfig**](ParameterizedJobConfig.md) |  | [optional] 
**ParentID** | **string** |  | [optional] 
**Payload** | **byte[]** |  | [optional] 
**Periodic** | [**PeriodicConfig**](PeriodicConfig.md) |  | [optional] 
**Priority** | **int** |  | [optional] 
**Region** | **string** |  | [optional] 
**Reschedule** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] 
**Spreads** | [**List&lt;Spread&gt;**](Spread.md) |  | [optional] 
**Stable** | **bool** |  | [optional] 
**Status** | **string** |  | [optional] 
**StatusDescription** | **string** |  | [optional] 
**Stop** | **bool** |  | [optional] 
**SubmitTime** | **long** |  | [optional] 
**TaskGroups** | [**List&lt;TaskGroup&gt;**](TaskGroup.md) |  | [optional] 
**Type** | **string** |  | [optional] 
**Update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] 
**VaultNamespace** | **string** |  | [optional] 
**VaultToken** | **string** |  | [optional] 
**Version** | **int** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

