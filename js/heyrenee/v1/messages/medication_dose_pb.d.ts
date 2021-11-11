// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/medication_dose.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class MedicationDose extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getMedicationId(): string;
  setMedicationId(value: string): void;

  getPrescriptionId(): string;
  setPrescriptionId(value: string): void;

  getMedicationDoseId(): string;
  setMedicationDoseId(value: string): void;

  hasTimeTaken(): boolean;
  clearTimeTaken(): void;
  getTimeTaken(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimeTaken(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MedicationDose.AsObject;
  static toObject(includeInstance: boolean, msg: MedicationDose): MedicationDose.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MedicationDose, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MedicationDose;
  static deserializeBinaryFromReader(message: MedicationDose, reader: jspb.BinaryReader): MedicationDose;
}

export namespace MedicationDose {
  export type AsObject = {
    patientId: string,
    medicationId: string,
    prescriptionId: string,
    medicationDoseId: string,
    timeTaken?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

