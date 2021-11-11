// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/refill.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Refill extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getMedicationId(): string;
  setMedicationId(value: string): void;

  getPrescriptionId(): string;
  setPrescriptionId(value: string): void;

  getRefillId(): string;
  setRefillId(value: string): void;

  hasDateTimeRefilled(): boolean;
  clearDateTimeRefilled(): void;
  getDateTimeRefilled(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateTimeRefilled(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getPharmacyId(): string;
  setPharmacyId(value: string): void;

  getRefillType(): RefillTypeMap[keyof RefillTypeMap];
  setRefillType(value: RefillTypeMap[keyof RefillTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Refill.AsObject;
  static toObject(includeInstance: boolean, msg: Refill): Refill.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Refill, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Refill;
  static deserializeBinaryFromReader(message: Refill, reader: jspb.BinaryReader): Refill;
}

export namespace Refill {
  export type AsObject = {
    patientId: string,
    medicationId: string,
    prescriptionId: string,
    refillId: string,
    dateTimeRefilled?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pharmacyId: string,
    refillType: RefillTypeMap[keyof RefillTypeMap],
  }
}

export interface RefillTypeMap {
  REFILL_TYPE_UNSPECIFIED: 0;
  REFILL_DELIVERY: 1;
  REFILL_PICKUP: 2;
}

export const RefillType: RefillTypeMap;

