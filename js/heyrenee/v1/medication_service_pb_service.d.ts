// package: heyrenee.v1
// file: heyrenee/v1/medication_service.proto

import * as heyrenee_v1_medication_service_pb from "../../heyrenee/v1/medication_service_pb";
import * as heyrenee_v1_messages_prescription_pb from "../../heyrenee/v1/messages/prescription_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MedicationServiceMedicationSuggest = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.MedicationSuggestRequest;
  readonly responseType: typeof heyrenee_v1_medication_service_pb.MedicationSuggestResponse;
};

type MedicationServiceCreatePrescription = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.CreatePrescriptionRequest;
  readonly responseType: typeof heyrenee_v1_messages_prescription_pb.Prescription;
};

type MedicationServiceUpdatePrescription = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.UpdatePrescriptionRequest;
  readonly responseType: typeof heyrenee_v1_messages_prescription_pb.Prescription;
};

type MedicationServiceListPrescriptions = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.ListPrescriptionsRequest;
  readonly responseType: typeof heyrenee_v1_medication_service_pb.ListPrescriptionsResponse;
};

type MedicationServiceListMedicationDoses = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.ListMedicationDosesRequest;
  readonly responseType: typeof heyrenee_v1_medication_service_pb.ListMedicationDosesResponse;
};

type MedicationServiceListRefills = {
  readonly methodName: string;
  readonly service: typeof MedicationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_medication_service_pb.ListRefillsRequest;
  readonly responseType: typeof heyrenee_v1_medication_service_pb.ListRefillsResponse;
};

export class MedicationService {
  static readonly serviceName: string;
  static readonly MedicationSuggest: MedicationServiceMedicationSuggest;
  static readonly CreatePrescription: MedicationServiceCreatePrescription;
  static readonly UpdatePrescription: MedicationServiceUpdatePrescription;
  static readonly ListPrescriptions: MedicationServiceListPrescriptions;
  static readonly ListMedicationDoses: MedicationServiceListMedicationDoses;
  static readonly ListRefills: MedicationServiceListRefills;
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

export class MedicationServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  medicationSuggest(
    requestMessage: heyrenee_v1_medication_service_pb.MedicationSuggestRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.MedicationSuggestResponse|null) => void
  ): UnaryResponse;
  medicationSuggest(
    requestMessage: heyrenee_v1_medication_service_pb.MedicationSuggestRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.MedicationSuggestResponse|null) => void
  ): UnaryResponse;
  createPrescription(
    requestMessage: heyrenee_v1_medication_service_pb.CreatePrescriptionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_prescription_pb.Prescription|null) => void
  ): UnaryResponse;
  createPrescription(
    requestMessage: heyrenee_v1_medication_service_pb.CreatePrescriptionRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_prescription_pb.Prescription|null) => void
  ): UnaryResponse;
  updatePrescription(
    requestMessage: heyrenee_v1_medication_service_pb.UpdatePrescriptionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_prescription_pb.Prescription|null) => void
  ): UnaryResponse;
  updatePrescription(
    requestMessage: heyrenee_v1_medication_service_pb.UpdatePrescriptionRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_prescription_pb.Prescription|null) => void
  ): UnaryResponse;
  listPrescriptions(
    requestMessage: heyrenee_v1_medication_service_pb.ListPrescriptionsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListPrescriptionsResponse|null) => void
  ): UnaryResponse;
  listPrescriptions(
    requestMessage: heyrenee_v1_medication_service_pb.ListPrescriptionsRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListPrescriptionsResponse|null) => void
  ): UnaryResponse;
  listMedicationDoses(
    requestMessage: heyrenee_v1_medication_service_pb.ListMedicationDosesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListMedicationDosesResponse|null) => void
  ): UnaryResponse;
  listMedicationDoses(
    requestMessage: heyrenee_v1_medication_service_pb.ListMedicationDosesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListMedicationDosesResponse|null) => void
  ): UnaryResponse;
  listRefills(
    requestMessage: heyrenee_v1_medication_service_pb.ListRefillsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListRefillsResponse|null) => void
  ): UnaryResponse;
  listRefills(
    requestMessage: heyrenee_v1_medication_service_pb.ListRefillsRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_medication_service_pb.ListRefillsResponse|null) => void
  ): UnaryResponse;
}

