// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_caregiver.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_caregiver_pb from "../../../heyrenee/v1/messages/caregiver_pb";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class PatientCaregiver extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  hasCaregiverId(): boolean;
  clearCaregiverId(): void;
  getCaregiverId(): string;
  setCaregiverId(value: string): void;

  hasCaregiverMessage(): boolean;
  clearCaregiverMessage(): void;
  getCaregiverMessage(): heyrenee_v1_messages_caregiver_pb.Caregiver | undefined;
  setCaregiverMessage(value?: heyrenee_v1_messages_caregiver_pb.Caregiver): void;

  getPreferredName(): string;
  setPreferredName(value: string): void;

  getPatientCaregiverType(): PatientCaregiverTypeMap[keyof PatientCaregiverTypeMap];
  setPatientCaregiverType(value: PatientCaregiverTypeMap[keyof PatientCaregiverTypeMap]): void;

  getPatientCaregiverAccess(): PatientCaregiverAccessMap[keyof PatientCaregiverAccessMap];
  setPatientCaregiverAccess(value: PatientCaregiverAccessMap[keyof PatientCaregiverAccessMap]): void;

  getPatientCaregiverRelationship(): PatientCaregiverRelationshipMap[keyof PatientCaregiverRelationshipMap];
  setPatientCaregiverRelationship(value: PatientCaregiverRelationshipMap[keyof PatientCaregiverRelationshipMap]): void;

  getCaregiverCase(): PatientCaregiver.CaregiverCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientCaregiver.AsObject;
  static toObject(includeInstance: boolean, msg: PatientCaregiver): PatientCaregiver.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientCaregiver, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientCaregiver;
  static deserializeBinaryFromReader(message: PatientCaregiver, reader: jspb.BinaryReader): PatientCaregiver;
}

export namespace PatientCaregiver {
  export type AsObject = {
    patientId: string,
    caregiverId: string,
    caregiverMessage?: heyrenee_v1_messages_caregiver_pb.Caregiver.AsObject,
    preferredName: string,
    patientCaregiverType: PatientCaregiverTypeMap[keyof PatientCaregiverTypeMap],
    patientCaregiverAccess: PatientCaregiverAccessMap[keyof PatientCaregiverAccessMap],
    patientCaregiverRelationship: PatientCaregiverRelationshipMap[keyof PatientCaregiverRelationshipMap],
  }

  export enum CaregiverCase {
    CAREGIVER_NOT_SET = 0,
    CAREGIVER_ID = 2,
    CAREGIVER_MESSAGE = 3,
  }
}

export interface PatientCaregiverTypeMap {
  PATIENT_CAREGIVER_TYPE_UNSPECIFIED: 0;
  PATIENT_CAREGIVER_TYPE_PRIMARY: 1;
}

export const PatientCaregiverType: PatientCaregiverTypeMap;

export interface PatientCaregiverAccessMap {
  PATIENT_CAREGIVER_ACCESS_UNSPECIFIED: 0;
  PATIENT_CAREGIVER_ACCESS_NONE: 1;
  PATIENT_CAREGIVER_ACCESS_DASHBOARD: 2;
}

export const PatientCaregiverAccess: PatientCaregiverAccessMap;

export interface PatientCaregiverRelationshipMap {
  PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED: 0;
  PATIENT_CAREGIVER_RELATIONSHIP_CHILD: 1;
  PATIENT_CAREGIVER_RELATIONSHIP_PARENT: 2;
  PATIENT_CAREGIVER_RELATIONSHIP_GRANDPARENT: 3;
  PATIENT_CAREGIVER_RELATIONSHIP_GRANDCHILD: 4;
  PATIENT_CAREGIVER_RELATIONSHIP_EXTENDED_FAMILY: 5;
  PATIENT_CAREGIVER_RELATIONSHIP_FRIEND: 6;
  PATIENT_CAREGIVER_RELATIONSHIP_HEALTH_AIDE: 7;
  PATIENT_CAREGIVER_RELATIONSHIP_OTHER: 8;
}

export const PatientCaregiverRelationship: PatientCaregiverRelationshipMap;

