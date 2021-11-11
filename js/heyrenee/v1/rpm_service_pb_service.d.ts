// package: heyrenee.v1
// file: heyrenee/v1/rpm_service.proto

import * as heyrenee_v1_rpm_service_pb from "../../heyrenee/v1/rpm_service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type RpmServiceListRpmSchedules = {
  readonly methodName: string;
  readonly service: typeof RpmService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_rpm_service_pb.ListRpmSchedulesRequest;
  readonly responseType: typeof heyrenee_v1_rpm_service_pb.ListRpmSchedulesResponse;
};

type RpmServiceListRpmMeasurements = {
  readonly methodName: string;
  readonly service: typeof RpmService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_rpm_service_pb.ListRpmMeasurementsRequest;
  readonly responseType: typeof heyrenee_v1_rpm_service_pb.ListRpmMeasurementsResponse;
};

export class RpmService {
  static readonly serviceName: string;
  static readonly ListRpmSchedules: RpmServiceListRpmSchedules;
  static readonly ListRpmMeasurements: RpmServiceListRpmMeasurements;
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

export class RpmServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listRpmSchedules(
    requestMessage: heyrenee_v1_rpm_service_pb.ListRpmSchedulesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_rpm_service_pb.ListRpmSchedulesResponse|null) => void
  ): UnaryResponse;
  listRpmSchedules(
    requestMessage: heyrenee_v1_rpm_service_pb.ListRpmSchedulesRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_rpm_service_pb.ListRpmSchedulesResponse|null) => void
  ): UnaryResponse;
  listRpmMeasurements(
    requestMessage: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsResponse|null) => void
  ): UnaryResponse;
  listRpmMeasurements(
    requestMessage: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsResponse|null) => void
  ): UnaryResponse;
}

