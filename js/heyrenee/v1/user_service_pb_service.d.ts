// package: heyrenee.v1
// file: heyrenee/v1/user_service.proto

import * as heyrenee_v1_user_service_pb from "../../heyrenee/v1/user_service_pb";
import * as heyrenee_v1_messages_patient_pb from "../../heyrenee/v1/messages/patient_pb";
import * as heyrenee_v1_messages_caregiver_pb from "../../heyrenee/v1/messages/caregiver_pb";
import {grpc} from "@improbable-eng/grpc-web";

type UserServiceCreatePatient = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.CreatePatientRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_pb.Patient;
};

type UserServiceGetPatient = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.GetPatientRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_pb.Patient;
};

type UserServiceUpdatePatient = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.UpdatePatientRequest;
  readonly responseType: typeof heyrenee_v1_messages_patient_pb.Patient;
};

type UserServiceCreateCaregiver = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.CreateCaregiverRequest;
  readonly responseType: typeof heyrenee_v1_messages_caregiver_pb.Caregiver;
};

type UserServiceGetCaregiver = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.GetCaregiverRequest;
  readonly responseType: typeof heyrenee_v1_messages_caregiver_pb.Caregiver;
};

type UserServiceUpdateCaregiver = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_user_service_pb.UpdateCaregiverRequest;
  readonly responseType: typeof heyrenee_v1_messages_caregiver_pb.Caregiver;
};

export class UserService {
  static readonly serviceName: string;
  static readonly CreatePatient: UserServiceCreatePatient;
  static readonly GetPatient: UserServiceGetPatient;
  static readonly UpdatePatient: UserServiceUpdatePatient;
  static readonly CreateCaregiver: UserServiceCreateCaregiver;
  static readonly GetCaregiver: UserServiceGetCaregiver;
  static readonly UpdateCaregiver: UserServiceUpdateCaregiver;
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

export class UserServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createPatient(
    requestMessage: heyrenee_v1_user_service_pb.CreatePatientRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  createPatient(
    requestMessage: heyrenee_v1_user_service_pb.CreatePatientRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  getPatient(
    requestMessage: heyrenee_v1_user_service_pb.GetPatientRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  getPatient(
    requestMessage: heyrenee_v1_user_service_pb.GetPatientRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  updatePatient(
    requestMessage: heyrenee_v1_user_service_pb.UpdatePatientRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  updatePatient(
    requestMessage: heyrenee_v1_user_service_pb.UpdatePatientRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_patient_pb.Patient|null) => void
  ): UnaryResponse;
  createCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.CreateCaregiverRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
  createCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.CreateCaregiverRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
  getCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.GetCaregiverRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
  getCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.GetCaregiverRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
  updateCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.UpdateCaregiverRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
  updateCaregiver(
    requestMessage: heyrenee_v1_user_service_pb.UpdateCaregiverRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_caregiver_pb.Caregiver|null) => void
  ): UnaryResponse;
}

