// package: heyrenee.v1
// file: heyrenee/v1/medication_service.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_messages_medication_dose_pb from "../../heyrenee/v1/messages/medication_dose_pb";
import * as heyrenee_v1_messages_prescription_pb from "../../heyrenee/v1/messages/prescription_pb";
import * as heyrenee_v1_messages_refill_pb from "../../heyrenee/v1/messages/refill_pb";
import * as heyrenee_v1_options_auth_pb from "../../heyrenee/v1/options/auth_pb";

export class MedicationSuggestRequest extends jspb.Message {
  getPartialText(): string;
  setPartialText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MedicationSuggestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MedicationSuggestRequest): MedicationSuggestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MedicationSuggestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MedicationSuggestRequest;
  static deserializeBinaryFromReader(message: MedicationSuggestRequest, reader: jspb.BinaryReader): MedicationSuggestRequest;
}

export namespace MedicationSuggestRequest {
  export type AsObject = {
    partialText: string,
  }
}

export class MedicationSuggestResponse extends jspb.Message {
  clearMedicationSuggestionsList(): void;
  getMedicationSuggestionsList(): Array<MedicationSuggestion>;
  setMedicationSuggestionsList(value: Array<MedicationSuggestion>): void;
  addMedicationSuggestions(value?: MedicationSuggestion, index?: number): MedicationSuggestion;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MedicationSuggestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MedicationSuggestResponse): MedicationSuggestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MedicationSuggestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MedicationSuggestResponse;
  static deserializeBinaryFromReader(message: MedicationSuggestResponse, reader: jspb.BinaryReader): MedicationSuggestResponse;
}

export namespace MedicationSuggestResponse {
  export type AsObject = {
    medicationSuggestionsList: Array<MedicationSuggestion.AsObject>,
  }
}

export class MedicationSuggestion extends jspb.Message {
  getMedicationId(): string;
  setMedicationId(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MedicationSuggestion.AsObject;
  static toObject(includeInstance: boolean, msg: MedicationSuggestion): MedicationSuggestion.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MedicationSuggestion, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MedicationSuggestion;
  static deserializeBinaryFromReader(message: MedicationSuggestion, reader: jspb.BinaryReader): MedicationSuggestion;
}

export namespace MedicationSuggestion {
  export type AsObject = {
    medicationId: string,
    displayName: string,
  }
}

export class CreatePrescriptionRequest extends jspb.Message {
  hasPrescription(): boolean;
  clearPrescription(): void;
  getPrescription(): heyrenee_v1_messages_prescription_pb.Prescription | undefined;
  setPrescription(value?: heyrenee_v1_messages_prescription_pb.Prescription): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePrescriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePrescriptionRequest): CreatePrescriptionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePrescriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePrescriptionRequest;
  static deserializeBinaryFromReader(message: CreatePrescriptionRequest, reader: jspb.BinaryReader): CreatePrescriptionRequest;
}

export namespace CreatePrescriptionRequest {
  export type AsObject = {
    prescription?: heyrenee_v1_messages_prescription_pb.Prescription.AsObject,
  }
}

export class UpdatePrescriptionRequest extends jspb.Message {
  hasPrescription(): boolean;
  clearPrescription(): void;
  getPrescription(): heyrenee_v1_messages_prescription_pb.Prescription | undefined;
  setPrescription(value?: heyrenee_v1_messages_prescription_pb.Prescription): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePrescriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePrescriptionRequest): UpdatePrescriptionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePrescriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePrescriptionRequest;
  static deserializeBinaryFromReader(message: UpdatePrescriptionRequest, reader: jspb.BinaryReader): UpdatePrescriptionRequest;
}

export namespace UpdatePrescriptionRequest {
  export type AsObject = {
    prescription?: heyrenee_v1_messages_prescription_pb.Prescription.AsObject,
  }
}

export class ListPrescriptionsRequest extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getPrescriptionStatus(): heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap[keyof heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap];
  setPrescriptionStatus(value: heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap[keyof heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPrescriptionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPrescriptionsRequest): ListPrescriptionsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPrescriptionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPrescriptionsRequest;
  static deserializeBinaryFromReader(message: ListPrescriptionsRequest, reader: jspb.BinaryReader): ListPrescriptionsRequest;
}

export namespace ListPrescriptionsRequest {
  export type AsObject = {
    patientId: string,
    prescriptionStatus: heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap[keyof heyrenee_v1_messages_prescription_pb.PrescriptionStatusMap],
  }
}

export class ListPrescriptionsResponse extends jspb.Message {
  clearPrescriptionsList(): void;
  getPrescriptionsList(): Array<heyrenee_v1_messages_prescription_pb.Prescription>;
  setPrescriptionsList(value: Array<heyrenee_v1_messages_prescription_pb.Prescription>): void;
  addPrescriptions(value?: heyrenee_v1_messages_prescription_pb.Prescription, index?: number): heyrenee_v1_messages_prescription_pb.Prescription;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPrescriptionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPrescriptionsResponse): ListPrescriptionsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPrescriptionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPrescriptionsResponse;
  static deserializeBinaryFromReader(message: ListPrescriptionsResponse, reader: jspb.BinaryReader): ListPrescriptionsResponse;
}

export namespace ListPrescriptionsResponse {
  export type AsObject = {
    prescriptionsList: Array<heyrenee_v1_messages_prescription_pb.Prescription.AsObject>,
  }
}

export class ListMedicationDosesRequest extends jspb.Message {
  getPrescriptionId(): string;
  setPrescriptionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListMedicationDosesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListMedicationDosesRequest): ListMedicationDosesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListMedicationDosesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListMedicationDosesRequest;
  static deserializeBinaryFromReader(message: ListMedicationDosesRequest, reader: jspb.BinaryReader): ListMedicationDosesRequest;
}

export namespace ListMedicationDosesRequest {
  export type AsObject = {
    prescriptionId: string,
  }
}

export class ListMedicationDosesResponse extends jspb.Message {
  clearMedicationDosesList(): void;
  getMedicationDosesList(): Array<heyrenee_v1_messages_medication_dose_pb.MedicationDose>;
  setMedicationDosesList(value: Array<heyrenee_v1_messages_medication_dose_pb.MedicationDose>): void;
  addMedicationDoses(value?: heyrenee_v1_messages_medication_dose_pb.MedicationDose, index?: number): heyrenee_v1_messages_medication_dose_pb.MedicationDose;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListMedicationDosesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListMedicationDosesResponse): ListMedicationDosesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListMedicationDosesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListMedicationDosesResponse;
  static deserializeBinaryFromReader(message: ListMedicationDosesResponse, reader: jspb.BinaryReader): ListMedicationDosesResponse;
}

export namespace ListMedicationDosesResponse {
  export type AsObject = {
    medicationDosesList: Array<heyrenee_v1_messages_medication_dose_pb.MedicationDose.AsObject>,
  }
}

export class ListRefillsRequest extends jspb.Message {
  getPrescriptionId(): string;
  setPrescriptionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRefillsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRefillsRequest): ListRefillsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRefillsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRefillsRequest;
  static deserializeBinaryFromReader(message: ListRefillsRequest, reader: jspb.BinaryReader): ListRefillsRequest;
}

export namespace ListRefillsRequest {
  export type AsObject = {
    prescriptionId: string,
  }
}

export class ListRefillsResponse extends jspb.Message {
  clearRefillsList(): void;
  getRefillsList(): Array<heyrenee_v1_messages_refill_pb.Refill>;
  setRefillsList(value: Array<heyrenee_v1_messages_refill_pb.Refill>): void;
  addRefills(value?: heyrenee_v1_messages_refill_pb.Refill, index?: number): heyrenee_v1_messages_refill_pb.Refill;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRefillsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRefillsResponse): ListRefillsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRefillsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRefillsResponse;
  static deserializeBinaryFromReader(message: ListRefillsResponse, reader: jspb.BinaryReader): ListRefillsResponse;
}

export namespace ListRefillsResponse {
  export type AsObject = {
    refillsList: Array<heyrenee_v1_messages_refill_pb.Refill.AsObject>,
  }
}

