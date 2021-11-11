// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient_satisfaction_questionnaire.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_enums_language_pb from "../../../heyrenee/v1/enums/language_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class PatientSatisfactionQuestionnaire extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientSatisfactionQuestionnaireId(): string;
  setPatientSatisfactionQuestionnaireId(value: string): void;

  getQuestionnaireLanguage(): heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap];
  setQuestionnaireLanguage(value: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap]): void;

  hasDateTimeAnswered(): boolean;
  clearDateTimeAnswered(): void;
  getDateTimeAnswered(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateTimeAnswered(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getHealthStatus(): number;
  setHealthStatus(value: number): void;

  getMedicalCareSatisfaction(): number;
  setMedicalCareSatisfaction(value: number): void;

  getHealthSatisfaction(): number;
  setHealthSatisfaction(value: number): void;

  getHeyReneeSatisfaction(): number;
  setHeyReneeSatisfaction(value: number): void;

  getPrimaryCareProviderSatisfaction(): number;
  setPrimaryCareProviderSatisfaction(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatientSatisfactionQuestionnaire.AsObject;
  static toObject(includeInstance: boolean, msg: PatientSatisfactionQuestionnaire): PatientSatisfactionQuestionnaire.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PatientSatisfactionQuestionnaire, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatientSatisfactionQuestionnaire;
  static deserializeBinaryFromReader(message: PatientSatisfactionQuestionnaire, reader: jspb.BinaryReader): PatientSatisfactionQuestionnaire;
}

export namespace PatientSatisfactionQuestionnaire {
  export type AsObject = {
    patientId: string,
    patientSatisfactionQuestionnaireId: string,
    questionnaireLanguage: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap],
    dateTimeAnswered?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    healthStatus: number,
    medicalCareSatisfaction: number,
    healthSatisfaction: number,
    heyReneeSatisfaction: number,
    primaryCareProviderSatisfaction: number,
  }
}

