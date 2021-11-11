// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_diagnosis.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_provider_pb from "../../../heyrenee/v1/messages/provider_pb";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_messages_diagnosis_pb from "../../../heyrenee/v1/messages/diagnosis_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class PatientDiagnosis extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  hasDiagnosisId(): boolean;
  clearDiagnosisId(): void;
  getDiagnosisId(): string;
  setDiagnosisId(value: string): void;

  hasDiagnosisMessage(): boolean;
  clearDiagnosisMessage(): void;
  getDiagnosisMessage(): heyrenee_v1_messages_diagnosis_pb.Diagnosis | undefined;
  setDiagnosisMessage(value?: heyrenee_v1_messages_diagnosis_pb.Diagnosis): void;

  getPatientDiagnosisId(): string;
  setPatientDiagnosisId(value: string): void;

  hasProviderId(): boolean;
  clearProviderId(): void;
  getProviderId(): string;
  setProviderId(value: string): void;

  hasProviderMessage(): boolean;
  clearProviderMessage(): void;
  getProviderMessage(): heyrenee_v1_messages_provider_pb.Provider | undefined;
  setProviderMessage(value?: heyrenee_v1_messages_provider_pb.Provider): void;

  getPatientDiagnosisStatus(): PatientDiagnosisStatusMap[keyof PatientDiagnosisStatusMap];
  setPatientDiagnosisStatus(value: PatientDiagnosisStatusMap[keyof PatientDiagnosisStatusMap]): void;

  hasDateDiagnosed(): boolean;
  clearDateDiagnosed(): void;
  getDateDiagnosed(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateDiagnosed(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getDiagnosisInstructions(): string;
  setDiagnosisInstructions(value: string): void;

  getDiagnosisNotes(): string;
  setDiagnosisNotes(value: string): void;

  getDiagnosisCase(): PatientDiagnosis.DiagnosisCase;
  getDiagnosingProviderCase(): PatientDiagnosis.DiagnosingProviderCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientDiagnosis.AsObject;
  static toObject(includeInstance: boolean, msg: PatientDiagnosis): PatientDiagnosis.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientDiagnosis, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientDiagnosis;
  static deserializeBinaryFromReader(message: PatientDiagnosis, reader: jspb.BinaryReader): PatientDiagnosis;
}

export namespace PatientDiagnosis {
  export type AsObject = {
    patientId: string,
    diagnosisId: string,
    diagnosisMessage?: heyrenee_v1_messages_diagnosis_pb.Diagnosis.AsObject,
    patientDiagnosisId: string,
    providerId: string,
    providerMessage?: heyrenee_v1_messages_provider_pb.Provider.AsObject,
    patientDiagnosisStatus: PatientDiagnosisStatusMap[keyof PatientDiagnosisStatusMap],
    dateDiagnosed?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    diagnosisInstructions: string,
    diagnosisNotes: string,
  }

  export enum DiagnosisCase {
    DIAGNOSIS_NOT_SET = 0,
    DIAGNOSIS_ID = 2,
    DIAGNOSIS_MESSAGE = 3,
  }

  export enum DiagnosingProviderCase {
    DIAGNOSING_PROVIDER_NOT_SET = 0,
    PROVIDER_ID = 5,
    PROVIDER_MESSAGE = 6,
  }
}

export interface PatientDiagnosisStatusMap {
  PATIENT_DIAGNOSIS_STATUS_UNSPECIFIED: 0;
  PATIENT_DIAGNOSIS_ACTIVE: 1;
  PATIENT_DIAGNOSIS_INACTIVE: 2;
}

export const PatientDiagnosisStatus: PatientDiagnosisStatusMap;

