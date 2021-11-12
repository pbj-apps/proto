// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/provider.proto

import * as jspb from "google-protobuf";
import * as heyrenee_v1_enums_specialty_pb from "../../../heyrenee/v1/enums/specialty_pb";

export class Provider extends jspb.Message {
  getProviderId(): string;
  setProviderId(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

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

  getPhone(): string;
  setPhone(value: string): void;

  getSpecialty(): heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap];
  setSpecialty(value: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap]): void;

  getNpi(): string;
  setNpi(value: string): void;

  getFacility(): string;
  setFacility(value: string): void;

  getHours(): string;
  setHours(value: string): void;

  getRibbonId(): string;
  setRibbonId(value: string): void;

  getSecondaryFacility(): string;
  setSecondaryFacility(value: string): void;

  clearSecondaryAddressLinesList(): void;
  getSecondaryAddressLinesList(): Array<string>;
  setSecondaryAddressLinesList(value: Array<string>): void;
  addSecondaryAddressLines(value: string, index?: number): string;

  getSecondaryCity(): string;
  setSecondaryCity(value: string): void;

  getSecondaryState(): string;
  setSecondaryState(value: string): void;

  getSecondaryZip(): string;
  setSecondaryZip(value: string): void;

  getSecondaryPhone(): string;
  setSecondaryPhone(value: string): void;

  getSecondarySpecialty(): heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap];
  setSecondarySpecialty(value: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap]): void;

  getSecondaryHours(): string;
  setSecondaryHours(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Provider.AsObject;
  static toObject(includeInstance: boolean, msg: Provider): Provider.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Provider, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Provider;
  static deserializeBinaryFromReader(message: Provider, reader: jspb.BinaryReader): Provider;
}

export namespace Provider {
  export type AsObject = {
    providerId: string,
    firstName: string,
    lastName: string,
    addressLinesList: Array<string>,
    city: string,
    state: string,
    zip: string,
    phone: string,
    specialty: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap],
    npi: string,
    facility: string,
    hours: string,
    ribbonId: string,
    secondaryFacility: string,
    secondaryAddressLinesList: Array<string>,
    secondaryCity: string,
    secondaryState: string,
    secondaryZip: string,
    secondaryPhone: string,
    secondarySpecialty: heyrenee_v1_enums_specialty_pb.SpecialtyMap[keyof heyrenee_v1_enums_specialty_pb.SpecialtyMap],
    secondaryHours: string,
  }
}

