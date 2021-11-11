// package: heyrenee.v1
// file: heyrenee/v1/insurance_service.proto

import * as heyrenee_v1_insurance_service_pb from "../../heyrenee/v1/insurance_service_pb";
import * as heyrenee_v1_messages_insurance_pb from "../../heyrenee/v1/messages/insurance_pb";
import {grpc} from "@improbable-eng/grpc-web";

type InsuranceServiceCreateInsurance = {
  readonly methodName: string;
  readonly service: typeof InsuranceService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_insurance_service_pb.CreateInsuranceRequest;
  readonly responseType: typeof heyrenee_v1_messages_insurance_pb.Insurance;
};

type InsuranceServiceUpdateInsurance = {
  readonly methodName: string;
  readonly service: typeof InsuranceService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_insurance_service_pb.UpdateInsuranceRequest;
  readonly responseType: typeof heyrenee_v1_messages_insurance_pb.Insurance;
};

type InsuranceServiceListInsurance = {
  readonly methodName: string;
  readonly service: typeof InsuranceService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof heyrenee_v1_insurance_service_pb.ListInsuranceRequest;
  readonly responseType: typeof heyrenee_v1_insurance_service_pb.ListInsuranceResponse;
};

export class InsuranceService {
  static readonly serviceName: string;
  static readonly CreateInsurance: InsuranceServiceCreateInsurance;
  static readonly UpdateInsurance: InsuranceServiceUpdateInsurance;
  static readonly ListInsurance: InsuranceServiceListInsurance;
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

export class InsuranceServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.CreateInsuranceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_insurance_pb.Insurance|null) => void
  ): UnaryResponse;
  createInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.CreateInsuranceRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_insurance_pb.Insurance|null) => void
  ): UnaryResponse;
  updateInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.UpdateInsuranceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_insurance_pb.Insurance|null) => void
  ): UnaryResponse;
  updateInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.UpdateInsuranceRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_messages_insurance_pb.Insurance|null) => void
  ): UnaryResponse;
  listInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.ListInsuranceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_insurance_service_pb.ListInsuranceResponse|null) => void
  ): UnaryResponse;
  listInsurance(
    requestMessage: heyrenee_v1_insurance_service_pb.ListInsuranceRequest,
    callback: (error: ServiceError|null, responseMessage: heyrenee_v1_insurance_service_pb.ListInsuranceResponse|null) => void
  ): UnaryResponse;
}

