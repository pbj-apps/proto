// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/caregiver.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";

export class Caregiver extends jspb.Message {
  getCaregiverId(): string;
  setCaregiverId(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getPrimaryPhone(): string;
  setPrimaryPhone(value: string): void;

  getMobilePhone(): string;
  setMobilePhone(value: string): void;

  getOtherPhone(): string;
  setOtherPhone(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  clearAddressLinesList(): void;
  getAddressLinesList(): Array<string>;
  setAddressLinesList(value: Array<string>): void;
  addAddressLines(value: string, index?: number): string;

  getCity(): string;
  setCity(value: string): void;

  getState(): string;
  setState(value: string): void;

  getZip(): string;
  setZip(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Caregiver.AsObject;
  static toObject(includeInstance: boolean, msg: Caregiver): Caregiver.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Caregiver, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Caregiver;
  static deserializeBinaryFromReader(message: Caregiver, reader: jspb.BinaryReader): Caregiver;
}

export namespace Caregiver {
  export type AsObject = {
    caregiverId: string,
    firstName: string,
    lastName: string,
    primaryPhone: string,
    mobilePhone: string,
    otherPhone: string,
    email: string,
    addressLinesList: Array<string>,
    city: string,
    state: string,
    zip: string,
  }
}

