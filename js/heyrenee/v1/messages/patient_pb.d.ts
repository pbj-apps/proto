// package: heyrenee.v1.messages
// file: heyrenee/v1/messages/patient.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as heyrenee_v1_options_auth_pb from "../../../heyrenee/v1/options/auth_pb";
import * as heyrenee_v1_enums_language_pb from "../../../heyrenee/v1/enums/language_pb";

export class Patient extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getMiddleName(): string;
  setMiddleName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getPrimaryPhone(): string;
  setPrimaryPhone(value: string): void;

  hasDateOfBirth(): boolean;
  clearDateOfBirth(): void;
  getDateOfBirth(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateOfBirth(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getSex(): SexMap[keyof SexMap];
  setSex(value: SexMap[keyof SexMap]): void;

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

  getPreferredName(): string;
  setPreferredName(value: string): void;

  getMaritalStatus(): MaritalStatusMap[keyof MaritalStatusMap];
  setMaritalStatus(value: MaritalStatusMap[keyof MaritalStatusMap]): void;

  getPreferredLanguage(): heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap];
  setPreferredLanguage(value: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap]): void;

  getEthnicity(): EthnicityMap[keyof EthnicityMap];
  setEthnicity(value: EthnicityMap[keyof EthnicityMap]): void;

  getMobilePhone(): string;
  setMobilePhone(value: string): void;

  getOtherPhone(): string;
  setOtherPhone(value: string): void;

  getNotes(): string;
  setNotes(value: string): void;

  getNamePronunciation(): string;
  setNamePronunciation(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Patient.AsObject;
  static toObject(includeInstance: boolean, msg: Patient): Patient.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Patient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Patient;
  static deserializeBinaryFromReader(message: Patient, reader: jspb.BinaryReader): Patient;
}

export namespace Patient {
  export type AsObject = {
    patientId: string,
    firstName: string,
    middleName: string,
    lastName: string,
    primaryPhone: string,
    dateOfBirth?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    sex: SexMap[keyof SexMap],
    email: string,
    addressLinesList: Array<string>,
    city: string,
    state: string,
    zip: string,
    preferredName: string,
    maritalStatus: MaritalStatusMap[keyof MaritalStatusMap],
    preferredLanguage: heyrenee_v1_enums_language_pb.LanguageMap[keyof heyrenee_v1_enums_language_pb.LanguageMap],
    ethnicity: EthnicityMap[keyof EthnicityMap],
    mobilePhone: string,
    otherPhone: string,
    notes: string,
    namePronunciation: string,
  }
}

export interface SexMap {
  SEX_UNSPECIFIED: 0;
  SEX_MALE: 1;
  SEX_FEMALE: 2;
}

export const Sex: SexMap;

export interface MaritalStatusMap {
  MARITAL_STATUS_UNSPECIFIED: 0;
  MARITAL_STATUS_SINGLE: 1;
  MARITAL_STATUS_MARRIED: 2;
}

export const MaritalStatus: MaritalStatusMap;

export interface EthnicityMap {
  ETHNICITY_UNSPECIFIED: 0;
  ETHNICITY_ASIAN_AMERICAN: 1;
  ETHNICITY_BLACK_OR_AFRICAN_AMERICAN: 2;
  ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN: 3;
  ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE: 4;
  ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDER: 5;
}

export const Ethnicity: EthnicityMap;

