// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_assessment.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_enums_language_pb from "../../../heyrenee/v1/enums/language_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class PatientAssessment extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientAssessmentId(): string;
  setPatientAssessmentId(value: string): void;

  getAssessmentType(): AssessmentTypeMap[keyof AssessmentTypeMap];
  setAssessmentType(value: AssessmentTypeMap[keyof AssessmentTypeMap]): void;

  getAssessmentLanguage(): heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap];
  setAssessmentLanguage(value: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap]): void;

  hasDateTimeTaken(): boolean;
  clearDateTimeTaken(): void;
  getDateTimeTaken(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateTimeTaken(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getScore(): number;
  setScore(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientAssessment.AsObject;
  static toObject(includeInstance: boolean, msg: PatientAssessment): PatientAssessment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientAssessment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientAssessment;
  static deserializeBinaryFromReader(message: PatientAssessment, reader: jspb.BinaryReader): PatientAssessment;
}

export namespace PatientAssessment {
  export type AsObject = {
    patientId: string,
    patientAssessmentId: string,
    assessmentType: AssessmentTypeMap[keyof AssessmentTypeMap],
    assessmentLanguage: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap],
    dateTimeTaken?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    score: number,
  }
}

export interface AssessmentTypeMap {
  ASSESSMENT_TYPE_UNSPECIFIED: 0;
  ASSESSMENT_TYPE_HEALTH_LITERACY: 1;
  ASSESSMENT_TYPE_TECHNOLOGY_LITERACY: 2;
}

export const AssessmentType: AssessmentTypeMap;

