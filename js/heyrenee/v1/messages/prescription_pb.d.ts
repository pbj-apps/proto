// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/prescription.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as heyrenee_v1_messages_medication_pb from "../../../heyrenee/v1/messages/medication_pb";
import * as heyrenee_v1_messages_provider_pb from "../../../heyrenee/v1/messages/provider_pb";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class Prescription extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  hasMedicationId(): boolean;
  clearMedicationId(): void;
  getMedicationId(): string;
  setMedicationId(value: string): void;

  hasMedicationMessage(): boolean;
  clearMedicationMessage(): void;
  getMedicationMessage(): heyrenee_v1_messages_medication_pb.Medication | undefined;
  setMedicationMessage(value?: heyrenee_v1_messages_medication_pb.Medication): void;

  getPrescriptionId(): string;
  setPrescriptionId(value: string): void;

  hasProviderId(): boolean;
  clearProviderId(): void;
  getProviderId(): string;
  setProviderId(value: string): void;

  hasProviderMessage(): boolean;
  clearProviderMessage(): void;
  getProviderMessage(): heyrenee_v1_messages_provider_pb.Provider | undefined;
  setProviderMessage(value?: heyrenee_v1_messages_provider_pb.Provider): void;

  getProviderInstructions(): string;
  setProviderInstructions(value: string): void;

  getPatientInstructions(): string;
  setPatientInstructions(value: string): void;

  getRefillCount(): number;
  setRefillCount(value: number): void;

  hasRefillFrequency(): boolean;
  clearRefillFrequency(): void;
  getRefillFrequency(): google_protobuf_duration_pb.Duration | undefined;
  setRefillFrequency(value?: google_protobuf_duration_pb.Duration): void;

  getPrescriptionStatus(): PrescriptionStatusMap[keyof PrescriptionStatusMap];
  setPrescriptionStatus(value: PrescriptionStatusMap[keyof PrescriptionStatusMap]): void;

  hasFirstMedicationRegimenStart(): boolean;
  clearFirstMedicationRegimenStart(): void;
  getFirstMedicationRegimenStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFirstMedicationRegimenStart(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasMedicationRegimenDuration(): boolean;
  clearMedicationRegimenDuration(): void;
  getMedicationRegimenDuration(): google_protobuf_duration_pb.Duration | undefined;
  setMedicationRegimenDuration(value?: google_protobuf_duration_pb.Duration): void;

  getMedicationDosesPerRegimen(): number;
  setMedicationDosesPerRegimen(value: number): void;

  clearMedicationDoseDurationsFromRegimenStartList(): void;
  getMedicationDoseDurationsFromRegimenStartList(): Array<google_protobuf_duration_pb.Duration>;
  setMedicationDoseDurationsFromRegimenStartList(value: Array<google_protobuf_duration_pb.Duration>): void;
  addMedicationDoseDurationsFromRegimenStart(value?: google_protobuf_duration_pb.Duration, index?: number): google_protobuf_duration_pb.Duration;

  hasDateWritten(): boolean;
  clearDateWritten(): void;
  getDateWritten(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateWritten(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getMedicationCase(): Prescription.MedicationCase;
  getProviderCase(): Prescription.ProviderCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Prescription.AsObject;
  static toObject(includeInstance: boolean, msg: Prescription): Prescription.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Prescription, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Prescription;
  static deserializeBinaryFromReader(message: Prescription, reader: jspb.BinaryReader): Prescription;
}

export namespace Prescription {
  export type AsObject = {
    patientId: string,
    medicationId: string,
    medicationMessage?: heyrenee_v1_messages_medication_pb.Medication.AsObject,
    prescriptionId: string,
    providerId: string,
    providerMessage?: heyrenee_v1_messages_provider_pb.Provider.AsObject,
    providerInstructions: string,
    patientInstructions: string,
    refillCount: number,
    refillFrequency?: google_protobuf_duration_pb.Duration.AsObject,
    prescriptionStatus: PrescriptionStatusMap[keyof PrescriptionStatusMap],
    firstMedicationRegimenStart?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    medicationRegimenDuration?: google_protobuf_duration_pb.Duration.AsObject,
    medicationDosesPerRegimen: number,
    medicationDoseDurationsFromRegimenStartList: Array<google_protobuf_duration_pb.Duration.AsObject>,
    dateWritten?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export enum MedicationCase {
    MEDICATION_NOT_SET = 0,
    MEDICATION_ID = 2,
    MEDICATION_MESSAGE = 3,
  }

  export enum ProviderCase {
    PROVIDER_NOT_SET = 0,
    PROVIDER_ID = 5,
    PROVIDER_MESSAGE = 6,
  }
}

export interface PrescriptionStatusMap {
  PRESCRIPTION_STATUS_UNSPECIFIED: 0;
  PRESCRIPTION_ACTIVE: 1;
  PRESCRIPTION_INACTIVE: 2;
}

export const PrescriptionStatus: PrescriptionStatusMap;

