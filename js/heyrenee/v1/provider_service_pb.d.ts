// package: heyrenee.v1
// file: heyrenee/v1/provider_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_patient_provider_pb from "../../heyrenee/v1/messages/patient_provider_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class ProviderSuggestRequest extends jspb.Message {
  getPartialText(): string;
  setPartialText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProviderSuggestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ProviderSuggestRequest): ProviderSuggestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProviderSuggestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProviderSuggestRequest;
  static deserializeBinaryFromReader(message: ProviderSuggestRequest, reader: jspb.BinaryReader): ProviderSuggestRequest;
}

export namespace ProviderSuggestRequest {
  export type AsObject = {
    partialText: string,
  }
}

export class ProviderSuggestResponse extends jspb.Message {
  clearProviderSuggestionsList(): void;
  getProviderSuggestionsList(): Array<ProviderSuggestion>;
  setProviderSuggestionsList(value: Array<ProviderSuggestion>): void;
  addProviderSuggestions(value?: ProviderSuggestion, index?: number): ProviderSuggestion;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProviderSuggestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ProviderSuggestResponse): ProviderSuggestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProviderSuggestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProviderSuggestResponse;
  static deserializeBinaryFromReader(message: ProviderSuggestResponse, reader: jspb.BinaryReader): ProviderSuggestResponse;
}

export namespace ProviderSuggestResponse {
  export type AsObject = {
    providerSuggestionsList: Array<ProviderSuggestion.AsObject>,
  }
}

export class ProviderSuggestion extends jspb.Message {
  getProviderId(): string;
  setProviderId(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProviderSuggestion.AsObject;
  static toObject(includeInstance: boolean, msg: ProviderSuggestion): ProviderSuggestion.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProviderSuggestion, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProviderSuggestion;
  static deserializeBinaryFromReader(message: ProviderSuggestion, reader: jspb.BinaryReader): ProviderSuggestion;
}

export namespace ProviderSuggestion {
  export type AsObject = {
    providerId: string,
    displayName: string,
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

