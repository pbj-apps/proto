// package: heyrenee.v1
// file: heyrenee/v1/diagnosis_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_patient_diagnosis_pb from "../../heyrenee/v1/messages/patient_diagnosis_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class DiagnosisSuggestRequest extends jspb.Message {
  getPartialText(): string;
  setPartialText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DiagnosisSuggestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DiagnosisSuggestRequest): DiagnosisSuggestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DiagnosisSuggestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DiagnosisSuggestRequest;
  static deserializeBinaryFromReader(message: DiagnosisSuggestRequest, reader: jspb.BinaryReader): DiagnosisSuggestRequest;
}

export namespace DiagnosisSuggestRequest {
  export type AsObject = {
    partialText: string,
  }
}

export class DiagnosisSuggestResponse extends jspb.Message {
  clearDiagnosisSuggestionsList(): void;
  getDiagnosisSuggestionsList(): Array<DiagnosisSuggestion>;
  setDiagnosisSuggestionsList(value: Array<DiagnosisSuggestion>): void;
  addDiagnosisSuggestions(value?: DiagnosisSuggestion, index?: number): DiagnosisSuggestion;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DiagnosisSuggestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DiagnosisSuggestResponse): DiagnosisSuggestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DiagnosisSuggestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DiagnosisSuggestResponse;
  static deserializeBinaryFromReader(message: DiagnosisSuggestResponse, reader: jspb.BinaryReader): DiagnosisSuggestResponse;
}

export namespace DiagnosisSuggestResponse {
  export type AsObject = {
    diagnosisSuggestionsList: Array<DiagnosisSuggestion.AsObject>,
  }
}

export class DiagnosisSuggestion extends jspb.Message {
  getDiagnosisId(): string;
  setDiagnosisId(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DiagnosisSuggestion.AsObject;
  static toObject(includeInstance: boolean, msg: DiagnosisSuggestion): DiagnosisSuggestion.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DiagnosisSuggestion, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DiagnosisSuggestion;
  static deserializeBinaryFromReader(message: DiagnosisSuggestion, reader: jspb.BinaryReader): DiagnosisSuggestion;
}

export namespace DiagnosisSuggestion {
  export type AsObject = {
    diagnosisId: string,
    displayName: string,
  }
}

export class CreatePatientDiagnosisRequest extends jspb.Message {
  hasPatientDiagnosis(): boolean;
  clearPatientDiagnosis(): void;
  getPatientDiagnosis(): heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis | undefined;
  setPatientDiagnosis(value?: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePatientDiagnosisRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePatientDiagnosisRequest): CreatePatientDiagnosisRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePatientDiagnosisRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePatientDiagnosisRequest;
  static deserializeBinaryFromReader(message: CreatePatientDiagnosisRequest, reader: jspb.BinaryReader): CreatePatientDiagnosisRequest;
}

export namespace CreatePatientDiagnosisRequest {
  export type AsObject = {
    patientDiagnosis?: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis.AsObject,
  }
}

export class UpdatePatientDiagnosisRequest extends jspb.Message {
  hasPatientDiagnosis(): boolean;
  clearPatientDiagnosis(): void;
  getPatientDiagnosis(): heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis | undefined;
  setPatientDiagnosis(value?: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePatientDiagnosisRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePatientDiagnosisRequest): UpdatePatientDiagnosisRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePatientDiagnosisRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePatientDiagnosisRequest;
  static deserializeBinaryFromReader(message: UpdatePatientDiagnosisRequest, reader: jspb.BinaryReader): UpdatePatientDiagnosisRequest;
}

export namespace UpdatePatientDiagnosisRequest {
  export type AsObject = {
    patientDiagnosis?: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis.AsObject,
  }
}

export class ListPatientDiagnosesRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPatientDiagnosisStatus(): heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap[keyof heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap];
  setPatientDiagnosisStatus(value: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap[keyof heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientDiagnosesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientDiagnosesRequest): ListPatientDiagnosesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientDiagnosesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientDiagnosesRequest;
  static deserializeBinaryFromReader(message: ListPatientDiagnosesRequest, reader: jspb.BinaryReader): ListPatientDiagnosesRequest;
}

export namespace ListPatientDiagnosesRequest {
  export type AsObject = {
    patientId: string,
    patientDiagnosisStatus: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap[keyof heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosisStatusMap],
  }
}

export class ListPatientDiagnosesResponse extends jspb.Message {
  clearPatientDiagnosesList(): void;
  getPatientDiagnosesList(): Array<heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis>;
  setPatientDiagnosesList(value: Array<heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis>): void;
  addPatientDiagnoses(value?: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis, index?: number): heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPatientDiagnosesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPatientDiagnosesResponse): ListPatientDiagnosesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPatientDiagnosesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPatientDiagnosesResponse;
  static deserializeBinaryFromReader(message: ListPatientDiagnosesResponse, reader: jspb.BinaryReader): ListPatientDiagnosesResponse;
}

export namespace ListPatientDiagnosesResponse {
  export type AsObject = {
    patientDiagnosesList: Array<heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis.AsObject>,
  }
}

