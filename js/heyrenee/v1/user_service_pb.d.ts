// package: heyrenee.v1
// file: heyrenee/v1/user_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_messages_patient_pb from "../../heyrenee/v1/messages/patient_pb";
import * as heyrenee_v1_messages_caregiver_pb from "../../heyrenee/v1/messages/caregiver_pb";

export class CreatePatientRequest extends jspb.Message {
  hasPatient(): boolean;
  clearPatient(): void;
  getPatient(): heyrenee_v1_messages_patient_pb.Patient | undefined;
  setPatient(value?: heyrenee_v1_messages_patient_pb.Patient): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientRequest): CreatePatientRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientRequest;
  static deserializeBinaryFromReader(message: CreatePatientRequest, reader: jspb.BinaryReader): CreatePatientRequest;
}

export namespace CreatePatientRequest {
  export type AsObject = {
    patient?: heyrenee_v1_messages_patient_pb.Patient.AsObject,
    password: string,
  }
}

export class GetPatientRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPatientRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPatientRequest): GetPatientRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPatientRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPatientRequest;
  static deserializeBinaryFromReader(message: GetPatientRequest, reader: jspb.BinaryReader): GetPatientRequest;
}

export namespace GetPatientRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class UpdatePatientRequest extends jspb.Message {
  hasPatient(): boolean;
  clearPatient(): void;
  getPatient(): heyrenee_v1_messages_patient_pb.Patient | undefined;
  setPatient(value?: heyrenee_v1_messages_patient_pb.Patient): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePatientRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePatientRequest): UpdatePatientRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePatientRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePatientRequest;
  static deserializeBinaryFromReader(message: UpdatePatientRequest, reader: jspb.BinaryReader): UpdatePatientRequest;
}

export namespace UpdatePatientRequest {
  export type AsObject = {
    patient?: heyrenee_v1_messages_patient_pb.Patient.AsObject,
  }
}

export class CreateCaregiverRequest extends jspb.Message {
  hasCaregiver(): boolean;
  clearCaregiver(): void;
  getCaregiver(): heyrenee_v1_messages_caregiver_pb.Caregiver | undefined;
  setCaregiver(value?: heyrenee_v1_messages_caregiver_pb.Caregiver): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCaregiverRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCaregiverRequest): CreateCaregiverRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateCaregiverRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCaregiverRequest;
  static deserializeBinaryFromReader(message: CreateCaregiverRequest, reader: jspb.BinaryReader): CreateCaregiverRequest;
}

export namespace CreateCaregiverRequest {
  export type AsObject = {
    caregiver?: heyrenee_v1_messages_caregiver_pb.Caregiver.AsObject,
    password: string,
  }
}

export class GetCaregiverRequest extends jspb.Message {
  getCaregiverId(): string;
  setCaregiverId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCaregiverRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCaregiverRequest): GetCaregiverRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetCaregiverRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCaregiverRequest;
  static deserializeBinaryFromReader(message: GetCaregiverRequest, reader: jspb.BinaryReader): GetCaregiverRequest;
}

export namespace GetCaregiverRequest {
  export type AsObject = {
    caregiverId: string,
  }
}

export class UpdateCaregiverRequest extends jspb.Message {
  hasCaregiver(): boolean;
  clearCaregiver(): void;
  getCaregiver(): heyrenee_v1_messages_caregiver_pb.Caregiver | undefined;
  setCaregiver(value?: heyrenee_v1_messages_caregiver_pb.Caregiver): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCaregiverRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCaregiverRequest): UpdateCaregiverRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateCaregiverRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCaregiverRequest;
  static deserializeBinaryFromReader(message: UpdateCaregiverRequest, reader: jspb.BinaryReader): UpdateCaregiverRequest;
}

export namespace UpdateCaregiverRequest {
  export type AsObject = {
    caregiver?: heyrenee_v1_messages_caregiver_pb.Caregiver.AsObject,
  }
}

