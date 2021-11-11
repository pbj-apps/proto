// package: heyrenee.v1
// file: heyrenee/v1/appointment_service.proto

import * as heyrenee_v1_appointment_service_pb from "../../heyrenee/v1/appointment_service_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AppointmentServiceCorsPreflight = {
  readonly methodName: string;
  readonly service: typeof AppointmentService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type AppointmentServiceListAppointments = {
  readonly methodName: string;
  readonly service: typeof AppointmentService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_appointment_service_pb.ListAppointmentsRequest;
  readonly responseType: typeof heyrenee_v1_appointment_service_pb.ListAppointmentsResponse;
};

export class AppointmentService {
  static readonly serviceName: string;
  static readonly CorsPreflight: AppointmentServiceCorsPreflight;
  static readonly ListAppointments: AppointmentServiceListAppointments;
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

export class AppointmentServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  corsPreflight(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  corsPreflight(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  listAppointments(
    requestMessage: heyrenee_v1_appointment_service_pb.ListAppointmentsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_appointment_service_pb.ListAppointmentsResponse|null) => void
  ): UnaryResponse;
  listAppointments(
    requestMessage: heyrenee_v1_appointment_service_pb.ListAppointmentsRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_appointment_service_pb.ListAppointmentsResponse|null) => void
  ): UnaryResponse;
}

