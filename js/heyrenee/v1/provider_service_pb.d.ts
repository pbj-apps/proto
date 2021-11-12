// package: heyrenee.v1
// file: heyrenee/v1/provider_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_patient_provider_pb from "../../heyrenee/v1/messages/patient_provider_pb";
import * as heyrenee_v1_messages_provider_pb from "../../heyrenee/v1/messages/provider_pb";
import * as heyrenee_v1_enums_specialty_pb from "../../heyrenee/v1/enums/specialty_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class SearchProvidersRequest extends jspb.Message {
  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getSpecialty(): heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap];
  setSpecialty(value: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap]): void;

  getCity(): string;
  setCity(value: string): void;

  getState(): string;
  setState(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchProvidersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SearchProvidersRequest): SearchProvidersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SearchProvidersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchProvidersRequest;
  static deserializeBinaryFromReader(message: SearchProvidersRequest, reader: jspb.BinaryReader): SearchProvidersRequest;
}

export namespace SearchProvidersRequest {
  export type AsObject = {
    firstName: string,
    lastName: string,
    specialty: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap],
    city: string,
    state: string,
  }
}

export class SearchProvidersResponse extends jspb.Message {
  clearProvidersList(): void;
  getProvidersList(): Array<heyrenee_v1_messages_provider_pb.Provider>;
  setProvidersList(value: Array<heyrenee_v1_messages_provider_pb.Provider>): void;
  addProviders(value?: heyrenee_v1_messages_provider_pb.Provider, index?: number): heyrenee_v1_messages_provider_pb.Provider;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchProvidersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SearchProvidersResponse): SearchProvidersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SearchProvidersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchProvidersResponse;
  static deserializeBinaryFromReader(message: SearchProvidersResponse, reader: jspb.BinaryReader): SearchProvidersResponse;
}

export namespace SearchProvidersResponse {
  export type AsObject = {
    providersList: Array<heyrenee_v1_messages_provider_pb.Provider.AsObject>,
  }
}

export class CreatePatientProviderRequest extends jspb.Message {
  hasPatientProvider(): boolean;
  clearPatientProvider(): void;
  getPatientProvider(): heyrenee_v1_messages_patient_provider_pb.PatientProvider | undefined;
  setPatientProvider(value?: heyrenee_v1_messages_patient_provider_pb.PatientProvider): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientProviderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientProviderRequest): CreatePatientProviderRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientProviderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientProviderRequest;
  static deserializeBinaryFromReader(message: CreatePatientProviderRequest, reader: jspb.BinaryReader): CreatePatientProviderRequest;
}

export namespace CreatePatientProviderRequest {
  export type AsObject = {
    patientProvider?: heyrenee_v1_messages_patient_provider_pb.PatientProvider.AsObject,
  }
}

export class UpdatePatientProviderRequest extends jspb.Message {
  hasPatientProvider(): boolean;
  clearPatientProvider(): void;
  getPatientProvider(): heyrenee_v1_messages_patient_provider_pb.PatientProvider | undefined;
  setPatientProvider(value?: heyrenee_v1_messages_patient_provider_pb.PatientProvider): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePatientProviderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePatientProviderRequest): UpdatePatientProviderRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePatientProviderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePatientProviderRequest;
  static deserializeBinaryFromReader(message: UpdatePatientProviderRequest, reader: jspb.BinaryReader): UpdatePatientProviderRequest;
}

export namespace UpdatePatientProviderRequest {
  export type AsObject = {
    patientProvider?: heyrenee_v1_messages_patient_provider_pb.PatientProvider.AsObject,
  }
}

export class ListPatientProvidersRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientProviderStatus(): heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap];
  setPatientProviderStatus(value: heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap]): void;

  getPatientProviderType(): heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap];
  setPatientProviderType(value: heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientProvidersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientProvidersRequest): ListPatientProvidersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientProvidersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientProvidersRequest;
  static deserializeBinaryFromReader(message: ListPatientProvidersRequest, reader: jspb.BinaryReader): ListPatientProvidersRequest;
}

export namespace ListPatientProvidersRequest {
  export type AsObject = {
    patientId: string,
    patientProviderStatus: heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderStatusMap],
    patientProviderType: heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap[keyof heyrenee_v1_messages_patient_provider_pb.PatientProviderTypeMap],
  }
}

export class ListPatientProvidersResponse extends jspb.Message {
  clearPatientProvidersList(): void;
  getPatientProvidersList(): Array<heyrenee_v1_messages_patient_provider_pb.PatientProvider>;
  setPatientProvidersList(value: Array<heyrenee_v1_messages_patient_provider_pb.PatientProvider>): void;
  addPatientProviders(value?: heyrenee_v1_messages_patient_provider_pb.PatientProvider, index?: number): heyrenee_v1_messages_patient_provider_pb.PatientProvider;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientProvidersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientProvidersResponse): ListPatientProvidersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientProvidersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientProvidersResponse;
  static deserializeBinaryFromReader(message: ListPatientProvidersResponse, reader: jspb.BinaryReader): ListPatientProvidersResponse;
}

export namespace ListPatientProvidersResponse {
  export type AsObject = {
    patientProvidersList: Array<heyrenee_v1_messages_patient_provider_pb.PatientProvider.AsObject>,
  }
}

