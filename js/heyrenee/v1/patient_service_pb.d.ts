// package: heyrenee.v1
// file: heyrenee/v1/patient_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_patient_caregiver_pb from "../../heyrenee/v1/messages/patient_caregiver_pb";
import * as heyrenee_v1_messages_patient_health_questionnaire_pb from "../../heyrenee/v1/messages/patient_health_questionnaire_pb";
import * as heyrenee_v1_messages_patient_satisfaction_questionnaire_pb from "../../heyrenee/v1/messages/patient_satisfaction_questionnaire_pb";
import * as heyrenee_v1_messages_patient_sdoh_questionnaire_pb from "../../heyrenee/v1/messages/patient_sdoh_questionnaire_pb";
import * as heyrenee_v1_messages_patient_assessment_pb from "../../heyrenee/v1/messages/patient_assessment_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class CreatePatientCaregiverRequest extends jspb.Message {
  hasPatientCaregiver(): boolean;
  clearPatientCaregiver(): void;
  getPatientCaregiver(): heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver | undefined;
  setPatientCaregiver(value?: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientCaregiverRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientCaregiverRequest): CreatePatientCaregiverRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientCaregiverRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientCaregiverRequest;
  static deserializeBinaryFromReader(message: CreatePatientCaregiverRequest, reader: jspb.BinaryReader): CreatePatientCaregiverRequest;
}

export namespace CreatePatientCaregiverRequest {
  export type AsObject = {
    patientCaregiver?: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver.AsObject,
  }
}

export class UpdatePatientCaregiverRequest extends jspb.Message {
  hasPatientCaregiver(): boolean;
  clearPatientCaregiver(): void;
  getPatientCaregiver(): heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver | undefined;
  setPatientCaregiver(value?: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePatientCaregiverRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePatientCaregiverRequest): UpdatePatientCaregiverRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePatientCaregiverRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePatientCaregiverRequest;
  static deserializeBinaryFromReader(message: UpdatePatientCaregiverRequest, reader: jspb.BinaryReader): UpdatePatientCaregiverRequest;
}

export namespace UpdatePatientCaregiverRequest {
  export type AsObject = {
    patientCaregiver?: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver.AsObject,
  }
}

export class ListPatientCaregiversRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientCaregiversRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientCaregiversRequest): ListPatientCaregiversRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientCaregiversRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientCaregiversRequest;
  static deserializeBinaryFromReader(message: ListPatientCaregiversRequest, reader: jspb.BinaryReader): ListPatientCaregiversRequest;
}

export namespace ListPatientCaregiversRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class ListPatientCaregiversResponse extends jspb.Message {
  clearPatientCaregiversList(): void;
  getPatientCaregiversList(): Array<heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver>;
  setPatientCaregiversList(value: Array<heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver>): void;
  addPatientCaregivers(value?: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver, index?: number): heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientCaregiversResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientCaregiversResponse): ListPatientCaregiversResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientCaregiversResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientCaregiversResponse;
  static deserializeBinaryFromReader(message: ListPatientCaregiversResponse, reader: jspb.BinaryReader): ListPatientCaregiversResponse;
}

export namespace ListPatientCaregiversResponse {
  export type AsObject = {
    patientCaregiversList: Array<heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver.AsObject>,
  }
}

export class CreatePatientAssessmentRequest extends jspb.Message {
  hasPatientAssessment(): boolean;
  clearPatientAssessment(): void;
  getPatientAssessment(): heyrenee_v1_messages_patient_assessment_pb.PatientAssessment | undefined;
  setPatientAssessment(value?: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientAssessmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientAssessmentRequest): CreatePatientAssessmentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientAssessmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientAssessmentRequest;
  static deserializeBinaryFromReader(message: CreatePatientAssessmentRequest, reader: jspb.BinaryReader): CreatePatientAssessmentRequest;
}

export namespace CreatePatientAssessmentRequest {
  export type AsObject = {
    patientAssessment?: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment.AsObject,
  }
}

export class ListPatientAssessmentsRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientAssessmentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientAssessmentsRequest): ListPatientAssessmentsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientAssessmentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientAssessmentsRequest;
  static deserializeBinaryFromReader(message: ListPatientAssessmentsRequest, reader: jspb.BinaryReader): ListPatientAssessmentsRequest;
}

export namespace ListPatientAssessmentsRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class ListPatientAssessmentsResponse extends jspb.Message {
  clearPatientAssessmentsList(): void;
  getPatientAssessmentsList(): Array<heyrenee_v1_messages_patient_assessment_pb.PatientAssessment>;
  setPatientAssessmentsList(value: Array<heyrenee_v1_messages_patient_assessment_pb.PatientAssessment>): void;
  addPatientAssessments(value?: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment, index?: number): heyrenee_v1_messages_patient_assessment_pb.PatientAssessment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientAssessmentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientAssessmentsResponse): ListPatientAssessmentsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientAssessmentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientAssessmentsResponse;
  static deserializeBinaryFromReader(message: ListPatientAssessmentsResponse, reader: jspb.BinaryReader): ListPatientAssessmentsResponse;
}

export namespace ListPatientAssessmentsResponse {
  export type AsObject = {
    patientAssessmentsList: Array<heyrenee_v1_messages_patient_assessment_pb.PatientAssessment.AsObject>,
  }
}

export class CreatePatientSatisfactionQuestionnaireRequest extends jspb.Message {
  hasPatientSatisfactionQuestionnaire(): boolean;
  clearPatientSatisfactionQuestionnaire(): void;
  getPatientSatisfactionQuestionnaire(): heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire | undefined;
  setPatientSatisfactionQuestionnaire(value?: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientSatisfactionQuestionnaireRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientSatisfactionQuestionnaireRequest): CreatePatientSatisfactionQuestionnaireRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientSatisfactionQuestionnaireRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientSatisfactionQuestionnaireRequest;
  static deserializeBinaryFromReader(message: CreatePatientSatisfactionQuestionnaireRequest, reader: jspb.BinaryReader): CreatePatientSatisfactionQuestionnaireRequest;
}

export namespace CreatePatientSatisfactionQuestionnaireRequest {
  export type AsObject = {
    patientSatisfactionQuestionnaire?: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire.AsObject,
  }
}

export class ListPatientSatisfactionQuestionnairesRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientSatisfactionQuestionnairesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientSatisfactionQuestionnairesRequest): ListPatientSatisfactionQuestionnairesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientSatisfactionQuestionnairesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientSatisfactionQuestionnairesRequest;
  static deserializeBinaryFromReader(message: ListPatientSatisfactionQuestionnairesRequest, reader: jspb.BinaryReader): ListPatientSatisfactionQuestionnairesRequest;
}

export namespace ListPatientSatisfactionQuestionnairesRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class ListPatientSatisfactionQuestionnairesResponse extends jspb.Message {
  clearPatientSatisfactionQuestionnairesList(): void;
  getPatientSatisfactionQuestionnairesList(): Array<heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire>;
  setPatientSatisfactionQuestionnairesList(value: Array<heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire>): void;
  addPatientSatisfactionQuestionnaires(value?: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire, index?: number): heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientSatisfactionQuestionnairesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientSatisfactionQuestionnairesResponse): ListPatientSatisfactionQuestionnairesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientSatisfactionQuestionnairesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientSatisfactionQuestionnairesResponse;
  static deserializeBinaryFromReader(message: ListPatientSatisfactionQuestionnairesResponse, reader: jspb.BinaryReader): ListPatientSatisfactionQuestionnairesResponse;
}

export namespace ListPatientSatisfactionQuestionnairesResponse {
  export type AsObject = {
    patientSatisfactionQuestionnairesList: Array<heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire.AsObject>,
  }
}

export class CreatePatientHealthQuestionnaireRequest extends jspb.Message {
  hasPatientHealthQuestionnaire(): boolean;
  clearPatientHealthQuestionnaire(): void;
  getPatientHealthQuestionnaire(): heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire | undefined;
  setPatientHealthQuestionnaire(value?: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientHealthQuestionnaireRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientHealthQuestionnaireRequest): CreatePatientHealthQuestionnaireRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientHealthQuestionnaireRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientHealthQuestionnaireRequest;
  static deserializeBinaryFromReader(message: CreatePatientHealthQuestionnaireRequest, reader: jspb.BinaryReader): CreatePatientHealthQuestionnaireRequest;
}

export namespace CreatePatientHealthQuestionnaireRequest {
  export type AsObject = {
    patientHealthQuestionnaire?: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire.AsObject,
  }
}

export class ListPatientHealthQuestionnairesRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientHealthQuestionnairesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientHealthQuestionnairesRequest): ListPatientHealthQuestionnairesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientHealthQuestionnairesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientHealthQuestionnairesRequest;
  static deserializeBinaryFromReader(message: ListPatientHealthQuestionnairesRequest, reader: jspb.BinaryReader): ListPatientHealthQuestionnairesRequest;
}

export namespace ListPatientHealthQuestionnairesRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class ListPatientHealthQuestionnairesResponse extends jspb.Message {
  clearPatientHealthQuestionnairesList(): void;
  getPatientHealthQuestionnairesList(): Array<heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire>;
  setPatientHealthQuestionnairesList(value: Array<heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire>): void;
  addPatientHealthQuestionnaires(value?: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire, index?: number): heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientHealthQuestionnairesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientHealthQuestionnairesResponse): ListPatientHealthQuestionnairesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientHealthQuestionnairesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientHealthQuestionnairesResponse;
  static deserializeBinaryFromReader(message: ListPatientHealthQuestionnairesResponse, reader: jspb.BinaryReader): ListPatientHealthQuestionnairesResponse;
}

export namespace ListPatientHealthQuestionnairesResponse {
  export type AsObject = {
    patientHealthQuestionnairesList: Array<heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire.AsObject>,
  }
}

export class CreatePatientSdohQuestionnaireRequest extends jspb.Message {
  hasPatientSdohQuestionnaire(): boolean;
  clearPatientSdohQuestionnaire(): void;
  getPatientSdohQuestionnaire(): heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire | undefined;
  setPatientSdohQuestionnaire(value?: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientSdohQuestionnaireRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientSdohQuestionnaireRequest): CreatePatientSdohQuestionnaireRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientSdohQuestionnaireRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientSdohQuestionnaireRequest;
  static deserializeBinaryFromReader(message: CreatePatientSdohQuestionnaireRequest, reader: jspb.BinaryReader): CreatePatientSdohQuestionnaireRequest;
}

export namespace CreatePatientSdohQuestionnaireRequest {
  export type AsObject = {
    patientSdohQuestionnaire?: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire.AsObject,
  }
}

export class ListPatientSdohQuestionnairesRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientSdohQuestionnairesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientSdohQuestionnairesRequest): ListPatientSdohQuestionnairesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientSdohQuestionnairesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientSdohQuestionnairesRequest;
  static deserializeBinaryFromReader(message: ListPatientSdohQuestionnairesRequest, reader: jspb.BinaryReader): ListPatientSdohQuestionnairesRequest;
}

export namespace ListPatientSdohQuestionnairesRequest {
  export type AsObject = {
    patientId: string,
  }
}

export class ListPatientSdohQuestionnairesResponse extends jspb.Message {
  clearPatientSdohQuestionnairesList(): void;
  getPatientSdohQuestionnairesList(): Array<heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire>;
  setPatientSdohQuestionnairesList(value: Array<heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire>): void;
  addPatientSdohQuestionnaires(value?: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire, index?: number): heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientSdohQuestionnairesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientSdohQuestionnairesResponse): ListPatientSdohQuestionnairesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientSdohQuestionnairesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientSdohQuestionnairesResponse;
  static deserializeBinaryFromReader(message: ListPatientSdohQuestionnairesResponse, reader: jspb.BinaryReader): ListPatientSdohQuestionnairesResponse;
}

export namespace ListPatientSdohQuestionnairesResponse {
  export type AsObject = {
    patientSdohQuestionnairesList: Array<heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire.AsObject>,
  }
}

