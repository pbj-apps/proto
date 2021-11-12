// package: heyrenee.v1
// file: heyrenee/v1/provider_service.proto

import * as heyrenee_v1_provider_service_pb from "../../heyrenee/v1/provider_service_pb";
import * as heyrenee_v1_messages_patient_provider_pb from "../../heyrenee/v1/messages/patient_provider_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ProviderServiceProviderSuggest = {
  readonly methodName: string;
  readonly service: typeof ProviderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_provider_service_pb.ProviderSuggestRequest;
  readonly responseType: typeof heyrenee_v1_provider_service_pb.ProviderSuggestResponse;
};

type ProviderServiceCreatePatientProvider = {
  readonly methodName: string;
  readonly service: typeof ProviderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_provider_service_pb.CreatePatientProviderRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_provider_pb.PatientProvider;
};

type ProviderServiceUpdatePatientProvider = {
  readonly methodName: string;
  readonly service: typeof ProviderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_provider_service_pb.UpdatePatientProviderRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_provider_pb.PatientProvider;
};

type ProviderServiceListPatientProviders = {
  readonly methodName: string;
  readonly service: typeof ProviderService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_provider_service_pb.ListPatientProvidersRequest;
  readonly responseType: typeof heyrenee_v1_provider_service_pb.ListPatientProvidersResponse;
};

export class ProviderService {
  static readonly serviceName: string;
  static readonly ProviderSuggest: ProviderServiceProviderSuggest;
  static readonly CreatePatientProvider: ProviderServiceCreatePatientProvider;
  static readonly UpdatePatientProvider: ProviderServiceUpdatePatientProvider;
  static readonly ListPatientProviders: ProviderServiceListPatientProviders;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ProviderServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  providerSuggest(
    requestMessage: heyrenee_v1_provider_service_pb.ProviderSuggestRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_provider_service_pb.ProviderSuggestResponse|null) => void
  ): UnaryResponse;
  providerSuggest(
    requestMessage: heyrenee_v1_provider_service_pb.ProviderSuggestRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_provider_service_pb.ProviderSuggestResponse|null) => void
  ): UnaryResponse;
  createPatientProvider(
    requestMessage: heyrenee_v1_provider_service_pb.CreatePatientProviderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_provider_pb.PatientProvider|null) => void
  ): UnaryResponse;
  createPatientProvider(
    requestMessage: heyrenee_v1_provider_service_pb.CreatePatientProviderRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_provider_pb.PatientProvider|null) => void
  ): UnaryResponse;
  updatePatientProvider(
    requestMessage: heyrenee_v1_provider_service_pb.UpdatePatientProviderRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_provider_pb.PatientProvider|null) => void
  ): UnaryResponse;
  updatePatientProvider(
    requestMessage: heyrenee_v1_provider_service_pb.UpdatePatientProviderRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_provider_pb.PatientProvider|null) => void
  ): UnaryResponse;
  listPatientProviders(
    requestMessage: heyrenee_v1_provider_service_pb.ListPatientProvidersRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_provider_service_pb.ListPatientProvidersResponse|null) => void
  ): UnaryResponse;
  listPatientProviders(
    requestMessage: heyrenee_v1_provider_service_pb.ListPatientProvidersRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_provider_service_pb.ListPatientProvidersResponse|null) => void
  ): UnaryResponse;
}

