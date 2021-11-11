// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/medication.proto

import * as jspb from "google-protobuf";

export class Medication extends jspb.Message {
  getMedicationId(): string;
  setMedicationId(value: string): void;

  getRxcui(): string;
  setRxcui(value: string): void;

  getGenericRxcui(): string;
  setGenericRxcui(value: string): void;

  getTermType(): string;
  setTermType(value: string): void;

  getFullName(): string;
  setFullName(value: string): void;

  getRxNormDoseForm(): string;
  setRxNormDoseForm(value: string): void;

  getFullGenericName(): string;
  setFullGenericName(value: string): void;

  getBrandName(): string;
  setBrandName(value: string): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  getRoute(): string;
  setRoute(value: string): void;

  getNewDoseForm(): string;
  setNewDoseForm(value: string): void;

  getStrength(): string;
  setStrength(value: string): void;

  getSuppressFor(): string;
  setSuppressFor(value: string): void;

  getDisplayNameSynonym(): string;
  setDisplayNameSynonym(value: string): void;

  getSxdgRxcui(): string;
  setSxdgRxcui(value: string): void;

  getSxdgTermType(): string;
  setSxdgTermType(value: string): void;

  getSxdgName(): string;
  setSxdgName(value: string): void;

  getPrescribableName(): string;
  setPrescribableName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Medication.AsObject;
  static toObject(includeInstance: boolean, msg: Medication): Medication.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Medication, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Medication;
  static deserializeBinaryFromReader(message: Medication, reader: jspb.BinaryReader): Medication;
}

export namespace Medication {
  export type AsObject = {
    medicationId: string,
    rxcui: string,
    genericRxcui: string,
    termType: string,
    fullName: string,
    rxNormDoseForm: string,
    fullGenericName: string,
    brandName: string,
    displayName: string,
    route: string,
    newDoseForm: string,
    strength: string,
    suppressFor: string,
    displayNameSynonym: string,
    sxdgRxcui: string,
    sxdgTermType: string,
    sxdgName: string,
    prescribableName: string,
  }
}

