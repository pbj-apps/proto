// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/diagnosis.proto

import * as jspb from "google-protobuf";

export class Diagnosis extends jspb.Message {
  getDiagnosisId(): string;
  setDiagnosisId(value: string): void;

  getIcd10cm(): string;
  setIcd10cm(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Diagnosis.AsObject;
  static toObject(includeInstance: boolean, msg: Diagnosis): Diagnosis.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Diagnosis, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Diagnosis;
  static deserializeBinaryFromReader(message: Diagnosis, reader: jspb.BinaryReader): Diagnosis;
}

export namespace Diagnosis {
  export type AsObject = {
    diagnosisId: string,
    icd10cm: string,
    name: string,
  }
}

