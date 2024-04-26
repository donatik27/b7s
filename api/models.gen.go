// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"github.com/blocklessnetwork/b7s/models/execute"
	"github.com/blocklessnetwork/b7s/node/aggregate"
)

// AggregatedResult Result of an Execution Request
type AggregatedResult = aggregate.Result

// AggregatedResults List of unique results of the Execution Request
type AggregatedResults = aggregate.Results

// AttributeAttestors Require specific attestors as vouchers
type AttributeAttestors = execute.AttributeAttestors

// ExecutionConfig Configuration options for the Execution Request
type ExecutionConfig = execute.Config

// ExecutionParameter defines model for ExecutionParameter.
type ExecutionParameter = execute.Parameter

// ExecutionRequest defines model for ExecutionRequest.
type ExecutionRequest struct {
	// Config Configuration options for the Execution Request
	Config ExecutionConfig `json:"config,omitempty"`

	// FunctionId CID of the function
	FunctionId string `json:"function_id"`

	// Method Name of the WASM file to execute
	Method string `json:"method"`

	// Parameters CLI arguments for the Blockless Function
	Parameters []ExecutionParameter `json:"parameters,omitempty"`

	// Topic In the scenario where workers form subgroups, you can target a specific subgroup by specifying its identifier
	Topic string `json:"topic,omitempty"`
}

// ExecutionResponse defines model for ExecutionResponse.
type ExecutionResponse struct {
	// Cluster Information about the cluster of nodes that executed this request
	Cluster NodeCluster `json:"cluster,omitempty"`

	// Code Status of the execution
	Code string `json:"code,omitempty"`

	// Message If the Execution Request failed, this message might have more info about the error
	Message string `json:"message,omitempty"`

	// RequestId ID of the Execution Request
	RequestId string `json:"request_id,omitempty"`

	// Results List of unique results of the Execution Request
	Results AggregatedResults `json:"results,omitempty"`
}

// ExecutionResult Actual outputs of the execution, like Standard Output, Standard Error, Exit Code etc..
type ExecutionResult = execute.RuntimeOutput

// FunctionInstallRequest defines model for FunctionInstallRequest.
type FunctionInstallRequest struct {
	// Cid CID of the function
	Cid string `json:"cid"`

	// Topic In a scenario where workers form subgroups, you can target a specific subgroup by specifying its identifier
	Topic string `json:"topic,omitempty"`
	Uri   string `json:"uri,omitempty"`
}

// FunctionInstallResponse defines model for FunctionInstallResponse.
type FunctionInstallResponse struct {
	Code string `json:"code,omitempty"`
}

// FunctionResultRequest Get the result of an Execution Request, identified by the request ID
type FunctionResultRequest struct {
	// Id ID of the Execution Request
	Id string `json:"id"`
}

// FunctionResultResponse defines model for FunctionResultResponse.
type FunctionResultResponse = ExecutionResponse

// HealthStatus Node status
type HealthStatus struct {
	Code string `json:"code,omitempty"`
}

// NamedValue A key-value pair
type NamedValue = execute.EnvVar

// NodeAttributes Attributes that the executing Node should have
type NodeAttributes = execute.Attributes

// NodeCluster Information about the cluster of nodes that executed this request
type NodeCluster = execute.Cluster

// ResultAggregation defines model for ResultAggregation.
type ResultAggregation = execute.ResultAggregation

// RuntimeConfig Configuration options for the Blockless Runtime
type RuntimeConfig = execute.BLSRuntimeConfig

// ExecuteFunctionJSONRequestBody defines body for ExecuteFunction for application/json ContentType.
type ExecuteFunctionJSONRequestBody = ExecutionRequest

// InstallFunctionJSONRequestBody defines body for InstallFunction for application/json ContentType.
type InstallFunctionJSONRequestBody = FunctionInstallRequest

// ExecutionResultJSONRequestBody defines body for ExecutionResult for application/json ContentType.
type ExecutionResultJSONRequestBody = FunctionResultRequest
