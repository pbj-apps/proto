// package: heyrenee.v1
// file: heyrenee/v1/diagnosis_service.proto

import * as heyrenee_v1_diagnosis_service_pb from "../../heyrenee/v1/diagnosis_service_pb";
import * as heyrenee_v1_messages_patient_diagnosis_pb from "../../heyrenee/v1/messages/patient_diagnosis_pb";
import {grpc} from "@improbable-eng/grpc-web";

type DiagnosisServiceDiagnosisSuggest = {
  readonly methodName: string;
  readonly service: typeof DiagnosisService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestRequest;
  readonly responseType: typeof heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestResponse;
};

type DiagnosisServiceCreatePatientDiagnosis = {
  readonly methodName: string;
  readonly service: typeof DiagnosisService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_diagnosis_service_pb.CreatePatientDiagnosisRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis;
};

type DiagnosisServiceUpdatePatientDiagnosis = {
  readonly methodName: string;
  readonly service: typeof DiagnosisService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_diagnosis_service_pb.UpdatePatientDiagnosisRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis;
};

type DiagnosisServiceListPatientDiagnoses = {
  readonly methodName: string;
  readonly service: typeof DiagnosisService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesRequest;
  readonly responseType: typeof heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesResponse;
};

export class DiagnosisService {
  static readonly serviceName: string;
  static readonly DiagnosisSuggest: DiagnosisServiceDiagnosisSuggest;
  static readonly CreatePatientDiagnosis: DiagnosisServiceCreatePatientDiagnosis;
  static readonly UpdatePatientDiagnosis: DiagnosisServiceUpdatePatientDiagnosis;
  static readonly ListPatientDiagnoses: DiagnosisServiceListPatientDiagnoses;
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

export class DiagnosisServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  diagnosisSuggest(
    requestMessage: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestResponse|null) => void
  ): UnaryResponse;
  diagnosisSuggest(
    requestMessage: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestResponse|null) => void
  ): UnaryResponse;
  createPatientDiagnosis(
    requestMessage: heyrenee_v1_diagnosis_service_pb.CreatePatientDiagnosisRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis|null) => void
  ): UnaryResponse;
  createPatientDiagnosis(
    requestMessage: heyrenee_v1_diagnosis_service_pb.CreatePatientDiagnosisRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis|null) => void
  ): UnaryResponse;
  updatePatientDiagnosis(
    requestMessage: heyrenee_v1_diagnosis_service_pb.UpdatePatientDiagnosisRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis|null) => void
  ): UnaryResponse;
  updatePatientDiagnosis(
    requestMessage: heyrenee_v1_diagnosis_service_pb.UpdatePatientDiagnosisRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis|null) => void
  ): UnaryResponse;
  listPatientDiagnoses(
    requestMessage: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesResponse|null) => void
  ): UnaryResponse;
  listPatientDiagnoses(
    requestMessage: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesResponse|null) => void
  ): UnaryResponse;
}

