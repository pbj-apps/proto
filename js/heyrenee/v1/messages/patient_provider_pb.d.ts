// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_provider.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_provider_pb from "../../../heyrenee/v1/messages/provider_pb";
import * as heyrenee_v1_messages_specialty_pb from "../../../heyrenee/v1/messages/specialty_pb";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class PatientProvider extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  hasProviderId(): boolean;
  clearProviderId(): void;
  getProviderId(): string;
  setProviderId(value: string): void;

  hasProviderMessage(): boolean;
  clearProviderMessage(): void;
  getProviderMessage(): heyrenee_v1_messages_provider_pb.Provider | undefined;
  setProviderMessage(value?: heyrenee_v1_messages_provider_pb.Provider): void;

  getCellPhone(): string;
  setCellPhone(value: string): void;

  getContactInstructions(): string;
  setContactInstructions(value: string): void;

  getContactInfo(): string;
  setContactInfo(value: string): void;

  getPatientProviderStatus(): PatientProviderStatusMap[keyof PatientProviderStatusMap];
  setPatientProviderStatus(value: PatientProviderStatusMap[keyof PatientProviderStatusMap]): void;

  getPatientProviderType(): PatientProviderTypeMap[keyof PatientProviderTypeMap];
  setPatientProviderType(value: PatientProviderTypeMap[keyof PatientProviderTypeMap]): void;

  getSpecialty(): heyrenee_v1_messages_specialty_pb.SpecialtyMap[keyof heyrenee_v1_messages_specialty_pb.SpecialtyMap];
  setSpecialty(value: heyrenee_v1_messages_specialty_pb.SpecialtyMap[keyof heyrenee_v1_messages_specialty_pb.SpecialtyMap]): void;

  getProviderCase(): PatientProvider.ProviderCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientProvider.AsObject;
  static toObject(includeInstance: boolean, msg: PatientProvider): PatientProvider.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientProvider, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientProvider;
  static deserializeBinaryFromReader(message: PatientProvider, reader: jspb.BinaryReader): PatientProvider;
}

export namespace PatientProvider {
  export type AsObject = {
    patientId: string,
    providerId: string,
    providerMessage?: heyrenee_v1_messages_provider_pb.Provider.AsObject,
    cellPhone: string,
    contactInstructions: string,
    contactInfo: string,
    patientProviderStatus: PatientProviderStatusMap[keyof PatientProviderStatusMap],
    patientProviderType: PatientProviderTypeMap[keyof PatientProviderTypeMap],
    specialty: heyrenee_v1_messages_specialty_pb.SpecialtyMap[keyof heyrenee_v1_messages_specialty_pb.SpecialtyMap],
  }

  export enum ProviderCase {
    PROVIDER_NOT_SET = 0,
    PROVIDER_ID = 2,
    PROVIDER_MESSAGE = 3,
  }
}

export interface PatientProviderStatusMap {
  PATIENT_PROVIDER_STATUS_UNSPECIFIED: 0;
  PATIENT_PROVIDER_ACTIVE: 1;
  PATIENT_PROVIDER_INACTIVE: 2;
}

export const PatientProviderStatus: PatientProviderStatusMap;

export interface PatientProviderTypeMap {
  PATIENT_PROVIDER_TYPE_UNSPECIFIED: 0;
  PATIENT_PROVIDER_TYPE_PRIMARY: 1;
}

export const PatientProviderType: PatientProviderTypeMap;

