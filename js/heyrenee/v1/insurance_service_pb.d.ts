// package: heyrenee.v1
// file: heyrenee/v1/insurance_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_insurance_pb from "../../heyrenee/v1/messages/insurance_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class CreateInsuranceRequest extends jspb.Message {
  hasInsurance(): boolean;
  clearInsurance(): void;
  getInsurance(): heyrenee_v1_messages_insurance_pb.Insurance | undefined;
  setInsurance(value?: heyrenee_v1_messages_insurance_pb.Insurance): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateInsuranceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateInsuranceRequest): CreateInsuranceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateInsuranceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateInsuranceRequest;
  static deserializeBinaryFromReader(message: CreateInsuranceRequest, reader: jspb.BinaryReader): CreateInsuranceRequest;
}

export namespace CreateInsuranceRequest {
  export type AsObject = {
    insurance?: heyrenee_v1_messages_insurance_pb.Insurance.AsObject,
  }
}

export class UpdateInsuranceRequest extends jspb.Message {
  hasInsurance(): boolean;
  clearInsurance(): void;
  getInsurance(): heyrenee_v1_messages_insurance_pb.Insurance | undefined;
  setInsurance(value?: heyrenee_v1_messages_insurance_pb.Insurance): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateInsuranceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateInsuranceRequest): UpdateInsuranceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateInsuranceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateInsuranceRequest;
  static deserializeBinaryFromReader(message: UpdateInsuranceRequest, reader: jspb.BinaryReader): UpdateInsuranceRequest;
}

export namespace UpdateInsuranceRequest {
  export type AsObject = {
    insurance?: heyrenee_v1_messages_insurance_pb.Insurance.AsObject,
  }
}

export class ListInsuranceRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getInsuranceStatus(): heyrenee_v1_messages_insurance_pb.InsuranceStatusMap[keyof heyrenee_v1_messages_insurance_pb.InsuranceStatusMap];
  setInsuranceStatus(value: heyrenee_v1_messages_insurance_pb.InsuranceStatusMap[keyof heyrenee_v1_messages_insurance_pb.InsuranceStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInsuranceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListInsuranceRequest): ListInsuranceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListInsuranceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInsuranceRequest;
  static deserializeBinaryFromReader(message: ListInsuranceRequest, reader: jspb.BinaryReader): ListInsuranceRequest;
}

export namespace ListInsuranceRequest {
  export type AsObject = {
    patientId: string,
    insuranceStatus: heyrenee_v1_messages_insurance_pb.InsuranceStatusMap[keyof heyrenee_v1_messages_insurance_pb.InsuranceStatusMap],
  }
}

export class ListInsuranceResponse extends jspb.Message {
  clearInsuranceList(): void;
  getInsuranceList(): Array<heyrenee_v1_messages_insurance_pb.Insurance>;
  setInsuranceList(value: Array<heyrenee_v1_messages_insurance_pb.Insurance>): void;
  addInsurance(value?: heyrenee_v1_messages_insurance_pb.Insurance, index?: number): heyrenee_v1_messages_insurance_pb.Insurance;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInsuranceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListInsuranceResponse): ListInsuranceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListInsuranceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInsuranceResponse;
  static deserializeBinaryFromReader(message: ListInsuranceResponse, reader: jspb.BinaryReader): ListInsuranceResponse;
}

export namespace ListInsuranceResponse {
  export type AsObject = {
    insuranceList: Array<heyrenee_v1_messages_insurance_pb.Insurance.AsObject>,
  }
}

